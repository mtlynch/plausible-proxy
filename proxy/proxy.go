package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const plausibleBaseURL = "https://plausible.io"

var (
	plausibleScriptURL = mustParseURL(plausibleBaseURL + "/js/plausible.js")
	plausibleEventURL  = mustParseURL(plausibleBaseURL + "/api/event")
)

func ProxyPlausibleScript(w http.ResponseWriter, r *http.Request) {
	proxy := makePlausibleProxy(plausibleScriptURL)
	proxy.ServeHTTP(w, r)
}

func ProxyPlausibleEvent(w http.ResponseWriter, r *http.Request) {
	proxy := makePlausibleProxy(plausibleEventURL)
	proxy.ServeHTTP(w, r)
}

func makePlausibleProxy(targetURL *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL = targetURL
		req.Host = targetURL.Host

		req.Header.Add("X-Forwarded-Proto", req.Proto)
		req.Header.Add("X-Forwarded-Host", req.Host)
	}
	return proxy
}

func mustParseURL(u string) *url.URL {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Printf("failed to parse %s", u)
		panic(err)
	}
	return parsed
}
