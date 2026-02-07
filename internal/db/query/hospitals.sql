-- name: CreateHospital :one
INSERT INTO hospitals (
    name,
    latitude,
    longitude,
    address,
    contact
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetAllHospitals :many
SELECT * FROM hospitals
ORDER BY name;

-- name: GetHospitalByID :one
SELECT * FROM hospitals
WHERE id = $1;
