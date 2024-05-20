package util

import (
	"github.com/goware/urlx"
)

func NormalizeURL(rawURL string) (string, error) {
	u, err := urlx.ParseWithDefaultScheme(rawURL, "https")
	if err != nil {
		return "", err
	}

	return urlx.Normalize(u)
}
