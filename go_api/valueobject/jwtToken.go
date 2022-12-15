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
	UserId uuid.UUID
}

type CustomClaims struct {
	UserId uuid.UUID `json:"userId"`
	jwt.RegisteredClaims
}

// TODO: keyを正しく設定する
var mySigningKey = []byte("AllYourBase")


func NewJwtToken(userId uuid.UUID) (JwtToken, error)  {
	return JwtToken{UserId: userId}, nil	
}

func (jt *JwtToken)Decode(tokenString string) (CustomClaims, error)  {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return *claims, nil	
	} else {
		return CustomClaims{}, err
	}
}

func (jt *JwtToken) Encode() (string, error)  {
	claims := CustomClaims{
		jt.UserId,
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
