create table users
(
  id serial primary key,
  username varchar(256) unique not null
);