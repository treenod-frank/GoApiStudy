package todoCtrl

import (
	"GinAPI/src/models/todoModel"
	"GinAPI/src/service/todoSvc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetTodos(c *gin.Context) {
	todos, _ := todoModel.GetTodos()
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	index, _ := strconv.Atoi(id)
	todo := todoSvc.GetTodo(index)
	c.JSON(http.StatusOK, gin.H{
		"todos": todo,
	})
}
