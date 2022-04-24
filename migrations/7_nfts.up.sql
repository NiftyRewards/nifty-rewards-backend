CREATE TABLE IF NOT EXISTS nfts
(
    collection_address VARCHAR PRIMARY KEY,
    collection_name    VARCHAR NOT NULL,
    total_supply       int     NOT NULL
);

INSERT INTO nfts (collection_address, collection_name, total_supply)
VALUES ('0xBAYC', 'bayc', 10),
       ('0x539d628a5fa811f3087648585bacb30870295f8b', 'Azuki', 5);
