-- +goose Up
CREATE TABLE VENDOR(
    VENDOR_ID CHAR(2), 
    VENDOR_NAME VARCHAR(35) NOT NULL, 
    PRIMARY KEY (VENDOR_ID)
);

CREATE TABLE INVENTORY_CATEGORY(
    CATEGORY_ID CHAR(2), 
    CATEGORY_NAME VARCHAR(50) NOT NULL, 
    PRIMARY KEY (CATEGORY_ID)
);

CREATE TABLE INVENTORY_SUBCATEGORY(
    SUBCATEGORY_ID CHAR(2), CATEGORY_ID CHAR(2), 
    DESCRIPTION VARCHAR(50), PRIMARY KEY (SUBCATEGORY_ID), 
    FOREIGN KEY (CATEGORY_ID) REFERENCES INVENTORY_CATEGORY (CATEGORY_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE FOOD_VENDOR(
    VENDOR_ID CHAR(2),
    PRIMARY KEY (VENDOR_ID),
    FOREIGN KEY (VENDOR_ID) REFERENCES VENDOR (VENDOR_ID) ON DELETE CASCADE
);

CREATE TABLE NONFOOD_VENDOR(
    VENDOR_ID CHAR(2),
    PRIMARY KEY (VENDOR_ID), FOREIGN KEY (VENDOR_ID) REFERENCES VENDOR (VENDOR_ID) ON DELETE CASCADE
);

CREATE TABLE SUPPLY(
    SUPPLY_ID CHAR(2), 
    SUPPLY_NAME VARCHAR(35) NOT NULL, 
    QUANTITY CHAR(10) NOT NULL, 
    REORDER_THRESHOLD CHAR(3) NOT NULL, 
    SUBCATEGORY_ID CHAR(2), 
    VENDOR_ID CHAR(2), 
    PRIMARY KEY (SUPPLY_ID), 
    FOREIGN KEY (SUBCATEGORY_ID) REFERENCES INVENTORY_SUBCATEGORY (SUBCATEGORY_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (VENDOR_ID) REFERENCES VENDOR (VENDOR_ID) ON UPDATE CASCADE ON DELETE CASCADE
); 

CREATE TABLE PROMOTION(
    PROMOTION_ID CHAR(2), 
    PROMOTION_NAME VARCHAR(35) NOT NULL, 
    PROMOTION_DESCRIPTION VARCHAR(100), 
    DISCOUNT DECIMAL(5, 2) DEFAULT 0.00 NOT NULL, 
    START_DATE DATE, 
    END_DATE DATE, 
    PRIMARY KEY (PROMOTION_ID)
); 

CREATE TABLE HOLIDAYS_RECOGNIZED(
    HOLIDAY_ID CHAR(2), 
    HOLIDAY_NAME VARCHAR(35) NOT NULL, 
    HOLIDAY_DATE DATE NOT NULL, 
    PRIMARY KEY (HOLIDAY_ID)
); 

CREATE TABLE HOLIDAY_PROMOTION(
    PROMOTION_ID CHAR(2), 
    HOLIDAY_ID CHAR(2), 
    PRIMARY KEY (PROMOTION_ID), 
    FOREIGN KEY (PROMOTION_ID) REFERENCES PROMOTION (PROMOTION_ID) ON UPDATE CASCADE ON DELETE CASCADE, 
    FOREIGN KEY (HOLIDAY_ID) REFERENCES HOLIDAYS_RECOGNIZED (HOLIDAY_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE MANAGERS_SPECIAL(
    PROMOTION_ID CHAR(2), 
    PRIMARY KEY (PROMOTION_ID), 
    FOREIGN KEY (PROMOTION_ID) REFERENCES PROMOTION (PROMOTION_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE FOOD_CATEGORY(
    CATEGORY_ID CHAR(2), 
    CATEGORY_NAME VARCHAR(35) NOT NULL, 
    PRIMARY KEY (CATEGORY_ID)
); 

CREATE TABLE FOOD_SUBCATEGORY(
    SUBCATEGORY_ID CHAR(2), 
    DESCRIPTION VARCHAR(50) NOT NULL, 
    CATEGORY_ID CHAR(2), 
    PRIMARY KEY (SUBCATEGORY_ID), 
    FOREIGN KEY (CATEGORY_ID) REFERENCES FOOD_CATEGORY (CATEGORY_ID) ON UPDATE CASCADE ON DELETE CASCADE
);  

CREATE TABLE FOOD_ITEM(
    FOOD_ITEM_ID CHAR(2), 
    FOOD_NAME VARCHAR(35) NOT NULL, 
    FOOD_SIZE CHAR(6) NOT NULL, 
    QUANTITY NUMERIC(4) NOT NULL,
    PRICE DECIMAL(5, 2) NOT NULL,
    SUBCATEGORY_ID CHAR(2),
    PRIMARY KEY (FOOD_ITEM_ID), FOREIGN KEY (SUBCATEGORY_ID) REFERENCES FOOD_SUBCATEGORY (SUBCATEGORY_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE PROMOTION_FOOD_BRIDGE(
    PROMOTION_ID CHAR(2), 
    FOOD_ITEM_ID CHAR(2), 
    ACTIVE BIT NOT NULL, 
    PRIMARY KEY (PROMOTION_ID, FOOD_ITEM_ID), 
    FOREIGN KEY (PROMOTION_ID) REFERENCES PROMOTION (PROMOTION_ID) ON UPDATE CASCADE ON DELETE CASCADE, 
    FOREIGN KEY (FOOD_ITEM_ID) REFERENCES FOOD_ITEM (FOOD_ITEM_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE ROLE(
    ROLE_ID CHAR(2),
    ROLE_NAME VARCHAR(50) NOT NULL,
    IS_HOURLY BIT NOT NULL,
    WAGE DECIMAL(10, 2),
    PRIMARY KEY (ROLE_ID)
);

CREATE TABLE EMPLOYEE(
    EMPLOYEE_ID CHAR(2),
    SUPERVISOR_ID CHAR(2) NULL,
    ROLE_ID CHAR(2) NOT NULL,
    F_NAME VARCHAR(50) NOT NULL,
    L_NAME VARCHAR(50) NOT NULL,
    STREET VARCHAR(50) NOT NULL,
    CITY VARCHAR(30) NOT NULL,
    STATE VARCHAR(20) NOT NULL,
    ZIP VARCHAR(10) NOT NULL,
    DATE_OF_BIRTH DATE NOT NULL,
    CURRENT_WAGE DECIMAL(10, 2),
    PRIMARY KEY (EMPLOYEE_ID),
    FOREIGN KEY (SUPERVISOR_ID) REFERENCES EMPLOYEE(EMPLOYEE_ID) ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY (ROLE_ID) REFERENCES ROLE(ROLE_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE SHIFT(
    SHIFT_ID CHAR(2),
    START_TIME TIME,
    END_TIME TIME,
    PRIMARY KEY (SHIFT_ID)
);

CREATE TABLE EMPLOYEE_SHIFT_BRIDGE(
    EMPLOYEE_ID CHAR(2),
    SHIFT_ID CHAR(2),
    HOURS_WORKED DECIMAL(5, 2) NOT NULL,
    CLOCK_IN_DATETIME DATETIME NOT NULL,
    CLOCK_OUT_DATETIME DATETIME NOT NULL,
    PRIMARY KEY (EMPLOYEE_ID, SHIFT_ID),
    FOREIGN KEY (EMPLOYEE_ID) REFERENCES EMPLOYEE(EMPLOYEE_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (SHIFT_ID) REFERENCES SHIFT(SHIFT_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE JOB_HISTORY(
    EMPLOYEE_ID CHAR(2),
    ROLE_ID CHAR(2),
    START_DATE DATE,
    END_DATE DATE,
    WAGE DECIMAL(10, 2),
    PRIMARY KEY (EMPLOYEE_ID),
    FOREIGN KEY (EMPLOYEE_ID) REFERENCES EMPLOYEE(EMPLOYEE_ID),
    FOREIGN KEY (ROLE_ID) REFERENCES ROLE(ROLE_ID)
);

CREATE TABLE CHEF(
    EMPLOYEE_ID CHAR(2),
    SPECIALTY VARCHAR(50),
    PRIMARY KEY (EMPLOYEE_ID),
    FOREIGN KEY (EMPLOYEE_ID) REFERENCES EMPLOYEE(EMPLOYEE_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE CERTIFICATION(
    CERTIFICATION_ID CHAR(2),
    NAME VARCHAR(50) NOT NULL,
    TYPE VARCHAR(50) NOT NULL,
    PRIMARY KEY (CERTIFICATION_ID)
);

CREATE TABLE CHEF_CERT_BRIDGE(
    CERTIFICATION_ID CHAR(2),
    EMPLOYEE_ID CHAR(2),
    VALID_TO DATE,
    PRIMARY KEY (CERTIFICATION_ID, EMPLOYEE_ID),
    FOREIGN KEY (CERTIFICATION_ID) REFERENCES CERTIFICATION(CERTIFICATION_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (EMPLOYEE_ID) REFERENCES EMPLOYEE(EMPLOYEE_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE MANAGER(
    EMPLOYEE_ID CHAR(2),
    TITLE VARCHAR(50),
    PRIMARY KEY (EMPLOYEE_ID)
);

CREATE TABLE DEGREE(
    DEGREE_ID CHAR(2),
    NAME VARCHAR(50) NOT NULL,
    TYPE VARCHAR(50) NOT NULL,
    COLLEGE_NAME VARCHAR(60) NOT NULL,
    GRADUATION_DATE DATE NOT NULL,
    PRIMARY KEY (DEGREE_ID)
);

CREATE TABLE MANAGER_DEGREE_BRIDGE(
    DEGREE_ID CHAR(2),
    EMPLOYEE_ID CHAR(2),
    PRIMARY KEY (DEGREE_ID, EMPLOYEE_ID),
    FOREIGN KEY (DEGREE_ID) REFERENCES DEGREE(DEGREE_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (EMPLOYEE_ID) REFERENCES EMPLOYEE(EMPLOYEE_ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE CLUB_CARD(
    CLUB_CARD_ID CHAR(2),
    POINTS INT NOT NULL DEFAULT 0,
    PRIMARY KEY (CLUB_CARD_ID)
);

CREATE TABLE USERS (
    ID UNIQUEIDENTIFIER DEFAULT NEWID(),
    Email VARCHAR(100) NOT NULL,
    PasswordHash VARCHAR(100) NOT NULL
    PRIMARY KEY (ID)
);

CREATE TABLE CUSTOMER(
    CUSTOMER_ID CHAR(2),
    F_NAME VARCHAR(50) NOT NULL,
    L_NAME VARCHAR(50) NOT NULL,
    STREET VARCHAR(50) NOT NULL,
    CITY VARCHAR(30) NOT NULL,
    STATE VARCHAR(20) NOT NULL,
    ZIP VARCHAR(10) NOT NULL,
    DATE_OF_BIRTH DATE NOT NULL,
    CLUB_CARD_ID CHAR(2),
    USER_ID UNIQUEIDENTIFIER NULL,
    PRIMARY KEY (CUSTOMER_ID),
    FOREIGN KEY (CLUB_CARD_ID) REFERENCES CLUB_CARD(CLUB_CARD_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (USER_ID) REFERENCES USERS(ID) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE CART (
    CUSTOMER_ID CHAR(2),
    FOOD_ITEM_ID CHAR(2),
    QUANTITY INT 
    PRIMARY KEY (CUSTOMER_ID, FOOD_ITEM_ID),
    FOREIGN KEY (CUSTOMER_ID) REFERENCES CUSTOMER(CUSTOMER_ID) ON DELETE CASCADE,
    FOREIGN KEY (FOOD_ITEM_ID) REFERENCES FOOD_ITEM(FOOD_ITEM_ID) ON DELETE CASCADE
);

CREATE TABLE ORDER_METHOD(
    ORDER_METHOD_ID CHAR(2),
    CHANNEL VARCHAR(20),
    PRIMARY KEY (ORDER_METHOD_ID)
);

CREATE TABLE PAYMENT_METHOD(
    PAYMENT_METHOD_ID CHAR(2),
    NAME VARCHAR(20),
    PRIMARY KEY (PAYMENT_METHOD_ID)
);

CREATE TABLE ORDERS(
    ORDER_ID CHAR(2),
    CUSTOMER_ID CHAR(2),
    ORDER_METHOD_ID CHAR(2),
    PAYMENT_METHOD_ID CHAR(2),
    ORDER_DATETIME DATETIME NOT NULL,
    RATING INT,
    PRIMARY KEY (ORDER_ID),
    FOREIGN KEY (CUSTOMER_ID) REFERENCES CUSTOMER(CUSTOMER_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (ORDER_METHOD_ID) REFERENCES ORDER_METHOD(ORDER_METHOD_ID),
    FOREIGN KEY (PAYMENT_METHOD_ID) REFERENCES PAYMENT_METHOD(PAYMENT_METHOD_ID)
);

CREATE TABLE ORDER_DETAIL(
    ORDER_DETAIL_ID CHAR(2),
    QUANTITY INT NOT NULL,
    RATING INT,
    FOOD_ITEM_ID CHAR(2),
    ORDER_ID CHAR(2),
    PROMOTION_ID CHAR(2),
    ORDER_TOTAL DECIMAL(4, 2) NOT NULL,
    TIP DECIMAL (4, 2) DEFAULT 0.00 NOT NULL,
    PRIMARY KEY (ORDER_DETAIL_ID),
    FOREIGN KEY (FOOD_ITEM_ID) REFERENCES FOOD_ITEM(FOOD_ITEM_ID),
    FOREIGN KEY (ORDER_ID) REFERENCES ORDERS(ORDER_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (PROMOTION_ID) REFERENCES PROMOTION(PROMOTION_ID)
);

CREATE TABLE ORDER_STATUS(
    ORDER_STATUS_ID CHAR(2),
    DESCRIPTION VARCHAR(30),
    PRIMARY KEY (ORDER_STATUS_ID)
);

CREATE TABLE ORDER_STATUS_BRIDGE(
    ORDER_DETAIL_ID CHAR(2),
    ORDER_STATUS_ID CHAR(2),
    FOOD_ITEM_ID CHAR(2),
    PRIMARY KEY (ORDER_DETAIL_ID, ORDER_STATUS_ID),
    FOREIGN KEY (ORDER_DETAIL_ID) REFERENCES ORDER_DETAIL(ORDER_DETAIL_ID) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (ORDER_STATUS_ID) REFERENCES ORDER_STATUS(ORDER_STATUS_ID) ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
