INSERT INTO articles (title, contents, username, nice, created_at)
VALUES ('firstPost', 'This is my first blog', 'moriT', 2, CURRENT_TIMESTAMP);


INSERT INTO articles (title, contents, username, nice)
VALUES ('2nd', 'Second blog post', 'moriT', 4);


INSERT INTO comments (article_id, message, created_at)
VALUES (1, '1st comment yeah', CURRENT_TIMESTAMP);


INSERT INTO comments (article_id, message)
VALUES (1, 'welcome');