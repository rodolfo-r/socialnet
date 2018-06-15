CREATE TABLE IF NOT EXISTS comments(
  id uuid NOT NULL UNIQUE,
  commenter_id uuid REFERENCES users(id),
  post_id uuid REFERENCES posts(id),
  text VARCHAR NOT NULL
);
