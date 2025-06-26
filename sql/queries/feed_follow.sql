-- name: CreateFeedFollow :one

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

-- name: GetFeedFollowsForUser :many

SELECT feeds.name 
FROM feeds 
INNER JOIN feed_follows 
ON feeds.id = feed_follows.feed_id WHERE 
feed_follows.user_id = (SELECT id FROM users WHERE users.name = $1);

-- name: DeleteFeedFollowWithFeedURL :exec

DELETE FROM feed_follows WHERE feed_id = (SELECT id FROM feeds WHERE url = $1) AND feed_follows.user_id = $2;