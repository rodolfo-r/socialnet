CREATE TABLE IF NOT EXISTS likes(
  id uuid UNIQUE,
  post_id uuid REFERENCES posts(id),
  liker_id uuid REFERENCES users(id)
);
