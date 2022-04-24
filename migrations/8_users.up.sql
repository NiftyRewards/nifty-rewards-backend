CREATE TABLE IF NOT EXISTS users
(
    address_w3a VARCHAR PRIMARY KEY,
    address_b   VARCHAR NOT NULL
);

INSERT INTO users (address_w3a, address_b)
VALUES ('0xUser1_w3a', '0xUser1_b'),
       ('0x32b2cc73939d4E58dCF8a751b2dc967d9fdEB3B4', '0x32b2cc73939d4E58dCF8a751b2dc967d9fdEB3B4')

