--
-- version: 1526417398
--

BEGIN;

-- Require that the ID is non-empty.
ALTER TABLE entities ADD CONSTRAINT entities_data_id_not_empty_check CHECK (
    trim(both from (data ->> '@id')) != ''
);

-- Require that the ID has a http:// or https:// protocol.
ALTER TABLE entities ADD CONSTRAINT entities_data_id_protocol_check CHECK (
    ((data ->> '@id') LIKE 'http://%') OR ((data ->> '@id') LIKE 'https://%')
);

COMMIT;
