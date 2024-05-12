package inspect

import (
	"net/http"

	"github.com/mkadirtan/feed-toolbelt/pkg/header_node"
)

func inspectHeaders(headers http.Header) (string, bool) {
	headerNode := header_node.NewHeaderNode()
	headerNode.ParseFields(headers)
	if headerNode.IsValidFeed() {
		return headerNode.FeedURL(), true
	}

	return "", false
}
