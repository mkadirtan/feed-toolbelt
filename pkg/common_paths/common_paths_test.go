package common_paths

import (
	"strings"
	"testing"
)

func TestWellKnownUrls(t *testing.T) {
	var loopAtLeastOnce bool
	for url := range CommonPaths {
		loopAtLeastOnce = true
		if !(strings.HasPrefix(url, "/") && len(url) > 1) {
			t.Errorf("invalid url: %s\n", url)
		}
	}
	if !loopAtLeastOnce {
		t.Errorf("did not loop")
	}
}
