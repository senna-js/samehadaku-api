package external

import (
	"strings"
)

func ExtractSlug(url string) string {
	if len(url) < 3 {
		return ""
	}

	if url[len(url)-1] == '/' {
		url = url[0 : len(url)-1]
	}

	urls := strings.Split(url, "/")

	return urls[len(urls)-1]
}
