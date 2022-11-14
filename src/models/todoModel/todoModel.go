package todoModel

import (
	"GinAPI/src/db"
	"gorm.io/gorm"
)

type Todo struct {
	Id        int    `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`
	Completed bool   `json:"completed" gorm:"column:completed"`
	UserId    int    `json:"userId" gorm:"column:userId"`
}

func GetTodos(userId int) ([]Todo, error) {
	var todos []Todo
	err := db.GetORM().Table("todos").Where("userId = ?", userId).Find(&todos).Error
	return todos, err
}

func GetTodoById(id int) (Todo, *gorm.DB) {
	var todos []Todo
	result := db.GetORM().Table("todos").Where("id = ?", id).Find(&todos)
	return todos[0], result
}

func GetTodoByUserId(id int) (Todo, *gorm.DB) {
	var todos []Todo
	result := db.GetORM().Table("todos").Where("userId = ?", id).Find(&todos)
	return todos[0], result
}

func CreateTodo(newData Todo) (Todo, error) {
	err := db.GetORM().Create(&newData).Error

	return newData, err
}

func UpdateTodo(newData Todo) (Todo, error) {
	err := db.GetORM().Save(&newData).Error
	return newData, err
}

func DeleteTodo(newData Todo) *gorm.DB {
	result := db.GetORM().Where("id = ?", newData.Id).Delete(&newData)
	return result
}
