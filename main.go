package main

import (
	"log"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func main() {

	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	router := r.Group("/todos")
	{
		router.POST("/", createTodo)
		router.GET("/", listTodo)
		router.DELETE("/:id", deleteTodo)
		router.PATCH("/:id", updateTodo)
	}

	r.Run(":8081")

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

func createTodo (c *gin.Context) {

	// check if the request body is empty
	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "Request body is empty"})
		return
	}
	var todo models.Todo
	
	c.BindJSON(&todo)

	if (todo.Name == "" || todo.Description == "" ) {
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

func updateTodo (c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	
	c.BindJSON(&todo)

	if (todo.Name == "" || todo.Description == "" ) {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	err := models.UpdateTodoById(id, todo)
	checkErr(err)

	c.JSON(200, gin.H{"message": "Todo updated successfully"})
}