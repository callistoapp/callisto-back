CREATE TABLE tasks (
    name        varchar(140) NOT NULL,
    id          integer CONSTRAINT taskId PRIMARY KEY,
    projectId   date,
    description varchar(140),
    type        integer,
    status      integer
);

CREATE TABLE projects (
    name        varchar(140) NOT NULL,
    id          integer CONSTRAINT projectId PRIMARY KEY,
    description varchar(140),
    repository  varchar(140),
    url         varchar(140),
    status      integer
);

CREATE TABLE users (
    name        varchar(140) NOT NULL,
    id          integer CONSTRAINT usersId PRIMARY KEY,
    email       varchar(140),
    phone       varchar(15)
);

CREATE TABLE releases (
    version     varchar(140) NOT NULL,
    id          integer CONSTRAINT releasesId PRIMARY KEY
);
