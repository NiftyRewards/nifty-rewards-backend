CREATE TABLE IF NOT EXISTS campaigns
(
    campaign_id        SERIAL PRIMARY KEY,
    merchant_id        INT REFERENCES merchants (merchant_id),
    collection_address VARCHAR REFERENCES nfts (collection_address),
    start_time         TIMESTAMP,
    end_time           TIMESTAMP
);

INSERT INTO campaigns (merchant_id, collection_address, start_time, end_time)
VALUES (1, '0xBAYC', now(), now());
