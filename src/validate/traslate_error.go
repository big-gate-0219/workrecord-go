package validate

import (
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Param string `json:"param"`
}

type ErrorResponse struct {
	Status string           `json:"status"`
	Errors []Error `json:"errors"`
}

func TranslateError(errs validator.ValidationErrors) []Error {
	result := []Error{}
	for _, err := range errs {
		error := Error{
			Field: err.Field(),
			Type:  err.Tag(),
			Param: err.Param(),
		}
		result = append(result, error)
	}
	return result
}

func CreateError(errorType string, errorField string) Error {
	return Error{Type: errorType, Field: errorField}
}

func  CreateSingleErrors(errorType string, errorField string) []Error {
	return []Error{CreateError(errorType, errorField)}
}

func CreateErrorResponse(errs []Error) ErrorResponse {
	return ErrorResponse{Status: "Error", Errors: errs}
}