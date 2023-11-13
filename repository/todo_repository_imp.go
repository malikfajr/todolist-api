package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/malikfajr/todolist-api/helper"
	"github.com/malikfajr/todolist-api/model/domain"
)

type TodoRepositoryImp struct {
}

// Delete implements TodoRepository.
func (*TodoRepositoryImp) Delete(ctx context.Context, tx *sql.Tx, todoId int) {
	SQL := "DELETE FROM todo WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, todoId)
	if err != nil {
		panic(err)
	}
}

// FindAll implements TodoRepository.
func (*TodoRepositoryImp) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo {
	SQL := "SELECT * FROM todo"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var todos []domain.Todo
	for rows.Next() {
		var todo domain.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}

	return todos
}

// FindById implements TodoRepository.
func (*TodoRepositoryImp) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error) {
	var todo domain.Todo

	SQL := "SELECT * FROM todo WHERE id = $1"
	row := tx.QueryRowContext(ctx, SQL, todoId)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)
	if err != nil {
		return todo, errors.New("todo not found")
	}

	return todo, nil
}

// Save implements TodoRepository.
func (*TodoRepositoryImp) Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "INSERT INTO todo (title, description, is_done) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRowContext(ctx, SQL, todo.Title, todo.Description, todo.IsDone).Scan(&todo.ID)
	helper.PanicIfError(err)

	return todo
}

// Update implements TodoRepository.
func (*TodoRepositoryImp) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "UPDATE todo SET title = $1, description = $2, is_done = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, SQL, todo.Title, todo.Description, todo.IsDone, todo.ID)
	if err != nil {
		panic(err)
	}

	return todo
}

func NewTodoRepositoryImp() TodoRepository {
	return &TodoRepositoryImp{}
}
