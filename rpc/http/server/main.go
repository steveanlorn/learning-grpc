package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/seteveanlorn/learning-grpc/rpc/todo"
)

func main() {
	td := todo.NewTodo()

	rpcServer := rpc.NewServer()
	if err := rpcServer.Register(td); err != nil {
		log.Fatalf("Could not register todo: %v\n", err)
	}

	rpcServer.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	log.Println("Listening on port :1234")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatalf("Could not run HTTP server on port 1234: %v\n", err)
	}
}
