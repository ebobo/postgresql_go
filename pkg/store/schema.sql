-- lego
-- -------------------------

CREATE TABLE IF NOT EXISTS lego (
    id          SERIAL PRIMARY KEY NOT NULL,
    model_id    INTEGER NOT NULL UNIQUE,
    name        TEXT NOT NULL,
    catalog     TEXT NOT NULL
);