-- Active: 1658457037724@@127.0.0.1@3306@jabin
create table users (
    id          SERIAL PRIMARY KEY,
    uuid        VARCHAR(64) not NULL UNIQUE,
    name        VARCHAR(255),
    email       VARCHAR(255) not NULL UNIQUE, 
    password    VARCHAR(255) not NULL,
    created_at  TIMESTAMP not NULL
);


CREATE Table sessions(
    id          SERIAL PRIMARY KEY, 
    uuid        VARCHAR(64) not null UNIQUE,
    email       VARCHAR(255),
    user_id     INTEGER REFERENCES users(id),
    created_at  TIMESTAMP not NULL
);


CREATE TABLE threads(
    id         serial primary key,
    uuid       varchar(64) not null unique,
    topic      TEXT,
    user_id    integer references users(id),  --创建该贴的用户id
    created_at timestamp not null
);


create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),    --该评论所属的用户id
  thread_id  integer references threads(id),  --该评论属于哪个贴
  created_at timestamp not null
);