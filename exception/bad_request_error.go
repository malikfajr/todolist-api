package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s is not valid email", fe.Field())
	case "boolean":
		return fmt.Sprintf("%s is not valid boolean", fe.Field())
	}

	return fe.Error()
}
