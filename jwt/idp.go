package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mainflux/manager"
)

const (
	issuer   string        = "mainflux"
	duration time.Duration = 10 * time.Hour
)

var _ manager.IdentityProvider = (*jwtIdentityProvider)(nil)

type jwtIdentityProvider struct {
	secret string
}

// NewIdentityProvider instantiates a JWT identity provider.
func NewIdentityProvider(secret string) manager.IdentityProvider {
	return &jwtIdentityProvider{}
}

func (idp *jwtIdentityProvider) Key(id string) (string, error) {
	now := time.Now().UTC()
	exp := now.Add(duration)

	claims := jwt.StandardClaims{
		Issuer:    issuer,
		IssuedAt:  now.Unix(),
		ExpiresAt: exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(idp.secret))
}

func (idp *jwtIdentityProvider) IsValid(key string) bool {
	token, _ := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(idp.secret), nil
	})

	return token.Valid
}
