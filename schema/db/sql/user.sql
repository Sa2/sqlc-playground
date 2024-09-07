-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  id, name, email, password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserInfos :many
SELECT
  *
FROM
  users u
INNER JOIN
  user_details ud
ON
  u.id = ud.user_id;

