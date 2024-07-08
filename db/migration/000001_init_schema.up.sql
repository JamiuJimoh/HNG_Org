create table users (
   id varchar(36) primary key, 
   first_name varchar(255) not null check (first_name <> ''),
   last_name varchar(255) not null check (last_name <> ''),
   email varchar(255) unique not null check (last_name <> ''),
   password varchar(255) not null,
   phone varchar(255)
);

create table organisations (
   org_id varchar(36) primary key,
   name varchar(255) not null,
   description varchar(255),
   user_id varchar(36) not null references users(id)
   check (name <> '')
);

create table org_members (
   id serial primary key,
   member_id varchar(36) not null references users(id),
   org_id varchar(36) not null references organisations(org_id) on delete cascade,
   creator_id varchar(36) not null references users(id) on delete cascade,
   unique (member_id, org_id)
);
