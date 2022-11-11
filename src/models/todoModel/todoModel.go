package todoModel

import (
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

func GetTodoById(id int) (Todo, *gorm.DB) {
	var todos []Todo
	db := getDbConn()
	result := db.Table("todos").Where("id = ?", id).Find(&todos)
	return todos[0], result
}

func GetTodoByUserId(id int) (Todo, *gorm.DB) {
	var todos []Todo
	db := getDbConn()
	result := db.Table("todos").Where("userId = ?", id).Find(&todos)
	return todos[0], result
}

func CreateTodo(newData Todo) *gorm.DB {
	db := getDbConn()
	result := db.Create(&newData)

	return result
}

func UpdateTodo(newData Todo) *gorm.DB {
	db := getDbConn()
	result := db.Save(&newData)
	return result
}

func DeleteTodo(newData Todo) *gorm.DB {
	db := getDbConn()
	result := db.Where("id = ?", newData.Id).Delete(&newData)
	return result
}
