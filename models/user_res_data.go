package models

import "github.com/JamiuJimoh/hngorg/db/sqlc"

type ResUser struct {
	UserId    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type SingleUserResData struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    ResUser `json:"data"`
}

func ResUserFromUser(u User) ResUser {
	return ResUser{
		UserId:    u.UserId,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}

func ResUserFromDBUser(u sqlc.CreateUserRow) ResUser {
	return ResUser{
		UserId:    u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone.String,
	}
}

func ResUserFromSQLUser(u sqlc.User) ResUser {
	return ResUser{
		UserId:    u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone.String,
	}
}

func ResUserFromSQLUserRow(u sqlc.GetUserByIDRow) ResUser {
	return ResUser{
		UserId:    u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone.String,
	}
}

func ResUserFromSQLUserSameOrgRow(u sqlc.GetUserInSameOrgByIDRow) ResUser {
	return ResUser{
		UserId:    u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone.String,
	}
}
