-- name: CreateService :one
INSERT INTO services (
    hospital_id,
    service_type_id,
    name,
    description,
    timings,
    eligibility,
    contact,
    latitude,
    longitude
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetAllServices :many
SELECT
    s.id,
    s.name,
    s.description,
    s.timings,
    s.eligibility,
    s.contact,
    s.latitude,
    s.longitude,
    s.created_at,

    h.id AS hospital_id,
    h.name AS hospital_name,

    st.id AS service_type_id,
    st.name AS service_type
FROM services s
JOIN hospitals h ON s.hospital_id = h.id
JOIN service_types st ON s.service_type_id = st.id
ORDER BY s.created_at DESC;

-- name: GetServiceByID :one
SELECT
    s.id,
    s.name,
    s.description,
    s.timings,
    s.eligibility,
    s.contact,
    s.latitude,
    s.longitude,
    s.created_at,

    h.id AS hospital_id,
    h.name AS hospital_name,

    st.id AS service_type_id,
    st.name AS service_type
FROM services s
JOIN hospitals h ON s.hospital_id = h.id
JOIN service_types st ON s.service_type_id = st.id
WHERE s.id = $1;

-- name: GetServicesNearLocation :many
SELECT
    s.*,
    (
        6371 * acos(
            cos(radians($1)) * cos(radians(s.latitude)) *
            cos(radians(s.longitude) - radians($2)) +
            sin(radians($1)) * sin(radians(s.latitude))
        )
    ) AS distance_km
FROM services s
HAVING (
    6371 * acos(
        cos(radians($1)) * cos(radians(s.latitude)) *
        cos(radians(s.longitude) - radians($2)) +
        sin(radians($1)) * sin(radians(s.latitude))
    )
) < $3
ORDER BY distance_km;

-- name: GetServicesByHospitalID :many
SELECT * FROM services
WHERE hospital_id = $1
ORDER BY created_at DESC;

-- name: GetServicesByServiceTypeID :many
SELECT * FROM services
WHERE service_type_id = $1
ORDER BY created_at DESC;
