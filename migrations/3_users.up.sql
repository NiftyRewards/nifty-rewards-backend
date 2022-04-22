CREATE TABLE IF NOT EXISTS users
(
    address_w3a VARCHAR PRIMARY KEY,
    address_b   VARCHAR NOT NULL
);

INSERT INTO users (address_w3a, address_b)
VALUES ('0x123', '0x456');

