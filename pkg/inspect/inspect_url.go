// Package inspect
// Feed Discovery Methods
// - Inside link, and a tags with rel="alternate" attribute
// - Inside script tag with structured json data
// - HTTP Header ( Link Header )
// - Brute forcing common feed paths
// Feed Validation Methods
// - Not Implemented Yet
package inspect

import (
	"io"
	"net/http"
	"slices"

	"github.com/mkadirtan/feed-toolbelt/pkg/util"
)

func InspectHTML(r io.Reader) []string {
	feedsOnPage, found := inspectPage(r)
	if found {
		return feedsOnPage
	} else {
		return nil
	}
}

func InspectURL(url string, checkHeaders bool, checkPage bool, checkCommonPaths bool, validateFeeds bool) []string {
	var feedURLs = make([]string, 0)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil
	}

	if checkHeaders {
		feedsOnHeader, found := inspectHeaders(resp.Header)
		if found {
			for _, feedOnHeader := range feedsOnHeader {
				feedURLs = append(feedURLs, feedOnHeader)
			}
		}
	}

	if checkPage {
		feedsOnPage, found := inspectPage(resp.Body)
		if found {
			for _, feedOnPage := range feedsOnPage {
				feedURLs = append(feedURLs, feedOnPage)
			}
		}
	}

	if validateFeeds {
		for i, u := range feedURLs {
			if response, httpErr := http.DefaultClient.Get(u); httpErr != nil || response.StatusCode < 200 || response.StatusCode > 299 {
				continue
			} else {
				if !util.ValidateFeed(response.Body) {
					slices.Delete(feedURLs, i, i+1)
				}
			}
		}
	}

	if checkCommonPaths {
		feedsOnCommonPaths, found := inspectCommonPaths(url)
		if found {
			for _, feedOnCommonPath := range feedsOnCommonPaths {
				feedURLs = append(feedURLs, feedOnCommonPath)
			}
		}
	}

	return feedURLs
}
