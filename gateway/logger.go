package gateway

import (
	"log"
	"net/http"
)

func Log(r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
}
