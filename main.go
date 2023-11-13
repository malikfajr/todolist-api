package main

import (
	"net/http"

	"github.com/malikfajr/todolist-api/app"
	"github.com/malikfajr/todolist-api/controller"
	"github.com/malikfajr/todolist-api/exception"
	"github.com/malikfajr/todolist-api/repository"
	"github.com/malikfajr/todolist-api/service"

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
	db := app.NewDb()

	validate := validator.New()
	todoRepository := repository.NewTodoRepositoryImp()
	todoService := service.NewTodoServiceImp(todoRepository, db, validate)
	todoController := controller.NewTodoControllerImp(todoService)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	routerV1 := router.PathPrefix("/api/").Subrouter()
	routerV1.HandleFunc("/todos", todoController.FindAll).Methods("GET")
	routerV1.HandleFunc("/todos", todoController.Create).Methods("POST")
	routerV1.HandleFunc("/todos/{todoId}", todoController.FindById).Methods("GET")
	routerV1.HandleFunc("/todos/{todoId}", todoController.Update).Methods("PUT")
	routerV1.HandleFunc("/todos/{todoId}", todoController.Delete).Methods("DELETE")

	router.Use(exception.ErrorHandler)

	return router
}
