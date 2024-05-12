package attributes

import "errors"

var ErrUnrecognizedRelType = errors.New("unrecognized rel type")

type RelType int

const (
	RelTypeAlternate RelType = iota
)

func ParseRelType(rel string) (RelType, error) {
	switch rel {
	case "alternate":
		return RelTypeAlternate, nil
	default:
		return 0, ErrUnrecognizedRelType
	}
}

func (r RelType) String() string {
	switch r {
	case RelTypeAlternate:
		return "alternate"
	default:
		return ""
	}
}
