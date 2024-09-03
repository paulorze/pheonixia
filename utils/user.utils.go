package utils

import (
	customErrors "phoenixia/errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUsername(username string) (string, error) {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]{4,}$`)
	if !regex.MatchString(username) {
		return "", &customErrors.InvalidUsername
	}
	trimmedUsername := strings.Trim(username, "")
	return trimmedUsername, nil
}

func ValidatePassword(password string) (string, error) {
	secure := true
	tests := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, password)
		if !t {
			secure = false
			break
		}
	}
	if !secure {
		return "", &customErrors.InvalidPassword
	}
	trimmedPassword := strings.Trim(password, "")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(trimmedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", &customErrors.ServerError
	}
	return string(hashedPassword), nil
}

func ValidateEmail(email string) (string, error) {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(email) {
		return "", &customErrors.InvalidEmail
	}
	trimmedEmail := strings.Trim(email, "")
	return trimmedEmail, nil
}

func ValidateName(name string) (string, error) {
	regex := regexp.MustCompile(`^[a-zA-ZáéíóúüñÁÉÍÓÚÜÑ]{2,}$`)
	if !regex.MatchString(name) {
		return "", &customErrors.InvalidName
	}
	return strings.Trim(name, ""), nil
}
