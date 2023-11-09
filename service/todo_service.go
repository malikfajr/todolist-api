package service

import (
	"context"
	"malikfajr/todolist-api/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Update(ctx context.Context, requst web.TodoUpdateRequest) web.TodoResponse
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) web.TodoResponse
	FindAll(ctx context.Context) []web.TodoResponse
}
