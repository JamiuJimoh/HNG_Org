-- name: CreateOrgMember :one
insert into org_members (member_id, org_id, creator_id)
values ($1, $2, $3)
RETURNING member_id, org_id, creator_id;

