-- 002_add_indexes.sql
-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_candidates_email ON candidates(email);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_candidates_status ON candidates(status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_candidates_email ON candidates;
-- +goose StatementEnd

-- +goose StatementBegin
DROP INDEX idx_candidates_status ON candidates;
-- +goose StatementEnd