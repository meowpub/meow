--
-- version: 1524739570
--

BEGIN;

ALTER TABLE entities ADD COLUMN id BIGINT PRIMARY KEY;
ALTER TABLE entities ADD CHECK (id != 0);

COMMIT;