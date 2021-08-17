package db

import "golang.org/x/crypto/bcrypt"

/* Encrypt Password is the Routine that allow the password encription */
func EncryptPasword(pass string) (string, error) {
	//Es un algoritmo que es 2 elevado al costo. La cantidad de pasadas que va a encriptar al texto. Entre mayor sea el costo, mas segura sera la contrasena, pero mas demora
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
