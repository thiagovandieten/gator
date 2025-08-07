-- Name: CreatePost :exec
INSERT INTO posts(created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- Name: GetPostsByUser :many
SELECT * 
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2;