CREATE TABLE IF NOT EXISTS anime (
    id              TEXT PRIMARY KEY, --ksuid
    kitsu_id        TEXT UNIQUE,
    last_crawled    TIMESTAMPTZ DEFAULT now(),
    title           TEXT NOT NULL, -- Attributes.CanonicalTitle
    cover_image     TEXT,
    poster_image    TEXT,
    popularity_rank INT,
    average_rating  DECIMAL,
    synopsis        TEXT,
    rating_freq     BYTES,
    rating_rank     INT,
    favorites_count INT,

);