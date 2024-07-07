// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: organisation.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrganisation = `-- name: CreateOrganisation :one
insert into organisations (org_id, name, description, user_id)
values ($1, $2, $3, $4)
RETURNING org_id, name, description, user_id
`

type CreateOrganisationParams struct {
	OrgID       string
	Name        string
	Description pgtype.Text
	UserID      string
}

func (q *Queries) CreateOrganisation(ctx context.Context, arg CreateOrganisationParams) (Organisation, error) {
	row := q.db.QueryRow(ctx, createOrganisation,
		arg.OrgID,
		arg.Name,
		arg.Description,
		arg.UserID,
	)
	var i Organisation
	err := row.Scan(
		&i.OrgID,
		&i.Name,
		&i.Description,
		&i.UserID,
	)
	return i, err
}

const getOrgByID = `-- name: GetOrgByID :one
select org_id, name, description, user_id from organisations
where org_id = $1
`

func (q *Queries) GetOrgByID(ctx context.Context, orgID string) (Organisation, error) {
	row := q.db.QueryRow(ctx, getOrgByID, orgID)
	var i Organisation
	err := row.Scan(
		&i.OrgID,
		&i.Name,
		&i.Description,
		&i.UserID,
	)
	return i, err
}

const getUserOrgsByID = `-- name: GetUserOrgsByID :many
select org_id, name, description, user_id from organisations
where user_id = $1
`

func (q *Queries) GetUserOrgsByID(ctx context.Context, userID string) ([]Organisation, error) {
	rows, err := q.db.Query(ctx, getUserOrgsByID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organisation
	for rows.Next() {
		var i Organisation
		if err := rows.Scan(
			&i.OrgID,
			&i.Name,
			&i.Description,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}