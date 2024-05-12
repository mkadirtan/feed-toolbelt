package inspect

import (
	"io"

	"github.com/mkadirtan/feed-toolbelt/pkg/link_node"
	"github.com/mkadirtan/feed-toolbelt/pkg/script_node"
	"golang.org/x/net/html"
)

func inspectPage(r io.Reader) ([]string, bool) {
	var hrefs = make([]string, 0)
	z := html.NewTokenizer(r)

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
				hrefs = append(hrefs, feedNode.FeedURL())
			}
		case "script":
			scriptNode := script_node.NewScriptNode(tag)
			scriptNode.ParseFields(z)
			if scriptNode.IsValidFeed() {
				hrefs = append(hrefs, scriptNode.FeedURL())
			}
		}
	}

	if len(hrefs) > 0 {
		return hrefs, true
	}

	return nil, false
}
