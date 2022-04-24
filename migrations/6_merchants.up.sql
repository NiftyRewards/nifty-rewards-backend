CREATE TABLE IF NOT EXISTS merchants
(
    merchant_id   SERIAL PRIMARY KEY,
    merchant_name VARCHAR NOT NULL
);

INSERT INTO merchants (merchant_name)
VALUES ('merchant1');