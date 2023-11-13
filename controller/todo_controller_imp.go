package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/malikfajr/todolist-api/helper"
	"github.com/malikfajr/todolist-api/model/web"
	"github.com/malikfajr/todolist-api/service"

	"github.com/gorilla/mux"
)

type TodoControllerImp struct {
	TodoService service.TodoService
}

// Create implements TodoController.
func (controller *TodoControllerImp) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	todo := web.TodoCreateRequest{}

	if r.Body != nil {
		err := decoder.Decode(&todo)
		helper.PanicIfError(err)
	}

	todoResponse := controller.TodoService.Create(r.Context(), todo)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(webResponse)
	helper.PanicIfError(err)
}

// Delete implements TodoController.
func (controller *TodoControllerImp) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["todoId"]
	todoId, err := strconv.ParseInt(id, 10, 64)
	helper.PanicIfError(err)

	controller.TodoService.Delete(r.Context(), int(todoId))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(webResponse)
	helper.PanicIfError(err)
}

// FindAll implements TodoController.
func (controller *TodoControllerImp) FindAll(w http.ResponseWriter, r *http.Request) {
	todos := controller.TodoService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todos,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(webResponse)
	helper.PanicIfError(err)
}

// FindById implements TodoController.
func (controller *TodoControllerImp) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["todoId"]
	todoId, err := strconv.ParseInt(id, 10, 64)
	helper.PanicIfError(err)

	todo := controller.TodoService.FindById(r.Context(), int(todoId))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todo,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(webResponse)
	helper.PanicIfError(err)
}

// Update implements TodoController.
func (controller *TodoControllerImp) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	todo := web.TodoUpdateRequest{}

	if r.Body != nil {
		err := decoder.Decode(&todo)
		helper.PanicIfError(err)
	}

	id := params["todoId"]
	todoId, err := strconv.ParseInt(id, 10, 64)
	helper.PanicIfError(err)

	todo.ID = int(todoId)

	todoResponse := controller.TodoService.Update(r.Context(), todo)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(webResponse)
	helper.PanicIfError(err)
}

func NewTodoControllerImp(todoService service.TodoService) TodoController {
	return &TodoControllerImp{
		TodoService: todoService,
	}
}
