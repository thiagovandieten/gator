-- name: CreatePost :exec
INSERT INTO posts(created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetPostsByUser :many
WITH user_feeds AS (
    SELECT id FROM feeds WHERE user_id = $1
)
SELECT * 
FROM posts
WHERE feed_id IN (SELECT id FROM user_feeds)
ORDER BY created_at DESC
LIMIT $2;