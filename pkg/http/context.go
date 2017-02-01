package http

import (
	"context"
	"errors"

	"github.com/danehans/todo/pkg/types"
)

// unexported key prevents collisions
type key int

const (
	todoKey key = iota
)

var (
	errNoTodoFromContext = errors.New("api: Context missing a Todo")
)

// withTodo returns a copy of ctx that stores the given Todo.
func withTodo(ctx context.Context, todo *types.Todo) context.Context {
	return context.WithValue(ctx, todoKey, todo)
}

// todoFromContext returns the Todo from the ctx.
func todoFromContext(ctx context.Context) (*types.Todo, error) {
	todo, ok := ctx.Value(todoKey).(*types.Todo)
	if !ok {
		return nil, errNoTodoFromContext
	}
	return todo, nil
}
