package todo_handlers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
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

func (t *Todos) Toggle(index int) error {
	todos := *t

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	isComplited := todos[index].Completed

	if !isComplited {
		completionTime := time.Now()
		todos[index].FinishedAt = &completionTime
	}

	todos[index].Completed = !isComplited

	return nil
}

func (t *Todos) Edit(index int, title string) error {
	todos := *t

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	todos[index].Title = title

	return nil
}

func (todos Todos) Print() {
	data := []string{"Title", "Completed", "CreatedAt", "FinishedAt"}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data)

	for _, t := range todos {
		finishedAt := "—"
		if t.FinishedAt != nil {
			finishedAt = t.FinishedAt.Format("2006-01-02 15:04:05")
		}
		table.Append([]string{
			t.Title,
			fmt.Sprintf("%v", t.Completed),
			t.CreatedAt.Format("2006-01-02 15:04:05"),
			finishedAt,
		})
	}

	table.Render()
}
