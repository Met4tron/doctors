-- name: GetPhysician :one
SELECT * FROM Physician
WHERE physician_id = $1 LIMIT 1;

-- name: ListPhysicians :many
SELECT * FROM Physician
ORDER BY name;

-- name: CreatePhysician :one
INSERT INTO Physician (
  name, email, crm
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeletePhysician :exec
DELETE FROM Physician
WHERE physician_id = $1;
