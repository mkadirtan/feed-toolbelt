package inspect

import (
	"net/http"
	"strings"

	"github.com/mkadirtan/feed-toolbelt/pkg/common_paths"
)

func inspectCommonPaths(url string) ([]string, bool) {
	var feeds = make([]string, 0)

	strippedURL, _ := strings.CutSuffix(url, "/")

	for path := range common_paths.CommonPaths {
		candidateFeedURL := strippedURL + path
		resp, err := http.DefaultClient.Get(candidateFeedURL)
		if err != nil {
			continue
		}

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			continue
		}

		feeds = append(feeds, candidateFeedURL)
	}

	if len(feeds) > 0 {
		return feeds, true
	} else {
		return nil, false
	}
}
