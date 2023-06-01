package chat

import (
	"fmt"
	"net"
)

type Server struct {
	rooms map[string]*Room
}

func NewServer() *Server {
	return &Server{
		rooms: make(map[string]*Room),
	}
}

func (s *Server) NewClient(conn net.Conn) *Client {
	fmt.Printf("new client has connected: %s\n", conn.RemoteAddr().String())
	return &Client{
		conn: conn,
		nick: "anon_user",
	}
}
