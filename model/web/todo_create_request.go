package web

type TodoCreateRequest struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	IsDone      bool   `json:"is_done" validate:"boolean"`
}
