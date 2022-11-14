package todoSvc

import (
	"GinAPI/src/models/todoModel"
	"fmt"
)

func GetTodoById(id int) todoModel.Todo {
	todo, result := todoModel.GetTodoById(id)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return todo
}

func GetTodoByUserId(id int) todoModel.Todo {
	todo, result := todoModel.GetTodoByUserId(id)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return todo
}

func GetTodos(userId int) []todoModel.Todo {
	todos, err := todoModel.GetTodos(userId)

	if err != nil {
		panic(err)
	}

	return todos
}

func CreateTodo(title string, userId int) (todoModel.Todo, error) {
	newData := todoModel.Todo{
		Title:  title,
		UserId: userId,
	}

	newData, err := todoModel.CreateTodo(newData)

	return newData, err
}

func DeleteTodo(newData todoModel.Todo) {
	result := todoModel.DeleteTodo(newData)
	if result.Error != nil {
		fmt.Println("Delete Error:", result.Error)
	}
}

func UpdateTodo(newData todoModel.Todo) (todoModel.Todo, error) {
	result, err := todoModel.UpdateTodo(newData)

	return result, err
}
