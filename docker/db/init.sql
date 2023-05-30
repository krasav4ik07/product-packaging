CREATE TABLE boxes
(
    pk_sscc VARCHAR(18) PRIMARY KEY,
    created_at TIMESTAMP
);

CREATE TABLE products
(
    pk_gtin VARCHAR(14) PRIMARY KEY,
    name VARCHAR(50),
    amount_packs SMALLINT
);

CREATE TABLE packs
(
    pk_serial_number VARCHAR(13) PRIMARY KEY,
    fk_box_sscc VARCHAR(18) REFERENCES boxes (pk_sscc),
    fk_product_gtin VARCHAR(14) REFERENCES products (pk_gtin)
);

INSERT INTO products (pk_gtin, name, amount_packs) VALUES
                                                   ('04603988000001', 'Продукт1', 4),
                                                   ('04603988000002', 'Продукт2', 48),
                                                   ('04603988000003', 'Продукт3', 180),
                                                   ('04603988000004', 'Продукт4', 360),
                                                   ('04603988000005', 'Продукт5', 18),
                                                   ('04603988000006', 'Продукт6', 36),
                                                   ('04603988000007', 'Продукт7', 15),
                                                   ('04603988000008', 'Продукт8', 90),
                                                   ('04603988000009', 'Продукт9', 144)
