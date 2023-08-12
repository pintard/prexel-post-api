package db

import (
	"database/sql"
	"errors"
	. "prexel-post-api/model"
)

func CreatePost(post PrexelPost) (int64, error) {
	query := `INSERT INTO prexelposts (username, contact, content) VALUES ($1, $2, $3) RETURNING uuid;`
	var uuid int64
	err := DB.QueryRow(query, post.Username, post.Contact, post.Content).Scan(&uuid)
	if err != nil {
		return 0, err
	}
	return uuid, nil
}

func GetPost(uuid int64) (PrexelPost, error) {
	query := `SELECT uuid, username, contact, content FROM prexelposts WHERE uuid=$1;`
	var post PrexelPost
	err := DB.QueryRow(query, uuid).Scan(&post.UUID, &post.Username, &post.Contact, &post.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return PrexelPost{}, errors.New("post not found")
		}
		return PrexelPost{}, err
	}
	return post, nil
}

func GetAllPosts() ([]PrexelPost, error) {
	query := `SELECT uuid, username, contact, content FROM prexelposts;`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []PrexelPost
	for rows.Next() {
		var post PrexelPost
		if err := rows.Scan(&post.UUID, &post.Username, &post.Contact, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
