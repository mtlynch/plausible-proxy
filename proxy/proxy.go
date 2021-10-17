package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var target = mustParseURL("https://plausible.io/api/event")

func ProxyRequest(w http.ResponseWriter, r *http.Request) {
	httputil.NewSingleHostReverseProxy(target).ServeHTTP(w, r)
}

func mustParseURL(u string) *url.URL {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Printf("failed to parse %s", u)
		panic(err)
	}
	return parsed
}
