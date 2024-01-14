package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

r:=gin.Default()

	router := r.Group("/todo")
	{
		router.POST("/create", addTodo)
		router.GET("/", listTodo)
		router.GET("/delete/:id", deleteTodo)
	}

	r.Run(":8081") 

}

func addTodo(c *gin.Context) {
	c.JSON(200, gin.H{"message": "A new Record Created!"})
}

func listTodo(c *gin.Context) {
	c.JSON(200, gin.H{"message": "All Records"})
}

func deleteTodo(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Deleted!"})
}