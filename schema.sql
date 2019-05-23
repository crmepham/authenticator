drop database if exists notes;
create database authenticator default charset utf8 default collate utf8_unicode_ci;
use authenticator;

create table user (
  id int(9) not null auto_increment primary key,
  username varchar(250) not null,
  password varchar(250) not null,
  email varchar(250) not null,
  application_id int(9) not null default 0,
  api boolean default false,
  active boolean default false,
  admin boolean default false,
  deleted boolean default false,
  created datetime not null default CURRENT_TIMESTAMP,
  created_by int(9) not null,
  updated datetime not null default CURRENT_TIMESTAMP,
  updated_by int(9) not null
) engine=InnoDB;

create index idx_u_u on user (username);
create index idx_u_e on user (email);

insert into user (username, password, email, api, active, admin, created_by, updated_by)
values ('temp', 'temp', '', true, true, true, 0, 0);

create table application (
  id int(9) not null auto_increment primary key,
  name varchar(250) not null,
  description varchar(1000),
  url varchar(500) not null,
  active boolean default false,
  deleted boolean default false,
  created datetime not null default CURRENT_TIMESTAMP,
  created_by int(9) not null,
  updated datetime not null default CURRENT_TIMESTAMP,
  updated_by int(9) not null
) engine=InnoDB;

create index idx_a_n on application (name);
create index idx_a_u on application (url);