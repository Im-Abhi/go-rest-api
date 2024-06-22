package models

import (
	"errors"

	"github.com/Im-Abhi/leaning-go/rest-api/db"
	"github.com/Im-Abhi/leaning-go/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query :=
		`
		INSERT INTO users(email, password) VALUES(?, ?)
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	user.ID = id
	return err
}

func (user *User) ValidateCredentials() error {
	query :=
		`
		SELECT id, password FROM users WHERE email = ?
	`

	row := db.DB.QueryRow(query, user.Email)

	var retrivedPassword string
	err := row.Scan(&user.ID, &retrivedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
