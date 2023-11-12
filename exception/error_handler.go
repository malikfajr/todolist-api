package exception

import (
	"encoding/json"
	"malikfajr/todolist-api/helper"
	"malikfajr/todolist-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if badRequestError(w, r, err) {
					return
				}

				if notFoundError(w, r, err) {
					return
				}

				internalServerErr(w, r, err)
			}
		}()

		h.ServeHTTP(w, r)
	})
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		code := http.StatusBadRequest

		errors := []map[string]string{}

		for _, err := range exception {
			errors = append(errors, map[string]string{err.Field(): msgForTag(err)})
			// errors[err.Field()] = msgForTag(err)
		}

		errorReponse := ErrorResponse(w, r, code, map[string]interface{}{"errors": errors})

		e := json.NewEncoder(w).Encode(errorReponse)
		helper.PanicIfError(e)
		return true

	} else {
		return false
	}

}
func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {

		code := http.StatusNotFound

		errorReponse := ErrorResponse(w, r, code, exception.Error)

		e := json.NewEncoder(w).Encode(errorReponse)
		helper.PanicIfError(e)

		return true
	} else {
		return false
	}

}

func internalServerErr(w http.ResponseWriter, r *http.Request, err interface{}) {
	code := http.StatusInternalServerError

	_ = err.(error)

	errorResponse := ErrorResponse(w, r, code, nil)

	e := json.NewEncoder(w).Encode(errorResponse)
	helper.PanicIfError(e)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, code int, err interface{}) web.ErrorResponse {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	webResponse := web.ErrorResponse{
		Code:   code,
		Status: http.StatusText(code),
		Errors: err,
	}

	return webResponse
}
