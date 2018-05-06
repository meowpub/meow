--
-- version: 1525550010
--

BEGIN;

CREATE TABLE users (
    id            BIGINT       PRIMARY KEY,
    entity_id     BIGINT       NOT NULL UNIQUE REFERENCES entities (id) ON DELETE CASCADE,
    email         VARCHAR(255) NOT NULL UNIQUE CHECK (email != ''),
    password_hash VARCHAR(255) NOT NULL CHECK (password_hash != '')
);

COMMIT;
