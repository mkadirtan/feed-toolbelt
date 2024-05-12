package script_node

import (
	"bytes"
	"encoding/json"
	"regexp"

	"golang.org/x/net/html"
)

var dynamicURLRegex = regexp.MustCompile(`\{[^{}]*}`)
var scriptTypeJSONLD = "application/ld+json"

type Target struct {
	TargetType  string `json:"@type"`
	UrlTemplate string `json:"urlTemplate"`
}
type PotentialAction struct {
	ActionType string `json:"@type"`
	Target     Target `json:"target"`
}
type JSONLD struct {
	PotentialAction PotentialAction `json:"potentialAction"`
}

type ScriptNode struct {
	NodeTag  string
	NodeType string // type="application/ld+json
	JSONLD   JSONLD
}

func NewScriptNode(tag string) ScriptNode {
	return ScriptNode{
		NodeTag: tag,
	}
}

func (s *ScriptNode) ParseFields(z *html.Tokenizer) {
	for {
		key, val, more := z.TagAttr()
		switch string(key) {
		case "type":
			s.NodeType = string(val)
		}

		if !more {
			break
		}

		if s.NodeType != "" {
			break
		}
	}

	// Future Work: support more schemas
	rawScript := z.Text()

	var decodedJSON JSONLD
	err := json.NewDecoder(bytes.NewBuffer(rawScript)).Decode(&decodedJSON)
	if err != nil {
		return
	}

	s.JSONLD = decodedJSON
}

func (s *ScriptNode) IsValidFeed() bool {
	if s.NodeType != scriptTypeJSONLD {
		return false
	}

	u := s.JSONLD.PotentialAction.Target.UrlTemplate

	if u == "" {
		return false
	}

	isDynamicURL := dynamicURLRegex.MatchString(u)
	if isDynamicURL {
		return false
	}

	return true
}

func (s *ScriptNode) FeedURL() string {
	return s.JSONLD.PotentialAction.Target.UrlTemplate
}
