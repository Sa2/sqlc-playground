-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (
  id, name, email, password
) VALUES (
  $1, $2, $3, $4
);

