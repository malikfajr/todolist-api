package service

import (
	"context"
	"database/sql"

	"github.com/malikfajr/todolist-api/exception"
	"github.com/malikfajr/todolist-api/helper"
	"github.com/malikfajr/todolist-api/model/domain"
	"github.com/malikfajr/todolist-api/model/web"
	"github.com/malikfajr/todolist-api/repository"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImp struct {
	TodoRepository repository.TodoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

// Create implements TodoService.
func (service *TodoServiceImp) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{
		Title:       request.Title,
		Description: request.Description,
		IsDone:      request.IsDone,
	}

	todo = service.TodoRepository.Save(ctx, tx, todo)

	return helper.ToTodoResponse(todo)

}

// Delete implements TodoService.
func (service *TodoServiceImp) Delete(ctx context.Context, todoId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoRepository.Delete(ctx, tx, todo.ID)
}

// FindAll implements TodoService.
func (service *TodoServiceImp) FindAll(ctx context.Context) []web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todos := service.TodoRepository.FindAll(ctx, tx)

	var todoResponses = []web.TodoResponse{}
	for _, todo := range todos {
		todoResponses = append(todoResponses, helper.ToTodoResponse(todo))
	}

	return todoResponses
}

// FindById implements TodoService.
func (service *TodoServiceImp) FindById(ctx context.Context, todoId int) web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTodoResponse(todo)
}

// Update implements TodoService.
func (service *TodoServiceImp) Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todo.Title = request.Title
	todo.Description = request.Description
	todo.IsDone = request.IsDone

	todo = service.TodoRepository.Update(ctx, tx, todo)

	return helper.ToTodoResponse(todo)
}

func NewTodoServiceImp(todoRepository repository.TodoRepository, db *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImp{
		TodoRepository: todoRepository,
		DB:             db,
		Validate:       validate,
	}
}
