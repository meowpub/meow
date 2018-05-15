--
-- version: 1526417398
--

BEGIN;

ALTER TABLE entities DROP CONSTRAINT entities_data_id_not_empty_check;

COMMIT;
