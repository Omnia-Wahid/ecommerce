CREATE TABLE IF NOT EXISTS products
(
    sku character(50) COLLATE pg_catalog."default" NOT NULL,
    name character(250) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT product_pkey PRIMARY KEY (sku)
    );
CREATE TABLE IF NOT EXISTS product_stocks
(
    sku character(50) COLLATE pg_catalog."default" NOT NULL,
    country character(50) COLLATE pg_catalog."default" NOT NULL,
    stock_change integer NOT NULL default 0,
    CONSTRAINT stock_unique UNIQUE (sku, country),
    CONSTRAINT sku_fk FOREIGN KEY (sku)
    REFERENCES products (sku) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE
    );