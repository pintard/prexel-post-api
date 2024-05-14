-- @block
CREATE TABLE IF NOT EXISTS prexel_posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL ,
    date TIMESTAMP NOT NULL,
    code TEXT NOT NULL,
    title TEXT NOT NULL,
    image_path TEXT,
    CONSTRAINT fk_prexel_user
        FOREIGN KEY(user_id)
        REFERENCES prexel_users(id)
        ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_prexel_posts_user_id ON prexel_posts(user_id);
CREATE INDEX IF NOT EXISTS idx_prexel_posts_date ON prexel_posts(date);