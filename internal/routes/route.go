package router

import (
	"database/sql"

	th "todolist/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		th.GetTodosHandler(db, c)
	})

	r.POST("/todos", func(c *gin.Context) {
		th.AddTodoHandler(db, c)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		th.DeleteTodoHandler(db, c)
	})

	r.PATCH("todos/:id/toggle", func(c *gin.Context) {
		th.ToggleTodoHandler(db, c)
	})

	return r
}
