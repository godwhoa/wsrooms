package main

import (
	"github.com/gorilla/websocket"
	"log"
)

/* To figure out if they wanna broadcast to all or broadcast to all except them */
type Message struct {
	mtype string
	msg   []byte
}

/* Reads and writes messages from client */
type Client struct {
	conn *websocket.Conn
	in   chan []byte
	out  chan Message
}

/* Reads and pumps to out channel */
func (c *Client) ReadLoop() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		c.out <- Message{"ex", message}
	}
}

/* Reads from in channel and pumps to client */
func (c *Client) WriteLoop() {
	for {
		in := <-c.in
		c.WriteMessage(in)
	}
}

/* Writes a message to the client */
func (c *Client) WriteMessage(msg []byte) {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
	}
}
