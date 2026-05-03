package main

import (
	"log"
	datab "todolist/internal/database"
	th "todolist/internal/handlers"
)

func main() {
	connStr := "postgres://postgres:k14062007@localhost:5432/todo_app?sslmode=disable"

	db, err := datab.NewDB(connStr)
	if err != nil {
		log.Fatal(err)
	}

	// th.AddTodo(db, "Create a database")
	// th.ToggleTodo(db, 2)
	// th.ToggleTodo(db, 1)
	// th.ToggleTodo(db, 3)
	// th.DeleteTodo(db, 1)
	th.AddTodo(db, "TEST")
	th.PrintTodos(db)

	defer db.Close()
}
