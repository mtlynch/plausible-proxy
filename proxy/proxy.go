package proxy

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func ProxyRequest(w http.ResponseWriter, r *http.Request) {
	httputil.NewSingleHostReverseProxy(url.Parse("https://plausible.io/api/event")).ServeHTTP(w, r)
}
