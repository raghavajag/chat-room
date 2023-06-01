package chat

import (
	"net"
)

type Room struct {
	name    string
	members map[net.Addr]*Client
}
