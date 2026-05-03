package main

import (
	"log"
	"os"
	datab "todolist/internal/database"
	th "todolist/internal/handlers"
	router "todolist/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	connStr := os.Getenv("DB_URL")

	db, err := datab.NewDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	r := router.SetupRouter(db)

	// th.AddTodo(db, "Create a database")
	// th.ToggleTodo(db, 2)
	// th.ToggleTodo(db, 1)
	// th.ToggleTodo(db, 3)
	// th.DeleteTodo(db, 1)
	// th.AddTodo(db, "TEST")
	th.PrintTodos(db)

	r.Run(":8080")

	defer db.Close()
}
