package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	todoList := List{}
	todoList.Add("test").Add("test2").Add("test3")

	actual := len(todoList)
	expected := 3
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDelete(t *testing.T) {
	todoList := List{}
	todoList.Add("test").Remove(1)

	actual := len(todoList)
	expected := 0
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestEdit(t *testing.T) {
	const EditTitle = "test_edited"
	todoList := List{}
	todoList.Add("test")
	todoList.Edit(1, EditTitle)

	actual := todoList[1].Title
	expected := EditTitle
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestChangeStatus(t *testing.T) {
	todoList := List{}
	todoList.Add("test")

	statusBefore := todoList[1].Completed
	todoList.ChangeStatus(1)
	statusAfter := todoList[1].Completed

	actual := statusAfter
	expected := !statusBefore
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
