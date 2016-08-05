package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	hub := NewHub()
	router := mux.NewRouter()
	router.HandleFunc("/ws/{room}", hub.HandleWS).Methods("GET")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
