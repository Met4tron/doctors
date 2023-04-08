CREATE TABLE IF NOT EXISTS  note (
  note_id SERIAL PRIMARY KEY,
  content text,
  created_at timestamp NOT NULL DEFAULT(NOW()),
  updated_at timestamp
);

ALTER TABLE appointment
  ADD note_id INT;

ALTER TABLE appointment
  ADD CONSTRAINT fk_note FOREIGN KEY(note_id) REFERENCES note(note_id) ON DELETE CASCADE;