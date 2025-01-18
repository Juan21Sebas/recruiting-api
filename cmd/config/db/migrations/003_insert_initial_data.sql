-- 003_insert_initial_data.sql
-- +goose Up
-- +goose StatementBegin
INSERT INTO candidates (name, email, gender, salary_expected, phone, experience_years) VALUES
    ('Ana García', 'ana.garcia@email.com', 'Female', 45000.00, '+1234567890', 3),
    ('Juan Pérez', 'juan.perez@email.com', 'Male', 52000.00, '+1234567891', 5),
    ('María Rodríguez', 'maria.rodriguez@email.com', 'Female', 48000.00, '+1234567892', 4),
    ('Carlos López', 'carlos.lopez@email.com', 'Male', 55000.00, '+1234567893', 6),
    ('Laura Torres', 'laura.torres@email.com', 'Female', 50000.00, '+1234567894', 4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM candidates WHERE email IN (
    'ana.garcia@email.com',
    'juan.perez@email.com',
    'maria.rodriguez@email.com',
    'carlos.lopez@email.com',
    'laura.torres@email.com'
);
-- +goose StatementEnd