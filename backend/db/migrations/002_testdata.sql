INSERT INTO users(id, username, created_at)
VALUES (
        '2326d0e8-18ba-4034-825b-8b8bdfc15353', 'user1', current_timestamp
       );

INSERT INTO groups(id, owner_id, name, created_at)
VALUES (
        'd9db7a53-5a2a-49ad-a3b5-1eaadeef5ca4', '2326d0e8-18ba-4034-825b-8b8bdfc15353', 'group1', current_timestamp
       );

INSERT INTO user_group(user_id, group_id)
VALUES (
        '2326d0e8-18ba-4034-825b-8b8bdfc15353', 'd9db7a53-5a2a-49ad-a3b5-1eaadeef5ca4'
       );

INSERT INTO posts(id, user_id, group_id, uri, description, plays, upvotes, downvotes)
VALUES ('5d480a17-a0d9-41b1-b07f-2803fcbaa878', '2326d0e8-18ba-4034-825b-8b8bdfc15353', 'd9db7a53-5a2a-49ad-a3b5-1eaadeef5ca4','http://testuri.co',
        'test description', 110, 1, 2);

INSERT INTO comments(id, user_id, post_id, text)
VALUES (
        '41ad9376-6390-43d8-82f4-80216baf5bf4', '2326d0e8-18ba-4034-825b-8b8bdfc15353', '5d480a17-a0d9-41b1-b07f-2803fcbaa878', 'test comment text'
       );
