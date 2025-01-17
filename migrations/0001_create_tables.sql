CREATE TABLE currencies
(
    id   SERIAL PRIMARY KEY,
    coin VARCHAR(10) UNIQUE
);

CREATE TABLE prices
(
    coin_id   INTEGER,
    price     NUMERIC,
    timestamp BIGINT,
    FOREIGN KEY (coin_id) REFERENCES currencies (id)
);