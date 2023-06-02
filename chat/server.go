package chat

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	rooms    map[string]*Room
	commands chan Command
}

func NewServer() *Server {
	return &Server{
		rooms:    make(map[string]*Room),
		commands: make(chan Command),
	}
}

func (s *Server) nick(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("nick is required. usage: /nick NAME")
		return
	}
	c.nick = args[1]
	c.msg(fmt.Sprintf("all right, I will call you %s", c.nick))
}
func (s *Server) NewClient(conn net.Conn) *Client {
	fmt.Printf("new client has connected: %s\n", conn.RemoteAddr().String())
	return &Client{
		conn:     conn,
		nick:     "anon_user",
		commands: s.commands,
	}
}
func (s *Server) Run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		}
	}
}
func (s *Server) join(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("room name is required. usage: /join ROOM_NAME")
		return
	}
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &Room{
			name:    roomName,
			members: make(map[net.Addr]*Client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c
	c.room = r
	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))
	c.msg(fmt.Sprintf("welcome to %s", roomName))
}

func (s *Server) quit(c *Client) {
	log.Printf("client has lefr the chat: %s", c.conn.RemoteAddr().String())
	c.msg("mkay!")
	c.conn.Close()
}
