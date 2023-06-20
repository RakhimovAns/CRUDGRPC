CREATE TABLE movies(
  id bigserial primary key ,
  title text not null ,
  genre text not null ,
  created timestamp not null default current_timestamp,
    updated timestamp not null default  current_timestamp
);