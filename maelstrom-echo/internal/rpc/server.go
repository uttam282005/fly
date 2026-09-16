package rpc

import (
	"log"
	"net"
	"net/rpc"
)

type LocalService struct{}

type Args struct{}
type Reply struct {
	Id int64
}

func (ls *LocalService) Start() {
	math := new(LocalService)

	rpc.Register(math)

	// Bind to a local port
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error:", err)
	}
	defer listener.Close()

	log.Println("RPC Server serving on port 1234...")

	// Accept connections indefinitely
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// ServeConn runs the gob-codec communication in its own goroutine
		go rpc.ServeConn(conn)
	}
}
