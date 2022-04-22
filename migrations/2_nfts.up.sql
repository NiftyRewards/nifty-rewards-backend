CREATE TABLE IF NOT EXISTS nfts
(
    collection_address VARCHAR PRIMARY KEY,
    collection_name    VARCHAR NOT NULL
);

INSERT INTO nfts (collection_address, collection_name)
VALUES ('0x9999999', 'bayc');

