// Package inspect
// Feed Discovery Methods
// - Inside link, and a tags with rel="alternate" attribute
// - Inside script tag with structured json data
// - HTTP Header ( Link Header )
// - Brute forcing common feed paths
package inspect

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
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

		var buf []byte
		var newReader = bytes.NewBuffer(buf)
		if _, cErr := io.Copy(newReader, resp.Body); cErr != nil {
			return cErr
		}
		i.body = bufio.NewReader(newReader)
		i.header = resp.Header
	}

	// case when given URL itself is a feed URL
	if i.config.TargetURL != nil {
		body, err := io.ReadAll(i.body)
		if err != nil {
			return err
		}

		if util.ValidateFeed(bufio.NewReader(bytes.NewBuffer(body))) {
			i.processFeedCandidate(*i.config.TargetURL, false)
			return nil
		}

		i.body = bufio.NewReader(bytes.NewBuffer(body))
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
	feedCandidateURL, err := util.NormalizeURL(feedCandidateURL)

	if err != nil {
		return
	}

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
