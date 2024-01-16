package main

import (
	"log"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	router := r.Group("/todos")
	{
		router.POST("", createTodo)
		router.GET("", listTodo)
		router.DELETE("/:id", deleteTodo)
		router.PATCH("/:id", updateTodo)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Add your frontend origin here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	r.Use(corsWrapper(c))

	r.Run()
}

func corsWrapper(c *cors.Cors) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Add your frontend origin here
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Next()
	}
}

func listTodo(c *gin.Context) {
	todos, err := models.GetTodos()
	checkErr(err)

	if todos == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(200, gin.H{"data": todos})
	}
}

func createTodo(c *gin.Context) {

	// check if the request body is empty
	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "Request body is empty"})
		return
	}
	var todo models.Todo

	c.BindJSON(&todo)

	if todo.Name == "" || todo.Description == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	err := models.CreateTodo(todo)
	checkErr(err)

	c.JSON(200, gin.H{"message": "Todo task created successfully"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteTodoById(id)
	checkErr(err)

	c.JSON(200, gin.H{"message": "Todo deleted successfully"})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	c.BindJSON(&todo)

	if todo.Name == "" || todo.Description == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	err := models.UpdateTodoById(id, todo)
	checkErr(err)

	c.JSON(200, gin.H{"message": "Todo updated successfully"})
}
