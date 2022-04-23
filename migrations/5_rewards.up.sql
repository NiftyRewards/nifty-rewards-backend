CREATE TABLE IF NOT EXISTS rewards
(
    reward_id          SERIAL PRIMARY KEY,
    merchant_id        INT REFERENCES merchants (merchant_id),
    collection_address VARCHAR REFERENCES nfts (collection_address),
    token_id           INT     NOT NULL,
    description        VARCHAR NOT NULL,
    max_quantity       INT     NOT NULL,
    quantity_used      INT DEFAULT 0
);

INSERT INTO rewards (merchant_id, collection_address, token_id, description, max_quantity, quantity_used)
VALUES (1, '0xBAYC', 555, 'rewards1_desc', 4, 0);

