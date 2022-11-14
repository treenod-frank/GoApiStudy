package todoCtrl

import (
	"GinAPI/src/models/todoModel"
	"GinAPI/src/pkpk"
	"GinAPI/src/service/todoSvc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const RET_OK = 1

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetTodos(c *gin.Context) {
	q := pkpk.GetRequestParam(c)

	userId, _ := strconv.Atoi(q[0])
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
	q := pkpk.GetRequestParam(c)
	title := q[0]
	userId, _ := strconv.Atoi(q[1])
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
	q := pkpk.GetRequestParam(c)

	id, _ := strconv.Atoi(q[0])
	userId, _ := strconv.Atoi(q[1])
	completed, _ := strconv.ParseBool(q[2])

	todos, err := todoModel.GetTodoById(id)

	if err != nil {
		panic(err)
	}

	todos.Completed = completed
	todos.UserId = userId

	result, err2 := todoSvc.UpdateTodo(todos)

	if err2 != nil {
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

func AddTodo(c *gin.Context) {
	q := pkpk.GetRequestParam(c)
	title := q[0]
	userId, _ := strconv.Atoi(q[1])
	_, err := todoSvc.CreateTodo(title, userId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret": RET_OK,
	})

}
