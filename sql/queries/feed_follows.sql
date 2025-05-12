-- name: CreateFeedFollow :one
WITH ff AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1, 
        $2, 
        $3, 
        $4, 
        $5
    )
    RETURNING *
)
SELECT ff.id, ff.created_at, ff.updated_at, ff.user_id, ff.feed_id, 
feeds.name AS feed_name, 
users.name AS user_name 
FROM ff
JOIN feeds
ON feeds.id = ff.feed_id
JOIN users
ON users.id = ff.user_id;

-- name: GetFeedFollowForUser :many
SELECT ff.id, ff.created_at, ff.updated_at, ff.user_id, ff.feed_id, 
feeds.name AS feed_name, 
users.name AS user_name 
FROM feed_follows AS ff
JOIN feeds
ON feeds.id = ff.feed_id
JOIN users
ON users.id = ff.user_id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
WITH f_id AS (
    SELECT id FROM feeds
    WHERE feeds.url = $1
    LIMIT 1
)
DELETE FROM feed_follows
WHERE feed_follows.feed_id IN (
    SELECT id FROM feeds
    WHERE feeds.url = $1
    LIMIT 1
) 
AND feed_follows.user_id = $2;