--
-- version: 1528746007
--

BEGIN;

CREATE TABLE stream_items (
    item_id       BIGINT       PRIMARY KEY,
    stream_id     BIGINT       NOT NULL,
    entity_id     BIGINT       NOT NULL
);

CREATE UNIQUE INDEX ON stream_items (stream_id, entity_id);
CREATE INDEX ON stream_items (stream_id, item_id);

COMMIT;
