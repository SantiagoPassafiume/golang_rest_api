package models

import (
	"errors"
	"github.com/SantiagoPassafiume/golang_rest_api/db"
	"github.com/SantiagoPassafiume/golang_rest_api/utils"
	"log"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)
`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		log.Print(err)
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return nil
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT id, password FROM users WHERE email = ?
`
	row := db.DB.QueryRow(query, &u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
