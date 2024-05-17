package inspect

import (
	"slices"
	"strings"

	"github.com/mkadirtan/feed-toolbelt/pkg/common_paths"
)

func (i *Inspector) applyStrategyCommon() {
	strippedURL, _ := strings.CutSuffix(*i.config.TargetURL, "/")
	for path := range common_paths.CommonPaths {
		feedCandidate := strippedURL + path

		if slices.Contains(i.foundFeeds, feedCandidate) {
			continue
		}

		i.processFeedCandidate(feedCandidate, true)
	}
}
