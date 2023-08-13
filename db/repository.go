package db

import (
	"database/sql"
	"errors"
	. "prexel-post-api/model"
)

func CreatePost(post PrexelPost) (int64, error) {
	var uuid int64
	query := `INSERT INTO prexelposts (username, contact, code, date) VALUES ($1, $2, $3, $4) RETURNING uuid;`
	err := DB.QueryRow(query, post.Username, post.Contact, post.Code, post.Date).Scan(&uuid)

	if err != nil {
		return 0, err
	}

	return uuid, nil
}

func GetPost(uuid int64) (PrexelPost, error) {
	var post PrexelPost
	query := `SELECT uuid, username, contact, code, date FROM prexelposts WHERE uuid=$1;`
	err := DB.QueryRow(query, uuid).Scan(&post.UUID, &post.Username, &post.Contact, &post.Code, &post.Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return PrexelPost{}, errors.New("post not found")
		}
		return PrexelPost{}, err
	}

	return post, nil
}

func PollPosts(lastUUID int64, limit int) ([]PrexelPost, error) {
	query := `SELECT uuid, username, contact, content, date FROM prexelposts WHERE uuid > $1 ORDER BY uuid ASC LIMIT $2;`
	rows, err := DB.Query(query, lastUUID, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var posts []PrexelPost

	for rows.Next() {
		var post PrexelPost
		if err := rows.Scan(&post.UUID, &post.Username, &post.Contact, &post.Code); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
