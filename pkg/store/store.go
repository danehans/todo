package store

import (
	"fmt"

	"github.com/danehans/todo/pkg/types"
)

var currentId int

var todos types.Todos

// Give us some seed data
func init() {
	StoreCreateTodo(types.Todo{Name: "Write presentation"})
	StoreCreateTodo(types.Todo{Name: "Host meetup"})
}

func StoreFindTodo(id int) types.Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return types.Todo{}
}

func StoreCreateTodo(t types.Todo) types.Todo {
	currentId = currentId + 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
