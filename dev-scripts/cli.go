package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mtlynch/plausible-proxy/proxy"
)

func main() {
	var port = os.Getenv("PORT")
	log.Print("starting up")
	http.HandleFunc("/", proxy.ProxyPlausible)
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
