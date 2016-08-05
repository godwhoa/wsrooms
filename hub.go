package main

/* Controls a bunch of rooms */
type Hub struct {
	hub      map[string]*Room
	upgrader websocket.Upgrader
}

/* If room doesn't exist creates it then returns it */
func (h *Hub) GetRoom(name string) *Room {
	if _, ok := h.hub[name]; ok {
		return hub[name]
	} else {
		h.hub[name] = NewRoom(name)
		return h.hub[name]
	}
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

	id := room.Join(c, room_name)
	go room.clients[id].Read()
	room.clients[id].Write()
}

/* Constructor */
func NewHub() *Hub {
	hub := &Hub{}
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
