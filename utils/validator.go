package utils

import (
	"encoding/json"

	"github.com/JamiuJimoh/hngorg/models"
)

type ValidatorError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidatorErrorsResponse struct {
	Errors []ValidatorError `json:"errors"`
}

func ValidateUserModel(user models.User) []byte {
	var errors []ValidatorError
	// if len(user.UserId) == 0 {
	// 	// todo: check unique user_id
	// 	errors = append(errors, ValidatorError{
	// 		Field:   "userId",
	// 		Message: "invalid user id",
	// 	})
	//
	// }
	if len(user.FirstName) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "firstName",
			Message: "First name cannot be null",
		})
	}

	if len(user.LastName) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "lastName",
			Message: "Last name cannot be null",
		})
	}
	if len(user.Phone) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "phone",
			Message: "Phone number cannot be null",
		})
	}
	if len(user.Email) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "email",
			Message: "Email cannot be null",
		})
	}
	if len(user.Password) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "password",
			Message: "Password cannot be null",
		})
	} else if len(user.Password) < 8 {
		errors = append(errors, ValidatorError{
			Field:   "password",
			Message: "Password length must be greater than 7",
		})
	}
	if len(errors) > 0 {
		res := ValidatorErrorsResponse{Errors: errors}

		payload, _ := json.Marshal(res)
		return payload
	}
	return nil
}

func ValidateOrgName(name string) []byte {
	var errors []ValidatorError

	if len(name) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "name",
			Message: "name must be a valid string",
		})
	}
	if len(errors) > 0 {
		res := ValidatorErrorsResponse{Errors: errors}

		payload, _ := json.Marshal(res)
		return payload
	}
	return nil
}
