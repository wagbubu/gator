-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (user_id, feed_id)
  VALUES ($1, $2)
  RETURNING *
)
SELECT
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds
ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users
ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT 
  feed_follows.id, 
  feed_follows.created_at, 
  feed_follows.updated_at, 
  feed_follows.user_id, 
  feed_follows.feed_id, 
  feeds.name AS feed_name, 
  users.name AS user_name 
FROM feed_follows
INNER JOIN feeds
ON feed_follows.feed_id = feeds.id
INNER JOIN users
ON feed_follows.user_id = users.id
WHERE users.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds, users
WHERE feed_follows.feed_id = feeds.id
AND feed_follows.user_id = users.id
AND users.id = @user_id 
AND feeds.id = @feed_id;