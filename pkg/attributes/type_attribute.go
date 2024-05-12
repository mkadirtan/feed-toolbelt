package attributes

import "errors"

var ErrUnrecognizedFeedType = errors.New("unrecognized feed type")

type FeedType int

const (
	FeedTypeRSS1 FeedType = iota
	FeedTypeRSS
	FeedTypeAtom
	FeedTypeJSON
	FeedTypeJSONGeneric
	FeedTypeXMLGeneric
)

func ParseFeedType(feedType string) (FeedType, error) {
	switch feedType {
	case "application/rdf+xml":
		return FeedTypeRSS1, nil
	case "application/rss+xml":
		return FeedTypeRSS, nil
	case "application/atom+xml":
		return FeedTypeAtom, nil
	case "application/feed+json":
		return FeedTypeJSON, nil
	case "application/json":
		return FeedTypeJSONGeneric, nil
	case "application/xml":
		return FeedTypeXMLGeneric, nil
	default:
		return 0, ErrUnrecognizedFeedType
	}
}

func (f FeedType) String() string {
	switch f {
	case FeedTypeRSS1:
		return "application/rdf+xml"
	case FeedTypeRSS:
		return "application/rss+xml"
	case FeedTypeAtom:
		return "application/atom+xml"
	case FeedTypeJSON:
		return "application/feed+json"
	case FeedTypeJSONGeneric:
		return "application/json"
	case FeedTypeXMLGeneric:
		return "application/xml"
	default:
		return ""
	}
}
