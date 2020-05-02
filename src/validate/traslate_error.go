package validate

import (
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Type    string `json:"type"`
	Field   string `json:"field"`
}

func TranslateError(errs validator.ValidationErrors) []Error {
	result := []Error{}
	for _, err := range errs {
		error := Error{
			Field: err.Field(),
			Type:  err.Tag(),
		}
		result = append(result, error)
	}
	return result
}

func CreateError(errorType string, errorField string) Error {
	return Error{Type: errorType, Field: errorField}
}
