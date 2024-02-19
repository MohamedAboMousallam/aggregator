-- +goose Up

ALTER TABLE feeds ADD COLUMN last_fetched_at timestamp with time zone;

-- +goose Down
ALTER TABLE feeds Drop COLUMN last_fetched_at;
