package main

import (
	"log"

	"github.com/gorilla/websocket"
)

/* Has a name, clients, count which holds the actual coutn and index which acts as the unique id */
type Room struct {
	name    string
	clients map[int]*Client
	count   int
	index   int
}

/* Add a conn to clients map so that it can be managed */
func (r *Room) Join(conn *websocket.Conn) int {
	r.index++
	r.clients[r.index] = NewClient(conn)
	log.Printf("New Client joined %s", r.name)
	r.count++
	return r.index
}

/* Send to specific client */
func (r *Room) SendTo(id int, msg []byte) {
	r.clients[id].WriteMessage(msg)
}

/* Broadcast to every client */
func (r *Room) BroadcastAll(msg []byte) {
	for _, client := range r.clients {
		client.WriteMessage(msg)
	}
}

/* Broadcast to all except */
func (r *Room) BroadcastEx(senderid int, msg []byte) {
	for id, client := range r.clients {
		if id != senderid {
			client.WriteMessage(msg)
		}
	}
}

/* Handle messages */
func (r *Room) HandleMsg(id int) {
	for {
		if r.clients[id] == nil {
			break
		}
		out := <-r.clients[id].out
		if out.mtype == "ex" {
			r.BroadcastEx(id, out.msg)
		} else {
			r.BroadcastAll(out.msg)
		}
	}
}

/* Constructor */
func NewRoom(name string) *Room {
	room := new(Room)
	room.name = name
	room.clients = make(map[int]*Client)
	room.count = 0
	room.index = 0
	return room
}
