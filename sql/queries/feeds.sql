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

-- name: FollowFeed :one
Insert into feed_follows (id,created_at,updated_at,user_id,feed_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *,
(select u.name from users u where user_id = users.id) As user_name,
(select f.name from feeds f where feed_id = feeds.id) As feed_name;

-- name: GetFeeds :many
Select * from feeds;

-- name: ResetFeeds :exec
DELETE from feeds;