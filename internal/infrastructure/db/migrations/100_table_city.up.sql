BEGIN;

CREATE TABLE IF NOT EXISTS "cities"
(
    id             BIGSERIAL    NOT NULL PRIMARY KEY,
    name           VARCHAR(100) NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT city_id_must_not_negative CHECK (id > 0)
);

COMMIT;  

