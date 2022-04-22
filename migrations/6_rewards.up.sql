CREATE TABLE IF NOT EXISTS rewards
(
    rewards_id   SERIAL PRIMARY KEY,
    campaign_id  INT REFERENCES campaigns (campaign_id),
    ownership_id INT REFERENCES ownerships (ownership_id),
    description  VARCHAR NOT NULL,
    quantity     INT     NOT NULL
);

INSERT INTO rewards (campaign_id, ownership_id, description, quantity)
VALUES (1, 1, 'rewards_desc', 4);

