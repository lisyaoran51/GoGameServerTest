package main

import (
	"context"
	"fmt"
	"net"
)

func main() {

	ip := "127.0.0.1"
	port := "8001"

	ipaddr := ip + ":" + port
	listener, err := net.Listen("tcp", ipaddr)
	if err != nil {
		fmt.Printf("fail to listen tcp %v", err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("fail to accept listener %v", err)
		return
	}

	connection := NewConnection(&conn)
	GetConnectionManager().AddConnection(GATE, connection)

	ctx, cancel := context.WithCancel(context.Background())
	go connection.Start(ctx)

	cancelChan := make(chan int, 0)
	for {
		select {
		case <-cancelChan:
			cancel()
			return
		}
	}

}
