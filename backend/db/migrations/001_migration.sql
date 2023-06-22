CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE groups(
    id uuid DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES users
);

CREATE TABLE user_group(
    user_id uuid NOT NULL,
    group_id uuid NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES users,
    FOREIGN KEY (group_id) REFERENCES groups
);

CREATE TABLE posts(
     id uuid DEFAULT uuid_generate_v4(),
     user_id uuid NOT NULL,
     group_id uuid NOT NULL,
     uri varchar NOT NULL,
     description varchar NOT NULL,
     plays integer NOT NULL DEFAULT 0,
     upvotes integer NOT NULL DEFAULT 0,
     downvotes integer NOT NULL DEFAULT 0,
     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (id),
     FOREIGN KEY (user_id) references users,
     FOREIGN KEY (group_id) references groups
);

CREATE TABLE comments(
    id uuid DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    text text NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users
);

CREATE TABLE post_comment(
    post_id uuid NOT NULL,
    comment_id uuid NOT NULL,
    PRIMARY KEY (post_id, comment_id),
    FOREIGN KEY (post_id) REFERENCES posts,
    FOREIGN KEY (comment_id) REFERENCES comments
);
