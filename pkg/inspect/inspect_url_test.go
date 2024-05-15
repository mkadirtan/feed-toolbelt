package inspect

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

type websiteMock struct {
	respHeaders map[string][]string
	respBody    []byte
}

func (m *websiteMock) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	for key := range m.respHeaders {
		for _, header := range m.respHeaders[key] {
			w.Header().Add(key, header)
		}
	}
	w.WriteHeader(200)
	_, _ = w.Write(m.respBody)
}

func TestInspectURLHeaders(t *testing.T) {
	feedUrl := "https://nooptoday.com/feed"
	linkHeader := fmt.Sprintf(`<%s>; rel="%s"; type="%s"; title="%s"`, feedUrl, "alternate", "application/rss+xml", "RSS")

	m := websiteMock{
		respHeaders: map[string][]string{"link": {"asd", linkHeader, "zxc"}},
		respBody:    nil,
	}

	s := httptest.NewServer(&m)
	feedLinks := InspectURL(s.URL, true, false, false, false)
	if !slices.Contains(feedLinks, feedUrl) {
		t.Errorf("expected feed ")
	}
}

func TestInspectURLPageLinks(t *testing.T) {
	feedUrl := "https://nooptoday.com/feed"
	body := fmt.Sprintf(`<html>
<head>
<link rel="alternate" type="application/rss+xml" href="%s" />
</head>
<body>
</body>
</html>`, feedUrl)

	m := websiteMock{
		respHeaders: nil,
		respBody:    []byte(body),
	}

	s := httptest.NewServer(&m)
	feedLinks := InspectURL(s.URL, false, true, false, false)
	if !slices.Contains(feedLinks, feedUrl) {
		t.Errorf("expected feed ")
	}
}

func TestInspectURLPageScripts(t *testing.T) {
	feedUrl := "https://nooptoday.com/feed"
	body := fmt.Sprintf(`<html>
<head>
<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "WebSite",
  "url": "https://nooptoday.com",
  "potentialAction": {
    "@type": "SubscribeAction",
    "target": {
      "@type": "EntryPoint",
      "urlTemplate": "%s"
    }
  }
}
</script>
</head>
<body>
</body>
</html>`, feedUrl)

	m := websiteMock{
		respHeaders: nil,
		respBody:    []byte(body),
	}

	s := httptest.NewServer(&m)
	feedLinks := InspectURL(s.URL, false, true, false, false)
	if !slices.Contains(feedLinks, feedUrl) {
		t.Errorf("expected feed")
	}
}
