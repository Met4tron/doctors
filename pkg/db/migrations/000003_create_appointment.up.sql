CREATE TABLE appointment (
  appointment_id SERIAL PRIMARY KEY,
  start_date timestamp NOT NULL,
  end_date timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT(NOW()),
  updated_at timestamp,

  physician_id int NOT NULL,
  patient_id int NOT NULL,
  CONSTRAINT fk_physician FOREIGN KEY(physician_id) REFERENCES physician(physician_id) ON DELETE CASCADE,
  CONSTRAINT fk_patient FOREIGN KEY(patient_id) REFERENCES patient(patient_id) ON DELETE CASCADE
);