package todoSvc

import (
	"GinAPI/src/models/todoModel"
	"fmt"
)

func GetTodo(id int) todoModel.Todo {
	todo, result := todoModel.GetTodo(id)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return todo
}

func GetTodos() []todoModel.Todo {
	todos, result := todoModel.GetTodos()

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return todos
}

func CreateTodo(newData todoModel.Todo) {
	result := todoModel.CreateTodo(newData)

	if result.Error != nil {
		fmt.Println("Insert Error:", result.Error)
	}
}

func DeleteTodo(newData todoModel.Todo) {
	result := todoModel.DeleteTodo(newData)
	if result.Error != nil {
		fmt.Println("Delete Error:", result.Error)
	}
}

func UpdateTodo(newData todoModel.Todo) {
	result := todoModel.UpdateTodo(newData)

	if result.Error != nil {
		fmt.Println("Update Error:", result.Error)
	}
}
