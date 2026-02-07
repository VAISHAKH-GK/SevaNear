-- +goose Up
-- +goose StatementBegin
ALTER TABLE services 
ADD COLUMN provider TEXT;

ALTER TABLE services 
ADD COLUMN required_docs TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE services DROP COLUMN provider;
ALTER TABLE services DROP COLUMN required_docs;
-- +goose StatementEnd
