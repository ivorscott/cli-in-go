package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// item struct represents a ToDo item
type item struct {
	Task       string
	Done       bool
	CreateAt   time.Time
	CompleteAt time.Time
}

// List represents a list of ToDo items
type List []item

// Add creates a new todo item and appends it to the List
func (l *List) Add(task string) {
	t := item{
		Task:       task,
		Done:       false,
		CreateAt:   time.Now(),
		CompleteAt: time.Time{},
	}

	// you need to dereference the list in order to avoid changing a copy
	// you dereference with providing *l to the append call
	*l = append(*l, t)
}

// Complete method marks a ToDo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompleteAt = time.Now()
	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save method encodes the List as JSON and writes it to a file using the provided filename
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes the JSON data and
// parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
