package main

/* Reads and writes messages from client */
type Client struct {
	conn *websocket.Conn
	in   chan []byte
	out  chan []byte
}

/* Reads and pumps to out channel */
func (c *Client) Read() {
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		c.out <- message
	}
}

/* Reads from in channel and pumps to client */
func (c *Client) Write() {
	for {
		in := <-c.in
		c.WriteMessage(in)
	}
}

/* Writes a message to the client */
func (c *Client) WriteMessage(msg []byte) {
	err := c.conn.WriteMessage(1, in)
	if err != nil {
		log.Println("write:", err)
		break
	}
}
