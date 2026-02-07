-- name: CreateServiceType :one
INSERT INTO service_types (name)
VALUES ($1)
ON CONFLICT (name) DO NOTHING
RETURNING *;

-- name: GetAllServiceTypes :many
SELECT * FROM service_types
ORDER BY name;

-- name: GetServiceTypeByID :one
SELECT * FROM service_types
WHERE id = $1;
