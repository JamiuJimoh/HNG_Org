package utils

import (
	"encoding/json"
	"regexp"

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
	} else {
		phonePattern := `^(\+?\d{1,3}[-.\s]?)?(\(?\d{3}\)?[-.\s]?)?[\d.\-\s]{7,10}$`
		re := regexp.MustCompile(phonePattern)

		if !re.MatchString(user.Phone) {
			errors = append(errors, ValidatorError{
				Field:   "phone",
				Message: "Invalid phone number",
			})
		}
	}

	if len(user.Email) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "email",
			Message: "Email cannot be null",
		})
	} else {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(emailRegex)
		if !re.MatchString(user.Email) {
			errors = append(errors, ValidatorError{
				Field:   "email",
				Message: "Invalid email",
			})
		}
	}

	if len(user.Password) == 0 {
		errors = append(errors, ValidatorError{
			Field:   "password",
			Message: "Password cannot be null",
		})
	} else if len(user.Password) < 7 {
		errors = append(errors, ValidatorError{
			Field:   "password",
			Message: "Password length must be greater than 6",
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
