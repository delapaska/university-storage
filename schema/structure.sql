CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE project_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    directory varchar(255)
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    project_id int references project_lists (id) on delete cascade not null
);