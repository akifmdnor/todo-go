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

	router := r.Group("/todo")
	{
		router.POST("/create", createTodo)
		router.GET("/", listTodo)
		router.GET("/delete/:id", deleteTodo)
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
	var todo models.Todo
	c.BindJSON(&todo)

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
