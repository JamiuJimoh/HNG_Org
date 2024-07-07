-- name: CreateUser :one
insert into users (id, first_name, last_name, email, password, phone)
values ($1, $2, $3, $4, $5, $6)
RETURNING id, first_name, last_name, email, phone;

-- name: GetUserByID :one
select id, first_name, last_name, email, phone from users 
where id = $1;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: GetUserInSameOrgByID :one
with org_ids as (
select * from organisations
join users on users.id = user_id
where user_id = $1
)
select id, first_name, last_name, email, password, phone
from org_ids
where user_id = $2;

-- select * 
-- from (select * from organisations
-- join users on users.id = user_id
-- where user_id = $1) as o
-- where o.user_id = $2;

