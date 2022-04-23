CREATE TABLE IF NOT EXISTS ownerships
(
    ownership_id      SERIAL PRIMARY KEY,
    address_w3a        VARCHAR REFERENCES users (address_w3a),
    collection_address VARCHAR REFERENCES nfts (collection_address),
    token_id           INT NOT NULL
);

INSERT INTO ownerships (address_w3a, collection_address, token_id)
VALUES ('0xUser1_w3a', '0x9999999', 1);

