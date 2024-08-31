
-- +migrate Up

CREATE SCHEMA sqlc;

CREATE TABLE sqlc.users (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE sqlc.users;

DROP SCHEMA sqlc;
