package main

import (
	"log"
	"todo-app/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()
	r.Use(cors.New(CORSConfig()))

	r.POST("/todos/", createTodo)
	r.GET("/todos/", listTodo)
	r.DELETE("/todos/:id", deleteTodo)
	r.PATCH("/todos/:id", updateTodo)

	r.Run(":8082")
}

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
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
	if c.Request.Body == nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Request body is empty"})
		return
	}
	var todo models.Todo

	err := c.BindJSON(&todo)
	if err != nil {
		log.Println(err)
		return
	}

	if todo.Name == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "All fields are required"})
		return
	}

	err = models.CreateTodo(todo)
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

	if todo.Name == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	err := models.UpdateTodoById(id, todo)
	checkErr(err)

	c.JSON(200, gin.H{"message": "Todo updated successfully"})
}
