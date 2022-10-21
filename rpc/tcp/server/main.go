package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/seteveanlorn/learning-grpc/rpc/todo"
)

func main() {
	td := todo.NewTodo()

	rpcServer := rpc.NewServer()
	if err := rpcServer.Register(td); err != nil {
		log.Fatalf("Could not register todo: %v\n", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Could not listen on TCP port 1234: %v\n", err)
	}

	log.Println("Listening on TPC port :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection: %v\n", err)
			continue
		}

		rpcServer.ServeConn(conn)
	}
}
