package models

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    string `json:"userId"`    // must be unique
	FirstName string `json:"firstName"` // must not be null
	LastName  string `json:"lastName"`  // must not be null
	Email     string `json:"email"`     // must be unique and must not be null
	Password  string `json:"password"`  // must not be null
	Phone     string `json:"phone"`
}

func (u *User) SanitizeUser() error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.UserId = id.String()
	u.Password = string(hash)
	return nil
}

func (u *User) CreateOrgFromUser() (*Organisation, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	org := Organisation{
		OrgId:       id.String(),
		Name:        fmt.Sprintf("%s's Organisation", u.FirstName),
		Description: "",
	}

	return &org, nil
}
