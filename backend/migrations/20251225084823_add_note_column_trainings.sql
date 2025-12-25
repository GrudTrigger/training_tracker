-- +goose Up
-- +goose StatementBegin
ALTER TABLE trainings ADD COLUMN note TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
