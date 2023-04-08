CREATE TABLE patient (
  patient_id SERIAL PRIMARY KEY,
  name text NOT NULL,
  email text NOT NULL,
  cpf varchar NOT NULL,
  birth_date varchar NOT NULL,
  physician_id int NOT NULL,
  created_at timestamp NOT NULL DEFAULT(NOW()),
  updated_at timestamp,
  
  CONSTRAINT fk_physician FOREIGN KEY(physician_id) REFERENCES physician(physician_id) ON DELETE CASCADE
);