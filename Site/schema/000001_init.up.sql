CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username TEXT not null unique,
    password_hash varchar(255) not null
);
CREATE TABLE images (
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    image_id BIGSERIAL,
    image_url TEXT UNIQUE,
    like_count INT,
    title TEXT,
    PRIMARY KEY (user_id, image_url)
);
CREATE TABLE promts (
    title TEXT,
    image_url TEXT REFERENCES images(image_url) ON DELETE CASCADE,
    PRIMARY KEY (title, image_url)
);
CREATE TABLE likes(
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    image_url TEXT REFERENCES images(image_url) ON DELETE CASCADE,
    -- like_check for understand what pictures like user
    like_check boolean,
    PRIMARY KEY (user_id, image_url)
);
CREATE TABLE consideration(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) on DELETE CASCADE,
    image_url TEXT,
    title TEXT
);