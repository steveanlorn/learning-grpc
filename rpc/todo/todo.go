package todo

import (
	"log"

	"github.com/google/uuid"
)

type Todo struct {
	tasks map[string]Task
}

func NewTodo() *Todo {
	todo := Todo{
		tasks: make(map[string]Task),
	}

	return &todo
}

func (t *Todo) GetTaskByID(taskID string, reply *Task) error {
	if val, ok := t.tasks[taskID]; ok {
		reply.ID = val.ID
		reply.Name = val.Name
		reply.Status = val.Status
		return nil
	}

	return ErrTaskNotFound
}

func (t *Todo) CreateTask(arg Task, replyNewTaskID *string) error {
	newID := uuid.New().String()

	newTask := Task{
		ID:     newID,
		Name:   arg.Name,
		Status: arg.Status,
	}

	t.tasks[newID] = newTask

	*replyNewTaskID = newTask.ID

	log.Printf("Creating a new task: %v", newTask)

	return nil
}
