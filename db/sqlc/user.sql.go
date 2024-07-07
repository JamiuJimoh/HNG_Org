// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
insert into users (id, first_name, last_name, email, password, phone)
values ($1, $2, $3, $4, $5, $6)
RETURNING id, first_name, last_name, email, phone
`

type CreateUserParams struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     pgtype.Text
}

type CreateUserRow struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Phone,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id, first_name, last_name, email, password, phone from users where email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Phone,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
select id, first_name, last_name, email, phone from users 
where id = $1
`

type GetUserByIDRow struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     pgtype.Text
}

func (q *Queries) GetUserByID(ctx context.Context, id string) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
	)
	return i, err
}

const getUserInSameOrgByID = `-- name: GetUserInSameOrgByID :one
with org_ids as (
select org_id, name, description, user_id, id, first_name, last_name, email, password, phone from organisations
join users on users.id = user_id
where user_id = $1
)
select id, first_name, last_name, email, password, phone
from org_ids
where user_id = $2
`

type GetUserInSameOrgByIDParams struct {
	UserID   string
	UserID_2 string
}

type GetUserInSameOrgByIDRow struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     pgtype.Text
}

func (q *Queries) GetUserInSameOrgByID(ctx context.Context, arg GetUserInSameOrgByIDParams) (GetUserInSameOrgByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserInSameOrgByID, arg.UserID, arg.UserID_2)
	var i GetUserInSameOrgByIDRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Phone,
	)
	return i, err
}