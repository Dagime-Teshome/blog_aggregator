
-- +goose Up
Create table feed_follows (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id UUID not null,
    feed_id UUID not null,
    FOREIGN key (user_id) References users (id) ON DELETE CASCADE,
    FOREIGN key (feed_id) References feeds (id) ON DELETE CASCADE,
    unique(user_id,feed_id)
);


-- +goose Down
Drop table feed_follows;