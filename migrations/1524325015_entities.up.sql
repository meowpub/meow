--
-- version: 1524325015
--

BEGIN;

-- The entities table is really simple - just JSON blobs and indices.
-- Turns out Postgres is better at being MongoDB than MongoDB is. Not webscale tho.
CREATE TABLE entities ( data JSONB NOT NULL );

-- Ensure that no entities lack an '@id', not just because that's incorrect, but because
-- that caused trouble with the UNIQUE constraint (NULL != NULL for some reason).
ALTER TABLE entities ADD CONSTRAINT entities_data_id_not_null_check CHECK (data ? '@id');

-- The indexing here is done on "->>", which coerces to string in case it's a different type.
CREATE UNIQUE INDEX entities_data_id_idx ON entities ((data ->> '@id'));

COMMIT;
