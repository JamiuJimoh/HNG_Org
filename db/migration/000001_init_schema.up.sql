create table users (
   id varchar(36) primary key unique not null,
   first_name varchar(255) not null,
   last_name varchar(255) not null,
   email varchar(255) unique not null,
   password varchar(255) not null,
   phone varchar(255)
);

create table organisations (
   org_id varchar(36) primary key unique not null,
   name varchar(255) not null,
   description varchar(255),
   user_id varchar(36) not null references users(id)
);
