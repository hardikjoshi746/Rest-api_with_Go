package models

import (
	"errors"
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	defer stmt.Close()

	hpass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hpass)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last insert id: %w", err)
	}

	u.ID = userId
	return nil
}

func (u *User) ValidateCridential() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)

	if err != nil {
		return errors.New("credentials incorrect")
	}

	passIsValid := utils.CheckPassword(u.Password, retrievePassword)

	if !passIsValid {
		return errors.New("credentials incorrect")
	}
	return nil
}
