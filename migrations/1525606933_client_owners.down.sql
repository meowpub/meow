--
-- version: 1525606933
--

BEGIN;

ALTER TABLE clients DROP COLUMN owner_id;

COMMIT;
