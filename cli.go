package main

import (
	"log"
	"net/http"

	"github.com/mtlynch/rss-proxy/proxy"
)

func main() {
	log.Print("starting up")
	http.HandlerFunc("/", proxy.ProxyRequest)
}
