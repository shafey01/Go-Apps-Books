package todolist_test

import (
	"testing"
	todolist "todo-list"
)

func TestAdd(t *testing.T) {
	l := todolist.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

}
