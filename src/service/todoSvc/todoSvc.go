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
		fmt.Println(result.Error)
	}
}

func DeleteTodo(newData todoModel.Todo) {

}

func UpdateTodo(newData todoModel.Todo) {
	
}
