CREATE TABLE projects (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(140) NOT NULL,
    description VARCHAR(140),
    repository  VARCHAR(140),
    url         VARCHAR(140),
    status      INTEGER
);

CREATE TABLE tasks (
  id          SERIAL PRIMARY KEY,
  projectId   INTEGER REFERENCES projects (id),
  name        VARCHAR(140) NOT NULL,
  description VARCHAR(140),
  type        INTEGER,
  statusId    INTEGER REFERENCES statuses (id)
);

CREATE TABLE statuses (
  id          SERIAL PRIMARY KEY,
  projectId   INTEGER REFERENCES projects (id),
  name        VARCHAR(140) NOT NULL,
  description VARCHAR(140),
  index       INTEGER
);

CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(140) NOT NULL,
    email       VARCHAR(140),
    phone       VARCHAR(15)
);

CREATE TABLE releases (
    id          SERIAL PRIMARY KEY,
    projectId   INTEGER REFERENCES projects (id),
    version     varchar(140) NOT NULL
);
