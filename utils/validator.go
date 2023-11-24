package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)
func CustomErrorMessage(err validator.ValidationErrors) map[string][]string {
	var message string
	var errors []string
	messages := make(map[string][]string)
	for _, e := range err {
		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required; ", e.Field())
			errors = append(errors, message)
			messages[e.Field()]=errors
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters; ", e.Field(), e.Param())
			errors = append(errors, message)
			messages[e.Field()]=errors
		case "max":
			message = fmt.Sprintf("%s cannot be longer than %s characters; ", e.Field(), e.Param())
			errors = append(errors, message)
			messages[e.Field()]=errors
		case "email":
			message = fmt.Sprintf("%s must be a valid email address; ", e.Field())
			errors = append(errors, message)
			messages[e.Field()]=errors
		default:
			message = fmt.Sprintf("%s is not valid; ", e.Field())
			errors = append(errors, message)
			messages[e.Field()]=errors
		}
	}
	return messages
}