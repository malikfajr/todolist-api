package repository

import (
	"context"
	"database/sql"
	"malikfajr/todolist-api/model/domain"
)

type TodoRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Delete(ctx context.Context, tx *sql.Tx, todoId int)
	FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo
}
