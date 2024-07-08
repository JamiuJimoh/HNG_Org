-- name: CreateOrganisation :one
insert into organisations (org_id, name, description, user_id)
values ($1, $2, $3, $4)
RETURNING org_id, name, description, user_id;

-- name: GetOrgByOrgID :many
select * from organisations
where org_id = $1;


-- name: GetUserOrgsByID :many
select o.org_id, o.name, o.description, o.user_id from org_members as om
join organisations as o on  o.org_id = om.org_id
where member_id = $1;

-- name: GetOrgByID :one
select * from organisations
where org_id = $1;
