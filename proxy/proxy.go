package proxy

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var targetURL = os.Getenv("TARGET_URL")

func ProxyRequest(w http.ResponseWriter, r *http.Request) {
	c := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := c.Get(targetURL)
	if err != nil {
		log.Printf("could not reach target RSS feed: %s", err)
		http.Error(w, "Failed to call remote RSS feed", http.StatusInternalServerError)
		return
	}

	copyHeaders(resp.Header, w.Header(), []string{"content-type", "cache-control", "access-control-allow-origin", "etag", "content-length"})
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func copyHeaders(src, dest http.Header, headers []string) {
	for _, k := range headers {
		dest.Set(k, src.Get(k))
	}
}
