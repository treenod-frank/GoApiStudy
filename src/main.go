package main

import (
	"GinAPI/src/controllers/todoCtrl"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
	UserId    int `gorm:"column:userId"`
}

func GormTest() {
	dsn := "treenod:xmflshem@tcp(ec2-13-125-229-201.ap-northeast-2.compute.amazonaws.com:3306)/study"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var todos []Todo
	db.Table("todos").Where("userId = ?", 1).Find(&todos)
	fmt.Println(todos)

	myTodo := new(Todo)
	myTodo.Title = "frankTodo"
	myTodo.Completed = false
	myTodo.UserId = 1052

	myTodo.Completed = true

}

func main() {
	r := gin.Default()

	r.GET("/ping", todoCtrl.Ping)
	r.GET("/todos", todoCtrl.GetTodos)
	r.GET("/todos/:id", todoCtrl.GetTodo)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
