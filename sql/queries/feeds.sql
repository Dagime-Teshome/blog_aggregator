-- name: CreateFeed :one
Insert INTO feeds (id,created_at,updated_at,name,url,user_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedById :one
Select * from
feeds
where  id = $1;

-- name: GetFeedByUrl :one
Select * from
feeds
where  url = $1;

-- name: GetFeeds :many
Select * from feeds;

-- name: ResetFeeds :exec
DELETE from feeds;