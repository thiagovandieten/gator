-- name: CreateFeed :one

INSERT INTO feeds(name, created_at, updated_at, url, user_id) VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetFeeds :many

SELECT f.name, f.url, u.name AS username FROM feeds AS f
JOIN users AS u ON f.user_id = u.id;

-- name: SearchFeedByURL :one
SELECT * FROM feeds
WHERE url = $1
LIMIT 1;

-- name: MarkFeedFatchedById :exec
UPDATE feeds 
SET last_fetched_at = $1
WHERE id = $2; 

-- name: GetNextFeedToFetch :many
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;