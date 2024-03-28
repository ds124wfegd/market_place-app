CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE mp_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    filepath    varchar(255) not null,
	price       DECIMAL(8,2) not null,
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references mp_lists (id) on delete cascade not null
);

CREATE TABLE mp_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    filepath    varchar(255) not null,
	price       DECIMAL(8,2) not null,
    done        boolean      not null default false
);


CREATE TABLE lists_items
(
    id      serial                                           not null unique,
    item_id int references mp_items (id) on delete cascade not null,
    list_id int references mp_lists (id) on delete cascade not null
);