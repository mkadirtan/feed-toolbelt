// HTTP header
// Link: <http://example.com/feed>; rel="alternate"; type="application/rss+xml"; title="RSS"

package header_node

import (
	"net/url"
	"strings"

	"github.com/mkadirtan/feed-toolbelt/pkg/attributes"
)

type HeaderNode struct {
	nodeHref  string
	nodeRel   string
	nodeType  string
	nodeTitle string
}

func NewHeaderNode() HeaderNode {
	return HeaderNode{}
}

func (h *HeaderNode) ParseFields(linkHeader string) {
	// Link: <http://example.com/feed>; rel="alternate"; type="application/rss+xml"; title="RSS"
	parts := strings.Split(linkHeader, ";")
	for _, part := range parts {
		if strings.HasSuffix(part, ">") && strings.HasPrefix(part, "<") {
			rawURL := part[1 : len(part)-1]
			u, err := url.Parse(rawURL)
			if err != nil {
				continue
			}

			h.nodeHref = u.String()
			continue
		}

		key, value, found := strings.Cut(part, "=")
		if !found {
			continue
		}
		key = strings.TrimSpace(key)

		value, found = strings.CutSuffix(value, `"`)
		if !found {
			continue
		}
		value, found = strings.CutPrefix(value, `"`)
		if !found {
			continue
		}

		switch key {
		case "rel":
			h.nodeRel = value
		case "type":
			h.nodeType = value
		case "title":
			h.nodeTitle = value
		}

		if h.nodeTitle != "" && h.nodeType != "" && h.nodeRel != "" {
			break
		}
	}
}

func (h *HeaderNode) checkRel() (attributes.RelType, error) {
	return attributes.ParseRelType(h.nodeRel)
}

func (h *HeaderNode) checkType() (attributes.FeedType, error) {
	return attributes.ParseFeedType(h.nodeType)
}

func (h *HeaderNode) checkHref() error {
	return attributes.ParseHref(h.nodeHref)
}

func (h *HeaderNode) IsValidFeed() bool {
	if _, err := h.checkRel(); err != nil {
		return false
	}

	if _, err := h.checkType(); err != nil {
		return false
	}

	if err := h.checkHref(); err != nil {
		return false
	}

	return true
}

func (h *HeaderNode) FeedURL() string {
	return h.nodeHref
}
