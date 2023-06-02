package chat

import (
	"net"
)

type Room struct {
	name    string
	members map[net.Addr]*Client
}

func (r *Room) broadcast(sender *Client, msg string) {
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
