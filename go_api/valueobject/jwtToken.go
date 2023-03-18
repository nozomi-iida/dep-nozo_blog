package valueobject

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ErrInvalidJwtToken = errors.New("A JwtToken has to have an a valid jwtToken")
)

type JwtToken struct {
	ID uuid.UUID
}

type CustomClaims struct {
	UserId uuid.UUID `json:"userId"`
	jwt.RegisteredClaims
}

// TODO: keyを正しく設定する
var mySigningKey = []byte("AllYourBase")

func NewJwtToken(id uuid.UUID) (JwtToken, error) {
	return JwtToken{ID: id}, nil
}

func (jt *JwtToken) Encode() (string, error) {
	claims := CustomClaims{
		jt.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Decode(tokenString string) (CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return *claims, nil
	} else {
		return CustomClaims{}, err
	}
}
