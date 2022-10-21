package main

import (
	"log"
	"net/rpc"

	"github.com/seteveanlorn/learning-grpc/rpc/todo"
)

func main() {
	rpcClient, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Could not dial TCP server: %v\n", err)
	}

	defer rpcClient.Close()

	newTask := todo.Task{
		Name:   "Register ID at ward office",
		Status: todo.StatusPending,
	}

	var newTaskID string
	if err := rpcClient.Call("Todo.CreateTask", newTask, &newTaskID); err != nil {
		log.Fatalf("Could not call 'Todo.CreateTask': %v\n", err)
	}

	log.Println(newTaskID)

	var taskReply todo.Task
	if err := rpcClient.Call("Todo.GetTaskByID", newTaskID, &taskReply); err != nil {
		log.Fatalf("Could not call 'Todo.GetTaskByID': %v\n", err)
	}

	log.Printf("%v", taskReply)
}
