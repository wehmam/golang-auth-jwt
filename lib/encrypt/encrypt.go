package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func EncrpytFromString(value string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	return string(encryptedPassword), err
}

func CompareHashPassword(value string, compare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(value), []byte(compare))
	return err == nil
}
