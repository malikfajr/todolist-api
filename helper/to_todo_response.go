package helper

import (
	"github.com/malikfajr/todolist-api/model/domain"
	"github.com/malikfajr/todolist-api/model/web"
)

func ToTodoResponse(todo domain.Todo) web.TodoResponse {
	return web.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
	}
}
