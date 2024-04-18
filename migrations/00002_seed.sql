-- +goose Up
INSERT INTO FOOD_CATEGORY VALUES ('01', 'Pizza')
INSERT INTO FOOD_SUBCATEGORY VALUES ('01', 'Pepperoni', '01')
INSERT INTO FOOD_ITEM VALUES ('01', 'Pepperoni Pizza', 'Large', 1, 5.99, '01')

INSERT INTO HOLIDAYS_RECOGNIZED VALUES ('01', 'Christmas', '2024-12-31')
INSERT INTO PROMOTION VALUES ('01', 'Default', 'Default', 0.00, NULL, NULL)
INSERT INTO HOLIDAY_PROMOTION VALUES ('01', '01');

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
