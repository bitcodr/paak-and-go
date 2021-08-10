BEGIN;

CREATE TABLE IF NOT EXISTS trips
(
    id              BIGSERIAL    NOT NULL PRIMARY KEY,
    origin_id       BIGINT       NOT NULL REFERENCES cities (id) ON DELETE RESTRICT,
    destination_id  BIGINT       NOT NULL REFERENCES cities (id) ON DELETE RESTRICT,
    dates           VARCHAR(255) NOT NULL,
    price           NUMERIC      NOT NULL,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT trip_id_must_not_negative CHECK (id > 0)
);

COMMIT;
