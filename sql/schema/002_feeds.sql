-- +goose Up
Create table feeds (
    id UUID primary key,
    user_id UUID not null,
    name VARCHAR(50) not null,
    url  VARCHAR(50) unique not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    FOREIGN key (user_id) References Users (id) ON DELETE CASCADE
);

-- +goose Down
Drop table feeds;