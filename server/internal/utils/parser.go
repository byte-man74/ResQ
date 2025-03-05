package utils

import "github.com/go-playground/validator/v10"


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



func FormatValidationErrors(err error) []validationErrorResponse {
	var errors []validationErrorResponse

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors = make([]validationErrorResponse, len(validationErrors))
		for i, fieldError := range validationErrors {
			errors[i] = validationErrorResponse{
				Field:   fieldError.Field(),
				Message: msgForTag(fieldError.Tag()),
			}
		}
	}

	return errors
}
