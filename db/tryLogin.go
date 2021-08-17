package db

import (
	"Dolphin/models"

	"golang.org/x/crypto/bcrypt"
)

/* TryLogin check the login state with the database */
func TryLogin(email string, password string) (models.User, bool) {

	usr, found, _ := CheckExistingUser(email)

	if !found {
		return usr, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usr.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usr, false
	}
	return usr, true
}
