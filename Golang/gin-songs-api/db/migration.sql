SET TIMEZONE = 'UTC';

CREATE TABLE artists
(
    id         VARCHAR PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    bio        TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

SELECT * FROM artists;


CREATE TABLE albums
(
    id           VARCHAR PRIMARY KEY,
    name         VARCHAR   NOT NULL,
    genre        VARCHAR   NOT NULL,
    artist_id    VARCHAR REFERENCES artists (id),
    release_date TIMESTAMP,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE songs
(
    id           VARCHAR PRIMARY KEY,
    name         VARCHAR   NOT NULL,
    album_id     VARCHAR REFERENCES albums (id),
    duration     FLOAT     NOT NULL,
    release_date TIMESTAMP,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
