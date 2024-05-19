package todolist

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// Item struct , Task, Done, CreatedAt, CompletedAt
// List []
// For this application, you’ll implement the following methods:
// 1. Complete: Marks a to-do item as completed.
// 2. Add: Creates a new to-do item and appends it to the list.
// 3. Delete: Deletes a to-do item from the list.
// 4. Save: Saves the list of items to a file using the JSON format.
// 5. Get: Obtains a list of items from a saved JSON file.

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {

	item := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, item)
}

func (l *List) Complete(i int) error {

	list := *l

	if i <= 0 || i > len(list) {

		return fmt.Errorf("item %d not exist int the list", i)

	}

	list[i-1].Done = true
	list[i-1].CompletedAt = time.Now()

	return nil

}

func (l *List) Delete(i int) error {

	list := *l

	if i <= 0 || i > len(list) {
		return fmt.Errorf("item %d not exist", i)

	}

	*l = append(list[:i-1], list[i:]...)
	return nil
}

func (l *List) Save(filename string) error {

	jsonList, err := json.Marshal(l)
	if err != nil {

		return err

	}

	// 4 2 1
	// r w x
	return os.WriteFile(filename, jsonList, 0644)

}

func (l *List) Get(filename string) error {

	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) <= 0 {
		return nil
	}

	return json.Unmarshal(file, l)

}
