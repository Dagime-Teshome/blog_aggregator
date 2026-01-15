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

Create table feed_follows (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id UUID not null,
    feed_id UUID not null,
    FOREIGN key (user_id) References users (id) ON DELETE CASCADE,
    FOREIGN key (feed_id) References feeds (id) ON DELETE CASCADE,
    unique(user_id,feed_id)
)

-- +goose Down
Drop table feeds;