-- migrate:up
create table if not exists users
(
    id          serial
        constraint users_pk
            primary key,
    created_at  timestamp default now(),
    modified_at timestamp default now(),
    deleted_at  timestamp,
    created_by  int,
    modified_by int,
    deleted_by  int,
    first_name  varchar(255),
    last_name   varchar(255),
    email       varchar(255)           not null,
    phone       varchar(20)            not null,
    is_active   boolean   default true not null,
    password    text
);

create unique index users_email_uindex
    on users (email);

create unique index users_phone_uindex
    on users (phone);

-- migrate:down
drop table if exists users;
drop index if exists users_email_uindex;
drop index if exists users_phone_uindex;
