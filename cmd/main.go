package main

import (
	"fmt"
	th "todolist/internal/handlers"
)

func main() {
	todos := th.Todos{}

	todos.Add("Buy milk")
	todos.Add("Read a book")
	fmt.Println(todos)
	todos.Toggle(1)
	// todos.Toggle(1)
	todos.Edit(0, "Buy book")
	fmt.Println(todos)

	todos.Print()
}
