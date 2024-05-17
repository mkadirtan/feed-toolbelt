// Package inspect
// Feed Discovery Methods
// - Inside link, and a tags with rel="alternate" attribute
// - Inside script tag with structured json data
// - HTTP Header ( Link Header )
// - Brute forcing common feed paths
package inspect

import (
	"net/http"
	"slices"
	"strings"

	"github.com/mkadirtan/feed-toolbelt/pkg/common_paths"
	"github.com/mkadirtan/feed-toolbelt/pkg/util"
)

func (i *Inspector) Find() {
	if i.config.PipedInput != nil {
		i.findTargetHTML()
		return
	}

	if i.config.TargetURL != nil {
		i.findTargetURL()
		return
	}
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

func (i *Inspector) findTargetHTML() {
	feedsOnPage, _ := inspectPage(*i.config.PipedInput)
	for _, feed := range feedsOnPage {
		i.processFeedCandidate(feed, false)
	}
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

func (i *Inspector) findTargetURL() {
	if i.config.Strategies.Header || i.config.Strategies.Page {
		i.pageAndHeadersStrategy()
	}

	if i.config.Strategies.Common {
		i.commonStrategy()
	}
}

func (i *Inspector) commonStrategy() {
	strippedURL, _ := strings.CutSuffix(*i.config.TargetURL, "/")
	for path := range common_paths.CommonPaths {
		feedCandidate := strippedURL + path

		if slices.Contains(i.foundFeeds, feedCandidate) {
			continue
		}

		i.processFeedCandidate(feedCandidate, true)
	}
}

func (i *Inspector) pageAndHeadersStrategy() {
	resp, err := http.DefaultClient.Get(*i.config.TargetURL)
	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return
	}

	if i.config.Strategies.Header {
		feedsOnHeader, _ := inspectHeaders(resp.Header)
		for _, feed := range feedsOnHeader {
			i.processFeedCandidate(feed, false)
		}
	}

	if i.config.Strategies.Page {
		feedsOnPage, _ := inspectPage(resp.Body)
		for _, feed := range feedsOnPage {
			i.processFeedCandidate(feed, false)
		}
	}
}
