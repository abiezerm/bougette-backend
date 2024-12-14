package handlers

import (
	"bouguette/common"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ValidateBodyRequest(c echo.Context, payload interface{}) []*common.ValidationError {
	var errors []*common.ValidationError
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(payload)
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		reflected := reflect.ValueOf(payload)

		for _, validationError := range validationErrors {
			field, _ := reflected.Type().FieldByName(validationError.StructField())
			var key string
			key = field.Tag.Get("json")
			condition := validationError.Tag()
			keyToTitleCase := strings.Replace(key, "_", " ", -1)

			errMessage := keyToTitleCase + " field is " + condition

			switch condition {
			case "required":
				errMessage = keyToTitleCase + " is required"
			case "email":
				errMessage = keyToTitleCase + " must be a valid email"
			}

			//fmt.Println(validationError.Field())
			//fmt.Println(validationError.Tag())
			// fmt.Println(key)
			// fmt.Println("----------------")

			currentValidationError := &common.ValidationError{
				Error:     errMessage,
				Key:       key,
				Condition: condition,
			}
			errors = append(errors, currentValidationError)
		}
	}
	return errors
}
