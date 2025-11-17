package models

import (
	"errors"

	"example.com/eventbookingrestapi/db"
	"example.com/eventbookingrestapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// func (u User) Save() error {
// 	query := "INSERT INTO users(email, password) VALUES($1, $2)"

// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	hashedPass, err := utils.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := stmt.Exec(u.Email, hashedPass)
// 	if err != nil {
// 		return err
// 	}

// 	userId, err := result.LastInsertId()

// 	u.ID = userId
// 	return err

// }

func (u *User) Save() error {
	// Use RETURNING id to get the inserted user's ID
	query := "INSERT INTO users(email, password) VALUES($1, $2) RETURNING id"

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Use QueryRow to execute the insert and get the ID
	err = db.DB.QueryRow(query, u.Email, hashedPass).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.UnHashPassword(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
