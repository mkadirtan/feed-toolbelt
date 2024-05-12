// Feed discovery methods
// Link tags with rel=alternate
// <link rel="alternate" type="application/rss+xml" title="RSS Feed" href="http://example.com/rss">
// HTTP header
// Link: <http://example.com/feed>; rel="alternate"; type="application/rss+xml"; title="RSS"
// Anchor Tags
// <a href="http://example.com/feed.xml">Subscribe to our feed</a>
// Structured JSON Data in json+ld format
// <script type="application/ld+json">
//{
//  "@context": "http://schema.org",
//  "@type": "WebSite",
//  "url": "http://example.com",
//  "potentialAction": {
//    "@type": "SubscribeAction",
//    "target": {
//      "@type": "EntryPoint",
//      "urlTemplate": "http://example.com/feed.json"
//    }
//  }
//}
//</script>
// JSON+LD schemas use "urlTemplate" instead of fixed urls. They may require some variables such as:
// "urlTemplate": "http://example.com/users/{userId}"
// In that case, the url should be ignored to avoid confusion.

// Feed confirmation methods, probably not the best idea to confirm feeds from this library.
// These are checked just so that if a feed url is directly given to the program, it will at least export itself again.
// link to self in feed contents
// <link rel="self" type="application/atom+xml" href="http://example.com/feed.atom" />
// <atom:link href="http://example.com/feed.rss" rel="self" type="application/rss+xml" />
package inspect

import (
	"net/http"
)

func InspectURL(url string) []string {
	var feedURLs = make([]string, 0)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil
	}

	feedOnHeader, found := inspectHeaders(resp.Header)
	if found {
		feedURLs = append(feedURLs, feedOnHeader)
	}

	feedsOnPage, found := inspectPage(resp.Body)
	if found {
		for _, feedOnPage := range feedsOnPage {
			feedURLs = append(feedURLs, feedOnPage)
		}
	}

	feedsOnCommonPaths, found := inspectCommonPaths(url)
	if found {
		for _, feedOnCommonPath := range feedsOnCommonPaths {
			feedURLs = append(feedURLs, feedOnCommonPath)
		}
	}

	return feedURLs
}
