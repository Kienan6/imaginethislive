CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE groups(
    id uuid DEFAULT uuid_generate_v4(),
    owner_id uuid,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES users
);

CREATE TABLE user_group(
    user_id uuid,
    group_id uuid,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES users,
    FOREIGN KEY (group_id) REFERENCES groups
);

CREATE TABLE posts(
     id uuid DEFAULT uuid_generate_v4(),
     user_id uuid,
     uri varchar NOT NULL,
     description varchar NOT NULL,
     plays integer NOT NULL DEFAULT 0,
     upvotes integer NOT NULL DEFAULT 0,
     downvotes integer NOT NULL DEFAULT 0,
     PRIMARY KEY (id),
     FOREIGN KEY (user_id) references users
);

CREATE TABLE group_post(
    group_id uuid NOT NULL,
    post_id uuid NOT NULL,
    PRIMARY KEY (group_id, post_id),
    FOREIGN KEY (group_id) REFERENCES groups,
    FOREIGN KEY (post_id) REFERENCES posts
);

CREATE TABLE comments(
    id uuid DEFAULT uuid_generate_v4(),
    user_id uuid,
    text text NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users
);

CREATE TABLE post_comment(
    post_id uuid,
    comment_id uuid,
    PRIMARY KEY (post_id, comment_id),
    FOREIGN KEY (post_id) REFERENCES posts,
    FOREIGN KEY (comment_id) REFERENCES comments
)
