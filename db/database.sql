create table fitnews (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
)

create table hcmus (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
)

create table lichthi (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
)

create table thongbaopkt (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
)

create table users (
  id text primary key,
  name text not null,
  is_subscribed boolean not null default false,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
)