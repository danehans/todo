package store

import (
	"fmt"

	"github.com/danehans/todo/pkg/types"
)

var currentId int

var todos types.Todos

// Give us some seed data
func init() {
	CreateTodo(types.Todo{Name: "Write presentation"})
	CreateTodo(types.Todo{Name: "Host meetup"})
}

// ListTodo takes a todo id and returns the associated todo.
func ListTodo(id int) types.Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return types.Todo{}
}

// CreateTodo takes a todo and creates an instance of a todo,
// including the todo id.
func CreateTodo(t types.Todo) types.Todo {
	currentId = currentId + 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

// DeleteTodo removes a todo by id from the list of todo's.
func DeleteTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
