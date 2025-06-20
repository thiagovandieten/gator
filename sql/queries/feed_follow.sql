-- name: CreateFeedFollow :many

WITH inserted_rows AS (
    INSERT INTO feed_follows(created_at, updated_at, user_id, feed_id) VALUES(
        $1,
        $2,
        $3,
        $4
    )
    RETURNING *
)

SELECT u.name AS username, f.name AS feedname, i.*
FROM inserted_rows AS i
INNER JOIN users AS u ON i.user_id = u.id
INNER JOIN feeds AS f on i.feed_id = f.id;
