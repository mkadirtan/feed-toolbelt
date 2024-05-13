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
	"net/http"
)

func InspectURL(url string, checkHeaders bool, checkPage bool, checkCommonPaths bool) []string {
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
