CREATE TABLE projects (
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(140) NOT NULL,
  description VARCHAR(140),
  repository  VARCHAR(140),
  url         VARCHAR(140),
  status      INTEGER,
  deleted     INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE statuses (
  id          SERIAL PRIMARY KEY,
  projectId   INTEGER REFERENCES projects (id),
  name        VARCHAR(140) NOT NULL,
  description VARCHAR(140),
  index       INTEGER,
  deleted     INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE tasks (
  id          SERIAL PRIMARY KEY,
  projectId   INTEGER REFERENCES projects (id),
  name        VARCHAR(140) NOT NULL,
  description VARCHAR(140),
  type        INTEGER,
  statusId    INTEGER REFERENCES statuses (id),
  deleted     INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE users (
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(140) NOT NULL,
  email       VARCHAR(140),
  phone       VARCHAR(15),
  deleted     INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE releases (
  id          SERIAL PRIMARY KEY,
  projectId   INTEGER REFERENCES projects (id),
  version     varchar(140) NOT NULL,
  deleted     INTEGER NOT NULL DEFAULT 0
);
