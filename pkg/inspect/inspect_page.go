package inspect

import (
	"github.com/mkadirtan/feed-toolbelt/pkg/link_node"
	"github.com/mkadirtan/feed-toolbelt/pkg/script_node"

	"golang.org/x/net/html"
)

func (i *Inspector) applyStrategyPage() {
	z := html.NewTokenizer(i.body)

	for {
		tt := z.Next()

		if err := z.Err(); err != nil {
			break
		}

		if tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tagBytes, hasAttributes := z.TagName()
		if !hasAttributes {
			continue
		}
		tag := string(tagBytes)

		switch tag {
		case "a", "link":
			feedNode := link_node.NewLinkNode(tag)
			feedNode.ParseFields(z)
			if feedNode.IsValidFeed() {
				i.processFeedCandidate(feedNode.FeedURL(), false)
			}
		case "script":
			scriptNode := script_node.NewScriptNode(tag)
			scriptNode.ParseFields(z)
			if scriptNode.IsValidFeed() {
				i.processFeedCandidate(scriptNode.FeedURL(), false)
			}
		}
	}
}
