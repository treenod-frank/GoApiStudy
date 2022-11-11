package todoCtrl

import (
	"GinAPI/src/models/todoModel"
	"GinAPI/src/service/todoSvc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const RET_OK = 1

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetTodos(c *gin.Context) {
	userId := 1
	todos, err := todoModel.GetTodos(userId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret":   RET_OK,
		"todos": todos,
	})
}

func AddTodos(c *gin.Context) {
	title := "new title"
	userId := 1127
	result, err := todoSvc.CreateTodo(title, userId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret":   RET_OK,
		"todos": result,
	})
}

func UpdateTodos(c *gin.Context) {
	userId := 1127

	if q, has := c.GetPostForm("q"); has {
		inputList := strings.Split(q, ",")
		fmt.Println(inputList)
	}

	todos, err := todoModel.GetTodos(userId)

	if err != nil {
		panic(err)
	}

	for index := range todos {
		todos[index].Completed = true
	}

	result, err := todoSvc.UpdateTodo(todos)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret":   RET_OK,
		"todos": result,
	})
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	index, _ := strconv.Atoi(id)
	todo := todoSvc.GetTodoById(index)
	c.JSON(http.StatusOK, gin.H{
		"todos": todo,
	})
}
