package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	bytes := []byte(pass)

	hashed, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)

	return string(hashed), err
}

func CheckPassword(hashed, current string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(current))
	return err == nil
}
