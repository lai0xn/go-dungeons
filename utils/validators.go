package utils

import (
	"errors"
	"strings"
)

func ValidateUsername(username string) error {
	if len(username) > 20 || len(username) < 6 {
		return errors.New("Your username is not valid")
	}
	if strings.ContainsAny(username, "@#$&%!~") {
		return errors.New("Your username is not valid")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 10 {
		return errors.New("Your password is too short")
	}
	if len(password) > 30 {
		return errors.New("Your password is too long")
	}
	if strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") == false {
		return errors.New("Your password has no uppercase characters")
	}

	return nil
}

func ValidateAuth(username string, password string) []error {
	var errors []error = nil
	if u_err := ValidateUsername(username); u_err != nil {
		errors = append(errors, u_err)
	}
	if p_err := ValidatePassword(password); p_err != nil {
		errors = append(errors, p_err)
	}

	return errors
}
