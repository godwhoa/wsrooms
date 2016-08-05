package main

/* Has a name, clients, count which holds the actual coutn and index which acts as the unique id */
type Room struct {
	name    string
	clients map[int]*Client
	count   int
	index   int
}

/* Add a conn to clients map so that it can be managed */
func (r *Room) Join(conn *websocket.Conn) int {
	r.clients[r.index] = conn
	log.Printf("New Client joined %s", r.name)
	r.index++
	r.count++
	return r.index
}

/* Send to specific client */
func (r *Room) SendTo(id int, msg []byte) {
	r.clients[id].in <- msg
}

/* Broadcast to every client */
func (r *Room) BroadcastAll(msg []byte) {
	for id, client := range r.clients {
		client.in <- msg
	}
}

/* Broadcast to all except */
func (r *Room) BroadcastEx(senderid int, msg []byte) {
	for id, client := range r.clients {
		if id != senderid {
			client.in <- msg
		}
	}
}

/* Constructor */
func NewRoom(name string) *Room {
	room := &Room{}
	room.name = name
	room.clients = make(map[int]*Client)
	room.count = 0
	room.index = 0
	return room
}
