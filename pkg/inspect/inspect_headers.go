package inspect

import (
	"net/http"

	"github.com/mkadirtan/feed-toolbelt/pkg/header_node"
)

func inspectHeaders(headers http.Header) ([]string, bool) {
	linkHeaders := headers.Values("link")
	var foundFeeds = make([]string, 0)
	for _, linkHeader := range linkHeaders {
		headerNode := header_node.NewHeaderNode()
		headerNode.ParseFields(linkHeader)
		if headerNode.IsValidFeed() {
			foundFeeds = append(foundFeeds, headerNode.FeedURL())
		}
	}

	if len(foundFeeds) > 0 {
		return foundFeeds, true
	}

	return nil, false
}
