-- +goose Up
CREATE TABLE USERS (
    ID UNIQUEIDENTIFIER DEFAULT NEWID(),
    Email VARCHAR(100) NOT NULL,
    PasswordHash VARCHAR(100) NOT NULL
);

CREATE TABLE CART (
    CUSTOMER_ID CHAR(2),
    FOOD_ITEM_ID CHAR(2),
    QUANTITY INT 
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE USER;
DROP TABLE CART;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
