package hash

import "golang.org/x/crypto/bcrypt"

func GetHashedPassword(rawPass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPass), 8)
	return string(hashedPassword), err
}
