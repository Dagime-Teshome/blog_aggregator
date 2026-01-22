
-- +goose Up
CREATE Table posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP not NULL,
    title   TEXT NOT NULL,
    url TEXT not NULL UNIQUE,
    description   TEXT,
    published_at TIMESTAMP ,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
Drop table posts;