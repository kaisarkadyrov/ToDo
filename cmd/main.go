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
	todos.Delete(1)
	fmt.Println(todos)

}
