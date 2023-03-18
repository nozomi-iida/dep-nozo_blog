package valueobject

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrTooShortPassword = errors.New("Password must be at least 8 characters long")
	ErrInvalidPassword  = errors.New("Password must be in string and numbers")
)

type Password struct {
	Value string
}

func NewPassword(plainText string) (Password, error) {
	if len(plainText) < 8 {
		return Password{}, ErrTooShortPassword
	}

	regString := regexp.MustCompile(`[a-zA-Z]`).Match([]byte(plainText))
	if !regString {
		return Password{}, ErrInvalidPassword
	}
	regInt := regexp.MustCompile(`[0-9]`).Match([]byte(plainText))
	if !regInt {
		return Password{}, ErrInvalidPassword
	}

	return Password{Value: plainText}, nil
}

func (p *Password) Encrypt() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
