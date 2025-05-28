CREATE TABLE category (
    category_id  SERIAL PRIMARY KEY NOT NULL,
    category_name VARCHAR(50) NOT NULL
);

CREATE TABLE product (
    product_id SERIAL PRIMARY KEY NOT NULL,
    category_id INTEGER NOT NULL,
    product_name VARCHAR(50) NOT NULL,
    characteristics VARCHAR(100) NOT NULL,
    FOREIGN KEY (category_id) 
        REFERENCES category(category_id)
        ON UPDATE CASCADE
        ON DELETE NO ACTION
);

CREATE TABLE store_product (
    upc VARCHAR(12) PRIMARY KEY NOT NULL,
    upc_prom VARCHAR(12),
    product_id INTEGER NOT NULL,
    selling_price DECIMAL(13,4) NOT NULL,
    products_number INTEGER NOT NULL,
    promotional_product BOOLEAN NOT NULL,
    FOREIGN KEY (upc_prom)
        REFERENCES store_product(upc)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY (product_id) 
        REFERENCES product(product_id)
        ON UPDATE CASCADE
        ON DELETE NO ACTION
);

CREATE TABLE employee (
    employee_id VARCHAR(10) PRIMARY KEY NOT NULL,
    empl_surname VARCHAR(50) NOT NULL,
    empl_name VARCHAR(50) NOT NULL,
    empl_patronymic VARCHAR(50),
    empl_role VARCHAR(10) NOT NULL,
    salary DECIMAL(13,4) NOT NULL,
    date_of_birth DATE NOT NULL,
    date_of_start DATE NOT NULL,
    phone_number VARCHAR(13) NOT NULL,
    city VARCHAR(50) NOT NULL,
    street VARCHAR(50) NOT NULL,
    zip_code VARCHAR(9) NOT NULL
);

CREATE TABLE customer_card (
    card_number VARCHAR(13) PRIMARY KEY NOT NULL,
    cust_surname VARCHAR(50) NOT NULL,
    cust_name VARCHAR(50) NOT NULL,
    cust_patronymic VARCHAR(50),
    phone_number VARCHAR(13) NOT NULL,
    city VARCHAR(50),
    street VARCHAR(50),
    zip_code VARCHAR(9),
    percent INTEGER NOT NULL
);

CREATE TABLE receipt (
    receipt_number VARCHAR(10) PRIMARY KEY NOT NULL,
    employee_id VARCHAR(10) NOT NULL,
    card_number VARCHAR(13),
    print_date TIMESTAMP NOT NULL,
    sum_total DECIMAL(13,4) NOT NULL,
    vat DECIMAL(13, 4) NOT NULL,
    FOREIGN KEY (employee_id)
        REFERENCES employee(employee_id)
        ON UPDATE CASCADE
        ON DELETE NO ACTION,
    FOREIGN KEY (card_number) 
        REFERENCES customer_card(card_number)
        ON UPDATE CASCADE
        ON DELETE NO ACTION
);


CREATE TABLE sale (
    upc VARCHAR(12) NOT NULL,
    receipt_number VARCHAR(10) NOT NULL,
    product_number INTEGER Not NULL,
    selling_price DECIMAL(13,4) NOT NULL,
    PRIMARY KEY (upc, receipt_number),
    FOREIGN KEY (upc)
        REFERENCES store_product(upc)
        ON UPDATE CASCADE
        ON DELETE NO ACTION,
    FOREIGN KEY (receipt_number)
        REFERENCES receipt(receipt_number)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
