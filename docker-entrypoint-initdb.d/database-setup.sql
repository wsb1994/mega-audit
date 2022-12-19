drop table if exists blockchain;

CREATE TABLE blockchain (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data TEXT NOT NULL,
    previous_block_hash VARCHAR(64),
    block_hash VARCHAR(64)
);

INSERT INTO blockchain (data, previous_block_hash, block_hash)
VALUES ('genesis', 'genesis', 'genesis');