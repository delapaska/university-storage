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
    title       varchar(255) not null
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    project_id int references project_lists (id) on delete cascade not null
);

CREATE TABLE project_folders
(
    id          serial       not null unique,
    project_id  int          not null references project_lists (id) on delete cascade,
    folder_name varchar(255) not null
);

CREATE TABLE project_tokens
(
    id          serial       not null unique,
    project_id  int          not null references project_lists (id) on delete cascade,
    token       varchar(255) not null unique,
    created_at  timestamp    not null default now()
);

CREATE TABLE files
(
    id          serial       not null unique,
    folder_id   int          not null references project_folders (id) on delete cascade,
    filename    varchar(255) not null,
    filepath    varchar(255) not null
);