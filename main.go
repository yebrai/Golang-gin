package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: "1", Title: "Learn Go", Status: "In Progress"},
	{ID: "2", Title: "Read a book", Status: "Pending"},
}

func main() {
	r := gin.Default()

	// Obtener todos los TODOs
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	// Crear un nuevo TODO
	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.BindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, newTodo)
		c.JSON(http.StatusCreated, newTodo)
	})

	// Obtener un TODO por ID
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, todo := range todos {
			if todo.ID == id {
				c.JSON(http.StatusOK, todo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
	})

	// Actualizar un TODO por ID
	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTodo Todo
		if err := c.BindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, todo := range todos {
			if todo.ID == id {
				todos[i] = updatedTodo
				c.JSON(http.StatusOK, updatedTodo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
	})

	// Eliminar un TODO por ID
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "TODO deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
	})

	r.Run() // Por defecto en localhost:8080
}
