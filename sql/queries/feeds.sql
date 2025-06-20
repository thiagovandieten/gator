-- name: CreateFeed :one

INSERT INTO feeds(name, url, user_id) VALUES(
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeeds :many

SELECT f.name, f.url, u.name AS username FROM feeds AS f
JOIN users AS u ON f.user_id = u.id;

-- name: SearchFeedByURL :one
SELECT * FROM feeds
WHERE url = $1
LIMIT 1;