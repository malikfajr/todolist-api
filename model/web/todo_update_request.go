package web

type TodoUpdateRequest struct {
	ID          int    `validate:"required"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	IsDone      bool   `json:"is_done" validate:"boolean"`
}
