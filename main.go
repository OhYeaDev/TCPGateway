package main

import (
	"log"
	"net/http"

	"tcpgateway/gateway"
)

func main() {
	gateway.Load()
        gateway.LoadKeys()

	mux := http.NewServeMux()

	mux.HandleFunc("/", gateway.Handle)
	mux.Handle("/panel/", http.StripPrefix("/panel/", http.FileServer(http.Dir("web"))))

	log.Println("TCPGateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
