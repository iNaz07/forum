drop table post_rating;
drop table thread_rating;
drop table posts;
drop table categories;
drop table threads;
drop table sessions;
drop table users;


create table users (
  id         INTEGER primary key AUTOINCREMENT,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         INTEGER primary key AUTOINCREMENT,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table threads (
  id         INTEGER primary key AUTOINCREMENT,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null       
);

create table categories (
  title      varchar(255),
  thread_id  integer references threads(id)
);

create table posts (
  id         INTEGER primary key AUTOINCREMENT,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null
);

create table thread_rating (
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  liked      boolean
);

create table post_rating (
  user_id    integer references users(id),
  post_id    integer references posts(id),
  liked      boolean
);
