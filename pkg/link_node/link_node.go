package link_node

import (
	"github.com/mkadirtan/feed-toolbelt/pkg/attributes"
	"golang.org/x/net/html"
)

func NewLinkNode(tag string) LinkNode {
	return LinkNode{
		nodeTag: tag,
	}
}

type LinkNode struct {
	// nodeTag is used for reporting only
	nodeTag  string
	nodeRel  string
	nodeHref string
	nodeType string
}

func (f *LinkNode) checkRel() (attributes.RelType, error) {
	return attributes.ParseRelType(f.nodeRel)
}

func (f *LinkNode) checkType() (attributes.FeedType, error) {
	return attributes.ParseFeedType(f.nodeType)
}

func (f *LinkNode) checkHref() error {
	return attributes.ParseHref(f.nodeHref)
}

// IsValidFeed future work:
//  1. In case of a soft error, such as rel or type error, try to validate the href contains a feed
//  2. Log debug information FeedType, RelType and Tag
//     Found feed without errors,
//     Found feed with warnings,
//     Could not find feed due to errors,
func (f *LinkNode) IsValidFeed() bool {
	_, relErr := f.checkRel()
	if relErr != nil {
		return false
	}

	_, typeErr := f.checkType()
	if typeErr != nil {
		return false
	}

	hrefErr := f.checkHref()
	if hrefErr != nil {
		return false
	}

	return true
}

func (f *LinkNode) FeedURL() string {
	return f.nodeHref
}

func (f *LinkNode) ParseFields(z *html.Tokenizer) {
	for {
		key, val, more := z.TagAttr()
		switch string(key) {
		case "href":
			f.nodeHref = string(val)
		case "type":
			f.nodeType = string(val)
		case "rel":
			f.nodeRel = string(val)
		}

		if !more {
			break
		}

		if f.nodeHref != "" && f.nodeType != "" && f.nodeRel != "" {
			break
		}
	}
}
