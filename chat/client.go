package chat

import (
	"net"
)

type Client struct {
	conn net.Conn
	nick string
	room *Room
}
