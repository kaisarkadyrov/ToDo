package todo_handlers

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title      string
	Completed  bool
	CreatedAt  time.Time
	FinishedAt *time.Time
}

type Todos []Todo

func (t *Todos) Add(title string) {
	todo := Todo{
		Title:      title,
		Completed:  false,
		CreatedAt:  time.Now(),
		FinishedAt: nil,
	}

	*t = append(*t, todo)
}

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (t *Todos) Delete(index int) error {
	todos := *t

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*t = append(todos[:index], todos[index+1:]...)
	return nil
}
