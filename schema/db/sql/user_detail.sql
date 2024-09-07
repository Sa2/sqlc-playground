-- name: GetUserDetail :many
SELECT
 * 
FROM
 user_details;

-- name: CreateUserDetail :one
INSERT INTO user_details (
  id, user_id, detail_info1, detail_info2
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

