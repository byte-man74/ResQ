package utils

import (
	"strings"
	"github.com/go-playground/validator"
)


func FormatValidationErrors(err error) map[string]string {
	if err == nil {
		return nil
	}



	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			// Split the field path on dots and get the last element
			fieldParts := strings.Split(e.Namespace(), ".")
			field := strings.ToLower(fieldParts[len(fieldParts)-1])
			tag := e.Tag()
			if tag != "" {
				errors[field] = tag + " validation failed"
			}
		}
	}

	return errors
}
