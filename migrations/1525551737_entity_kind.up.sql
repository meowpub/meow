--
-- version: 1525551737
--

BEGIN;

ALTER TABLE entities ADD COLUMN kind VARCHAR(16) NOT NULL;
ALTER TABLE entities ADD CHECK (kind != '');

COMMIT;
