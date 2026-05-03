package todo_handlers

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Todo struct {
	ID         int
	Title      string
	Completed  bool
	CreatedAt  time.Time
	FinishedAt *time.Time
}

func AddTodo(db *sql.DB, title string) error {
	query := `INSERT INTO todos (title) VALUES ($1)`
	_, err := db.Exec(query, title)
	return err
}

func GetTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query(`SELECT id, title, completed, created_at, finished_at FROM todos ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt, &t.FinishedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func ToggleTodo(db *sql.DB, id int) error {
	query := `
		UPDATE todos 
		SET completed = NOT completed,
		    finished_at = CASE 
		        WHEN completed = FALSE THEN NOW() 
		        ELSE NULL 
		    END
		WHERE id = $1
	`
	_, err := db.Exec(query, id)
	return err
}

func EditTodo(db *sql.DB, id int, title string) error {
	query := `UPDATE todos SET title = $1 WHERE id = $2`
	_, err := db.Exec(query, title, id)
	return err
}

func DeleteTodo(db *sql.DB, id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

func PrintTodos(db *sql.DB) error {
	todos, err := GetTodos(db)
	if err != nil {
		return err
	}

	data := []string{"ID", "Title", "Completed", "CreatedAt", "FinishedAt"}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data)

	for _, t := range todos {
		finishedAt := "—"
		if t.FinishedAt != nil {
			finishedAt = t.FinishedAt.Format("2006-01-02 15:04:05")
		}
		table.Append([]string{
			fmt.Sprintf("%d", t.ID),
			t.Title,
			fmt.Sprintf("%v", t.Completed),
			t.CreatedAt.Format("2006-01-02 15:04:05"),
			finishedAt,
		})
	}
	table.Render()
	return nil
}
