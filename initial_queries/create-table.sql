CREATE TABLE products (
    id int AUTO_INCREMENT,
    product_code varchar(10) NOT NULL,
    price FLOAT(5) NOT NULL,
    stock int NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (product_code)
);

CREATE TABLE orders (
    id int AUTO_INCREMENT,
    product_code varchar(10) NOT NULL,
    quantity int NOT NULL,
    price FLOAT(5) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (product_code) REFERENCES products(product_code)
);

CREATE TABLE campaigns (
    id int AUTO_INCREMENT,
    name varchar(10) NOT NULL,
    product_code varchar(10) NOT NULL,
    duration int NOT NULL,
    price_manipulation_limit int NOT NULL,
    target_sales_count int NOT NULL,
    started_at varchar(10) NOT NULL,
    finished_at varchar(10) NOT NULL,
    status varchar(10) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (name),
    FOREIGN KEY (product_code) REFERENCES products(product_code)
);

/* Tests. Check for Logs */
/*
INSERT INTO products (product_code, price, stock) 
VALUES ('Product1', 1.20, 5);

INSERT INTO orders (product_code, quantity, price) 
VALUES ('Product1', 5, 6.00);

INSERT INTO campaigns (name, product_code, duration, price_manipulation_limit, target_sales_count) 
VALUES ('Campaign1', 'Product1', 1, 40, 50);

SELECT * FROM products;
SELECT * FROM orders;
SELECT * FROM campaigns;
*/