package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
	Id string
}

type JWT struct {
	Secret string
}

func NewJwt(secret string) *JWT {
	return &JWT{Secret: secret}
}

func(j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
		"id": data.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	email := t.Claims.(jwt.MapClaims)["email"]
	id := t.Claims.(jwt.MapClaims)["id"]
	return t.Valid, &JWTData{Email: email.(string), Id: id.(string)}
}

