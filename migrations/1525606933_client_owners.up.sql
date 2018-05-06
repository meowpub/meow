--
-- version: 1525606933
--

BEGIN;

ALTER TABLE clients ADD COLUMN owner_id BIGINT NOT NULL REFERENCES users (id);
ALTER TABLE clients ADD CHECK (owner_id != 0);

COMMIT;
