package db

import (
	"database/sql"
	"errors"
	. "prexel-post-api/model"
)

func CreatePost(post PrexelPost) (int64, error) {
	var id int64
	query := `INSERT INTO prexelposts (user_id, code, date) VALUES ($1, $2, $3) RETURNING id;`
	err := DB.QueryRow(query, post.UserId, post.Code, post.Date).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetPost(id int64) (PrexelPost, error) {
	var post PrexelPost
	query := `SELECT id, user_id, code, date FROM prexelposts WHERE id=$1;`
	err := DB.QueryRow(query, id).Scan(&post.ID, &post.UserId, &post.Code, &post.Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return PrexelPost{}, errors.New("post not found")
		}
		return PrexelPost{}, err
	}

	return post, nil
}

func PollPosts(lastID *int64, limit int) ([]PrexelPost, error) {
	var posts []PrexelPost
	var query string
	var rows *sql.Rows
	var err error

	if lastID != nil {
		query = `SELECT id, user_id, code, date FROM prexelposts WHERE id > $1 ORDER BY id ASC LIMIT $2;`
		rows, err = DB.Query(query, *lastID, limit)
	} else {
		query = `SELECT id, user_id, code, date FROM prexelposts ORDER BY id ASC LIMIT $1;`
		rows, err = DB.Query(query, limit)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post PrexelPost
		if err := rows.Scan(&post.ID, &post.UserId, &post.Code, &post.Date); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
