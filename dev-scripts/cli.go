package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mtlynch/rss-proxy/proxy"
)

func main() {
	var port = os.Getenv("PORT")
	log.Print("starting up")
	http.HandleFunc("/", proxy.ProxyRequest)
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
