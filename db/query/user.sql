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
with orga as (
select distinct m2.member_id as member2 
from org_members as m1
join org_members as m2 on m1.org_id = m2.org_id
where m1.member_id = $1 and m2.member_id = $2
)
select id, first_name, last_name, email, phone from users
join orga on orga.member2 = users.id;

