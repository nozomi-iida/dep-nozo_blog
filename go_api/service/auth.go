package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/domain/user/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

var (
	ErrDuplicateUsername = errors.New("Duplicate username")
	ErrUnMatchPassword = errors.New("Unmatch password")
)

type AuthConfiguration func(as *AuthService) error

type AuthService struct {
	users user.UserRepository
}

type AuthResponse struct {
	User entity.User `json:"user"`
	Token string `json:"token"`
}

func NewAuthService(cfgs ...AuthConfiguration) (*AuthService, error)  {
	os := &AuthService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithSqliteUserRepository(fileString string) AuthConfiguration  {
	return func(as *AuthService) error {
		u, err := sqlite.New(fileString)
		if err != nil {
			return err
		}
		as.users = u

		return nil
	}
}

func (as *AuthService) SignUp(username string, password string) (AuthResponse, error)  {
	ps, err := valueobject.NewPassword(password)
	u, err := entity.NewUser(username, ps)
	if err != nil {
		return AuthResponse{}, err
	}

	user, err := as.users.Create(u)
	if err != nil {
		return AuthResponse{}, err
	}

	token, err := generateToken(user.GetID());
	if err != nil {
		return AuthResponse{}, err
	}
	
	return AuthResponse{User: user, Token: token}, nil
}

func (as *AuthService) SignIn(username string, password string) (AuthResponse, error)  {
	user, err := as.users.FindByUsername(username)	
	if err != nil {
		return AuthResponse{}, err
	}
	if !user.IsMatchPassword(password) {
		return AuthResponse{}, ErrUnMatchPassword
	}
	token, err := generateToken(user.GetID());
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{User: user, Token: token}, nil
}

func generateToken(id uuid.UUID) (string, error)  {
	tokenString, _ := valueobject.NewJwtToken(id)
	token, err := tokenString.Encode();
	if err != nil {
		return "", err
	}
	return token, nil
}
