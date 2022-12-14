package valueobject

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrTooShortPassword = errors.New("Password must be at least 8 characters long")
	ErrInvalidPassword = errors.New("Password must be in string and numbers")
)

type Password struct {
	Value string
}

func NewPassword(plainText string) (Password, error) {
	if(len(plainText) < 8) {
		return Password{}, ErrTooShortPassword 
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if(!re.MatchString(plainText)) {
		return Password{}, ErrInvalidPassword 
	}
	return Password{Value: plainText}, nil
}

func (p *Password)Encrypt(password string) (Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}

	return Password{Value: string(hash)}, nil
}

func (p *Password) IsMatchPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Value), []byte(password))	
	return err != bcrypt.ErrMismatchedHashAndPassword
}
