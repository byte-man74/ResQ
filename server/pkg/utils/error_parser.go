package utils

import (
	"reflect"
	"github.com/go-playground/validator/v10"
)

type validationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	default:
		return "Invalid value"
	}
}

// Function to get JSON tag name
func getJSONFieldName(structType reflect.Type, fieldName string) string {
	field, found := structType.FieldByName(fieldName)
	if !found {
		return fieldName // Default to struct field name if no JSON tag
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return fieldName
	}

	return jsonTag
}

func FormatValidationErrors(err error, model interface{}) []validationErrorResponse {
	var errors []validationErrorResponse

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		structType := reflect.TypeOf(model)

		errors = make([]validationErrorResponse, len(validationErrors))
		for i, fieldError := range validationErrors {
			jsonField := getJSONFieldName(structType, fieldError.Field())

			errors[i] = validationErrorResponse{
				Field:   jsonField,
				Message: msgForTag(fieldError.Tag()),
			}
		}
	}

	return errors
}
