package inspect

import (
	"github.com/mkadirtan/feed-toolbelt/pkg/header_node"
)

func (i *Inspector) applyStrategyHeader() {
	linkHeaders := i.header.Values("link")

	for _, linkHeader := range linkHeaders {
		headerNode := header_node.NewHeaderNode()
		headerNode.ParseFields(linkHeader)
		if headerNode.IsValidFeed() {
			i.processFeedCandidate(headerNode.FeedURL(), false)
		}
	}
}
