-- +goose Up
CREATE TABLE collections (
    id TEXT PRIMARY KEY,

	title TEXT NOT NULL CHECK(title<>''),

    created INTEGER NOT NULL,
    updated INTEGER NOT NULL
);

CREATE TABLE images (
    collection_id TEXT NOT NULL REFERENCES collections(id) ON DELETE CASCADE, 
    hash TEXT NOT NULL CHECK(hash<>''),

	filename TEXT NOT NULL CHECK(filename<>''),

    created INTEGER NOT NULL,
    updated INTEGER NOT NULL,

    PRIMARY KEY(collection_id, hash)
);

-- +goose Down
DROP TABLE images;
DROP TABLE collections;
