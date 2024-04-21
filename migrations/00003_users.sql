-- +goose Up

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE USER;
DROP TABLE CART;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
