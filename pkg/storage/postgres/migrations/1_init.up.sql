CREATE TABLE IF NOT EXISTS users(
  id uuid,
  created_at date,
  updated_at date,
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL,
  PRIMARY KEY(id)
 );

CREATE TABLE IF NOT EXISTS posts(
  id uuid,
  created_at date,
  updated_at date,
  users_id uuid REFERENCES users(id),
  title VARCHAR NOT NULL,
  body VARCHAR NOT NULL
);
