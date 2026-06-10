-- PostgreSQL schema for the post & comment application.
-- The API currently runs with an in-memory repository; this script is the
-- equivalent relational design, ready to plug in via a PostgreSQL repository
-- implementation of the same domain interfaces.

CREATE TABLE posts (
    id          SERIAL PRIMARY KEY,
    author_name VARCHAR(100) NOT NULL,
    image_url   VARCHAR(500),
    posted_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE comments (
    id          SERIAL PRIMARY KEY,
    post_id     INTEGER      NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    author_name VARCHAR(100) NOT NULL,
    message     TEXT         NOT NULL CHECK (BTRIM(message) <> ''),
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_comments_post_id ON comments (post_id, created_at);

-- Seed data matching the in-memory repository.
INSERT INTO posts (author_name, image_url, posted_at)
VALUES ('Change can', 'assets/post.png', '2021-10-16 16:00:00+07');

INSERT INTO comments (post_id, author_name, message, created_at)
VALUES (1, 'Blend 285', 'have a good day', '2021-10-16 16:05:00+07');
