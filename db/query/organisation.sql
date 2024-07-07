-- name: CreateOrganisation :one
insert into organisations (org_id, name, description, user_id)
values ($1, $2, $3, $4)
RETURNING org_id, name, description, user_id;

-- name: GetUserOrgsByID :many
select * from organisations
where user_id = $1;

-- name: GetOrgByID :one
select * from organisations
where org_id = $1;
