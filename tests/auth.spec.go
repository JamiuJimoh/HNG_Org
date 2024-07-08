package tests

//
// import (
// 	"encoding/json"
// 	"testing"
//
// 	"github.com/JamiuJimoh/hngorg/models"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestValidateUserModel(t *testing.T) {
// 	t.Run("test_validate_user_model_empty_first_name", func(t *testing.T) {
// 		user := models.User{
// 			FirstName: "",
// 			LastName:  "Doe",
// 			Phone:     "+1234567890",
// 			Email:     "john.doe@example.com",
// 			Password:  "password123",
// 		}
// 		expectedError := ValidatorErrorsResponse{
// 			Errors: []ValidatorError{
// 				{
// 					Field:   "firstName",
// 					Message: "First name cannot be null",
// 				},
// 			},
// 		}
// 		expectedPayload, _ := json.Marshal(expectedError)
// 		result := ValidateUserModel(user)
// 		assert.Equal(t, expectedPayload, result)
// 	})
//
// 	t.Run("test_validate_user_model_invalid_phone", func(t *testing.T) {
// 		user := models.User{
// 			FirstName: "John",
// 			LastName:  "Doe",
// 			Phone:     "invalid_phone",
// 			Email:     "john.doe@example.com",
// 			Password:  "password123",
// 		}
// 		expectedError := ValidatorErrorsResponse{
// 			Errors: []ValidatorError{
// 				{
// 					Field:   "phone",
// 					Message: "Invalid phone number",
// 				},
// 			},
// 		}
// 		expectedPayload, _ := json.Marshal(expectedError)
// 		result := ValidateUserModel(user)
// 		assert.Equal(t, expectedPayload, result)
// 	})
//
// 	t.Run("test_validate_user_model_valid_user", func(t *testing.T) {
// 		user := models.User{
// 			FirstName: "John",
// 			LastName:  "Doe",
// 			Phone:     "+1234567890",
// 			Email:     "john.doe@example.com",
// 			Password:  "password123",
// 		}
// 		result := ValidateUserModel(user)
// 		assert.Nil(t, result)
// 	})
// }

