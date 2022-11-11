package todoModel

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
	UserId    int `gorm:"column:userId"`
}

const dsn = "treenod:xmflshem@tcp(ec2-13-125-229-201.ap-northeast-2.compute.amazonaws.com:3306)/study"

func getDbConn() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetTodos() ([]Todo, *gorm.DB) {
	db := getDbConn()
	var todos []Todo
	result := db.Table("todos").Find(&todos)

	return todos, result
}

func GetTodo(id int) (Todo, *gorm.DB) {
	var todos []Todo
	db := getDbConn()
	result := db.Table("todos").Where("id = ?", id).Find(&todos)
	return todos[0], result
}

func CreateTodo(newData Todo) *gorm.DB {
	var todos []Todo
	db := getDbConn()
	result := db.Create(&newData)
	if result.Error == nil {
		result := db.Table("todos").Find(&todos)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println(todos)
		}
	} else {
		fmt.Println("Insert Error:", result.Error)
	}

	return result
}

func UpdateTodo(newData Todo) *gorm.DB {
	var todos []Todo
	db := getDbConn()
	result := db.Save(&newData)
	if result.Error == nil {
		result := db.Table("todos").Where("id = ?", newData.Id).Find(&todos)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println(todos)
		}
	} else {

		fmt.Println("Update Error:", result.Error)
	}

	return result
}

func DeleteTodo(newData Todo) *gorm.DB {
	var todos []Todo
	db := getDbConn()
	result := db.Where("id = ?", newData.Id).Delete(&newData)
	if result.Error == nil {
		result := db.Table("todos").Find(&todos)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println(todos)
		}
	} else {
		fmt.Println("Delete Error:", result.Error)
	}

	return result
}
