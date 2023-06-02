package main

import (
	"chat/chat"
	"fmt"
	"net"
)

func main() {
	s := chat.NewServer()
	go s.Run()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %s\n", err)
			continue
		}
		c := s.NewClient(conn)
		go c.ReadInput()
	}
}
