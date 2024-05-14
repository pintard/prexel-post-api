package repository

import (
	"database/sql"
	"errors"
	. "prexel-post-api/src/model"
	"prexel-post-api/src/utils"
)

func CreateUser(user PrexelUser) (int64, error) {
	var id int64
	query := `INSERT INTO prexelusers (email, service, username, contact, contact_url, date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	err := utils.DB.QueryRow(query, user.Email, user.Service, user.Username, user.Contact, user.ContactURL, user.Date).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUser(id int64) (PrexelUser, error) {
	var user PrexelUser
	query := `SELECT id, email, service, username, contact, contact_url, date FROM prexelusers WHERE id=$1;`
	err := utils.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Service, &user.Contact, &user.ContactURL, &user.Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return PrexelUser{}, errors.New("user not found")
		}
		return PrexelUser{}, err
	}

	return user, nil
}
