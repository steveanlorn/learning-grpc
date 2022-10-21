package todo

import "fmt"

var (
	ErrTaskNotFound = fmt.Errorf("Task is not found")
)

const (
	StatusPending = iota + 1
	StatusOnProgress
	StatusDone
)

type Task struct {
	// ID identifies this task. This ID is assigned by the server and is
	// populated for newly created Task.
	//
	// This field is read-only.
	ID string

	Name   string
	Status int
}
