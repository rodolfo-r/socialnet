CREATE TABLE IF NOT EXISTS users(
  id uuid UNIQUE,
  created_at timestamp,
  updated_at timestamp,
  username VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL,
  PRIMARY KEY(id)
 );

CREATE TABLE IF NOT EXISTS posts(
  id uuid UNIQUE,
  created_at timestamp,
  updated_at timestamp,
  users_id uuid REFERENCES users(id),
  title VARCHAR NOT NULL,
  body VARCHAR NOT NULL
);
