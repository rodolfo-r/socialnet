CREATE TABLE IF NOT EXISTS follow(
  id uuid,
  follower_id uuid REFERENCES users(id),
  followee_id uuid REFERENCES users(id)
);

