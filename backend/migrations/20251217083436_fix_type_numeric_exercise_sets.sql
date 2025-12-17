-- +goose Up
-- +goose StatementBegin
ALTER TABLE exercise_sets
ALTER COLUMN weight TYPE NUMERIC(5,2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE exercise_sets
ALTER COLUMN weight TYPE NUMERIC(3,2);
-- +goose StatementEnd
