package attributes

import (
	"errors"
	"net/url"
)

func ParseHref(href string) error {
	if href == "" {
		return errors.New("empty href")
	}

	u, err := url.Parse(href)
	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("invalid scheme")
	}

	return nil
}
