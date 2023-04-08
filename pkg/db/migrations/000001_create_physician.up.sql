CREATE TABLE physician (
  physician_id SERIAL PRIMARY KEY,
  name text NOT NULL,
  email text UNIQUE NOT NULL,
  crm varchar unique NOT NULL,
  created_at timestamp NOT NULL DEFAULT(NOW()),
  updated_at timestamp
);