package main

import (
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()

	serveMux := http.NewServeMux()

	//Web Page Flutter
	serveMux.Handle("/", http.FileServer(http.Dir("public")))

	//Server WebSocket
	serveMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(hub, w, r)
	})

	log.Println("Running server :8080")
	log.Println(http.ListenAndServe(":8080", serveMux))
}
