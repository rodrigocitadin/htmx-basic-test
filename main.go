package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(ctx *gin.Context) {
		todos := ReadToDoList()
        fmt.Println(todos)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	e.POST("/todos", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		status := ctx.PostForm("status")
		id, _ := CreateToDo(title, status)

        fmt.Println(title)
        fmt.Println(status)
        fmt.Println(id)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":  title,
			"status": status,
			"id":     id,
		})
	})

	e.DELETE("/todos/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteTodo(id)
	})

	e.Run(":8080")
}
