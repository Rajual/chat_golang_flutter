package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
WriteBufferSize:1024
}

type Message struct {
	Nickname string `json:"nickname,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Client struct {
	nickname     string
	hub          *Hub
	conn         *websocket.Conn
	queryMessage chan Message
}

func (c *Client) writeWS() {

}

func (c *Client) readWS() {

}

func handleWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query()["nickname"]
	if len(nickname) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
