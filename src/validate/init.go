package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"reflect"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Init(e *echo.Echo) {
	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		fieldName := field.Tag.Get("json")
		if fieldName == "=" {
			return ""
		}
		return fieldName
	})

	e.Validator = &CustomValidator{validator: v}
}
