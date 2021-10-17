package proxy

import (
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var (
	targetURL = mustParseURL("https://plausible.io")
)

func ProxyPlausible(w http.ResponseWriter, r *http.Request) {
	canonicalPath, err := canonicalizePath(r.URL.Path)
	if err != nil {
		log.Printf("path %s is not supported", r.URL.Path)
		http.Error(w, "Unsupported path", http.StatusNotFound)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL = targetURL
		req.URL.Path = canonicalPath
		req.Host = targetURL.Host

		req.Header.Add("X-Forwarded-Proto", req.Proto)
		req.Header.Add("X-Forwarded-Host", req.Host)
	}
	proxy.ServeHTTP(w, r)
}

func canonicalizePath(path string) (string, error) {
	mappings := map[string]string{
		"/api/event":       "/api/event",
		"/js/plausible.js": "/js/plausible.js",
		// Alias for plausible.js in case of ad-blockers who match the plausible.js
		// string.
		"/js/script.js": "/js/plausible.js",
	}
	for k, v := range mappings {
		// Use HasSuffix instead of direct matching because there might be a prefix
		// in the path, depending on how the client is forwarding to the proxy.
		if strings.HasSuffix(path, k) {
			return v, nil
		}
	}
	return "", errors.New("unsupported path")
}

// Only proxy requests to a whitelist of paths that are necessary for loading
// the plausible.js script and registering an event.
func isWhitelistedPath(path string) bool {
	whitelistedPaths := []string{
		"/js/plausible.js",
		"/api/event",
	}
	for _, p := range whitelistedPaths {
		if p == path {
			return true
		}
	}
	return false
}

func mustParseURL(u string) *url.URL {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Printf("failed to parse URL %s", u)
		panic(err)
	}
	return parsed
}
