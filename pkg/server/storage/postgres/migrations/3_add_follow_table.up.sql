CREATE TABLE IF NOT EXISTS follows(
  id uuid UNIQUE,
  follower_id uuid REFERENCES users(id),
  followee_id uuid REFERENCES users(id)
);

