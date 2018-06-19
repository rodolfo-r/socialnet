ALTER TABLE comments
  ADD COLUMN IF NOT EXISTS created_at timestamp NOT NULL;
