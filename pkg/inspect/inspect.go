// Package inspect
// Feed Discovery Methods
// - Inside link, and a tags with rel="alternate" attribute
// - Inside script tag with structured json data
// - HTTP Header ( Link Header )
// - Brute forcing common feed paths
package inspect

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/mkadirtan/feed-toolbelt/pkg/util"
)

func (i *Inspector) Find() error {
	if i.config.PipedInput != nil {
		i.body = i.config.PipedInput
	} else {
		resp, err := http.DefaultClient.Get(*i.config.TargetURL)
		if err != nil {
			return err
		}

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			return errors.New(fmt.Sprintf("invalid status code: %d", resp.StatusCode))
		}

		i.body = bufio.NewReader(resp.Body)
		i.header = resp.Header
	}

	if i.config.Strategies.Header {
		i.applyStrategyHeader()
	}
	if i.config.Strategies.Page {
		i.applyStrategyPage()
	}
	if i.config.Strategies.Common {
		i.applyStrategyCommon()
	}

	return nil
}

func (i *Inspector) processFeedCandidate(feedCandidateURL string, mustValidate bool) {
	if slices.Contains(i.foundFeeds, feedCandidateURL) {
		return
	}

	if (mustValidate || i.config.Validate) && !i.validateFeedURL(feedCandidateURL) {
		return
	}

	i.config.OutputHandler(feedCandidateURL)
	i.foundFeeds = append(i.foundFeeds, feedCandidateURL)
}

func (i *Inspector) validateFeedURL(feedURL string) bool {
	resp, err := http.DefaultClient.Get(feedURL)
	if err != nil {
		return false
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return false
	}

	return util.ValidateFeed(resp.Body)
}
