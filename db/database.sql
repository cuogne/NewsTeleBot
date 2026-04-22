-- This is the only user data I store about you
create table users (
  id text primary key,
  name text not null,
  is_subscribed boolean not null default false,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
);

/* 
This is the table to store articles for each category (feed)
If you add new feeds, make sure table names match categories in config/resource.go.
New tables must keep exactly the same structure as the feed tables below; 

          create table <name_of_table_is_equal_to_category> (
            url text primary key,
            title text not null,
            send_at timestamp,
            prompt_token int,
            completion_token int
          );

only the table name should differ.
*/
create table fithcmus (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
);

create table hcmus (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
);

create table lichthipkt (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
);

create table thongbaopkt (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
);

create table ctda (
  url text primary key,
  title text not null,
  send_at timestamp,
  prompt_token int,
  completion_token int
);