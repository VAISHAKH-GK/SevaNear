-- +goose Up
-- +goose StatementBegin
CREATE TABLE hospitals (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    address TEXT,
    contact TEXT
);
CREATE TABLE service_types (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    hospital_id INTEGER NOT NULL,
    service_type_id INTEGER NOT NULL,

    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,

    name TEXT NOT NULL,
    description TEXT,
    timings TEXT,
    eligibility TEXT,
    contact TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_hospital
        FOREIGN KEY (hospital_id)
        REFERENCES hospitals(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_service_type
        FOREIGN KEY (service_type_id)
        REFERENCES service_types(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS service_types;
DROP TABLE IF EXISTS hospitals;
-- +goose StatementEnd
