--
-- version: 1524569269
--

BEGIN;

CREATE TABLE clients (
    id           BIGINT       PRIMARY KEY,
    name         VARCHAR(255) NOT NULL CHECK (name != ''),
    description  VARCHAR(255) NOT NULL CHECK (description != ''),
	redirect_uri VARCHAR(255) NOT NULL CHECK (redirect_uri != ''),
	secret       VARCHAR(255) NOT NULL CHECK (secret != '')
);

COMMIT;
