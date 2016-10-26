package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/* Controls a bunch of rooms */
type Hub struct {
	hub      map[string]*Room
	upgrader websocket.Upgrader
}

/* If room doesn't exist creates it then returns it */
func (h *Hub) GetRoom(name string) *Room {
	if _, ok := h.hub[name]; !ok {
		h.hub[name] = NewRoom(name)
	}
	return h.hub[name]
}

/* Get ws conn. and hands it over to correct room */
func (h *Hub) HandleWS(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	room_name := params["room"]

	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close()
	room := h.GetRoom(room_name)
	id := room.Join(c)

	/* Reads from the client's out bound channel and broadcasts it */
	go room.HandleMsg(id)

	/* Reads from client and if this loop breaks then client disconnected. */
	room.clients[id].ReadLoop()
	room.Leave(id)
}

/* Constructor */
func NewHub() *Hub {
	hub := new(Hub)
	hub.hub = make(map[string]*Room)
	hub.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return hub
}
