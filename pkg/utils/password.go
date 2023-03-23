package utils

import "golang.org/x/crypto/bcrypt"

// GeneratePassword - crypt password
func GeneratePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ValidatePassword - validate password
//  @hashedPassowrd: password stored at db, hashed form
//  @password: password from form, not hashed
func ValidatePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}
	return true
}
