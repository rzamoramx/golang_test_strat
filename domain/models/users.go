package models

import (
	"errors"
	"unicode"
)

// validate rules for user login
func ValidateUser(user string, pwd string, userEntity User) error {
	// if user is empty, return user not found
	if userEntity.User == "" {
		return errors.New("User not found")
	}

	// Validate user and password
	if user == "" || pwd == "" {
		return errors.New("user and pwd are required")
	}

	if userEntity.User != user {
		return errors.New("wrong user")
	}

	if userEntity.Password != pwd {
		return errors.New("wrong password")
	}

	return nil
}

// validate rules for user register
func RulesForRegisterUser(req map[string]any, user User) error {
	// user already exists
	if user.Email != "" {
		return errors.New("user already exists")
	}

	if user.Phone != "" {
		return errors.New("phone already exists")
	}

	// phone policy
	if len(req["Phone"].(string)) != 10 {
		return errors.New("phone must be 10 digits")
	}

	pwd := req["Password"].(string)

	// password policy
	if len(pwd) < 6 || len(pwd) > 12 {
		return errors.New("password must be between 6 and 12 characters")
	}

	if !verifyPassword(pwd) {
		return errors.New("password must have at least one lowercase, one uppercase, one special char")
	}

	return nil
}

// helper, check password complexity policy
// TODO: make regex implementation, but keep in mind that regex is quite expensive for CPU
func verifyPassword(s string) bool {
	var hasUpperCase, hasLowercase, hasSpecial bool
	for _, c := range s {
		switch {
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasUpperCase && hasLowercase && hasSpecial
}
