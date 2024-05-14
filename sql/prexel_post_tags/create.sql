CREATE TABLE prexel_post_tags (
    post_id INT REFERENCES prexel_posts(id) ON DELETE CASCADE,
    tag_id INT REFERENCES prexel_tags(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);
