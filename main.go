package main

import (
	"malikfajr/todolist-api/controller"
	"malikfajr/todolist-api/helper"
	"malikfajr/todolist-api/repository"
	"malikfajr/todolist-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	routes := Routes()

	s := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	s.ListenAndServe()
}

func Routes() http.Handler {
	db := helper.NewDb()

	validate := validator.New()
	todoRepository := repository.NewTodoRepositoryImp()
	todoService := service.NewTodoServiceImp(todoRepository, db, validate)
	todoController := controller.NewTodoControllerImp(todoService)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	router.HandleFunc("/todos", todoController.FindAll).Methods("GET")
	router.HandleFunc("/todos", todoController.Create).Methods("POST")
	router.HandleFunc("/todos/{todoId}", todoController.FindById).Methods("GET")
	router.HandleFunc("/todos/{todoId}", todoController.Update).Methods("PUT")
	router.HandleFunc("/todos/{todoId}", todoController.Delete).Methods("DELETE")

	return router
}
