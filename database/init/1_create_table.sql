create table users
(
  id serial primary key,
  username varchar(256) unique not null,
  created_at date,
  updated_at date,
  deleted_at date
);