package store_test

import (
	"fmt"
	"github.com/danehans/todo/pkg/store"
	"github.com/danehans/todo/pkg/types"
	"testing"
)

var testTodo types.Todo

const succeed = "\u2713"
const failed = "\u2717"

var id int

func TestCreateRepo(t *testing.T) {

	testTodo := types.Todo{Name: "Testing"}
	t.Logf("Testing Repo")
	t.Logf("\tTesting creating a repo")
	createdTodo := store.CreateTodo(testTodo)
	id = createdTodo.Id

	if id == 0 {
		t.Errorf("\t%s testTodo id not incremented", failed)
	} else {
		t.Logf("\t%s ToDo was created", succeed)
	}

	storedTodo := store.ListTodo(id)
	if storedTodo.Name != testTodo.Name {
		t.Errorf("\t%s testTodo name does not match", failed)
	} else {
		t.Logf("\t%s Created Todo name matches stored Todo", succeed)
	}

}

func TestDeleteRepo(t *testing.T) {

	t.Logf("\tTesting deleting a repo")

	if store.DeleteTodo(id) != nil {
		t.Errorf("\t%s Error deleting Todo")
	} else {
		t.Logf("\t%s ToDo was deleted", succeed)
	}
	storedTodo := store.ListTodo(id)
	if storedTodo.Name != "" {
		t.Errorf("\t%s StoredTodo name not empty", failed)
	} else {
		t.Logf("\t%s Stored Todo was removed from list", succeed)
	}

}

func ExampleRepo() {
	storedTodo := store.ListTodo(1)
	fmt.Println(storedTodo.Name)
	// Output:
	//    Write presentation

}
