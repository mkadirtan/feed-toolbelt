package util

import (
	"io"

	"github.com/mmcdole/gofeed"
)

func ValidateFeed(r io.Reader) bool {
	feedType := gofeed.DetectFeedType(r)
	return feedType != gofeed.FeedTypeUnknown
}
