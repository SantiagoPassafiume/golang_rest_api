package models

import (
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

	log.Print(err)

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
