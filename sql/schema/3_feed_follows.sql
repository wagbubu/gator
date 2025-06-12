-- +goose Up
CREATE TABLE feed_follows (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE NOT NULL,
  CONSTRAINT unique_feed_follows UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;