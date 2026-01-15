-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: GetUser :one
Select *
from users
where name = $1;

-- name: GetUserById :one
Select * from users
where id = $1;

-- name: GetUsers :many
Select * from users;

-- name: ResetUserTable :exec
DELETE  FROM users;

