CREATE TABLE IF NOT EXISTS articles (
    article_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    contents TEXT NOT NULL,
    username VARCHAR(100) NOT NULL,
    nice INTEGER NOT NULL,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id SERIAL PRIMARY KEY,
    article_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles(article_id)
);