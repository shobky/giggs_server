package _websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Brodcast   chan Message
	Clients    map[*Client]bool
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Brodcast:   make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			if p.Clients[client] == true {
				break
			}
			p.Clients[client] = true
			fmt.Println("new client joined, id:", len(p.Clients))
			for client := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "new user joined.."})
			}
			break
		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println(fmt.Sprintf("client %v disconected, size of pool:", client.ID), len(p.Clients))
			for client := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "user disconnected.."})
			}
			break
		case message := <-p.Brodcast:
			for client := range p.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	}
}
