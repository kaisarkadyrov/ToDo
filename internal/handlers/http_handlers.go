package todo_handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodosHandler(db *sql.DB, c *gin.Context) {
	todos, err := GetTodos(db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todos)
}

func AddTodoHandler(db *sql.DB, c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	err := AddTodo(db, input.Title)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "created"})
}

func DeleteTodoHandler(db *sql.DB, c *gin.Context) {
	id := c.Param("id")

	intID, _ := strconv.Atoi(id)

	err := DeleteTodo(db, intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "deleted"})
}

func ToggleTodoHandler(db *sql.DB, c *gin.Context) {
	id := c.Param("id")

	intID, _ := strconv.Atoi(id)

	err := ToggleTodo(db, intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "toggled"})
}
