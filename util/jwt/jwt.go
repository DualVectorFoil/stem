package jwt

import (
	"errors"
	"github.com/DualVectorFoil/stem/app/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userId int64, userName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := jwt.StandardClaims{
		Audience:  userName,
		ExpiresAt: time.Now().Add(conf.USER_INFO_TTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(conf.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	if len(token) == 0 {
		return nil, errors.New("Uncorrected token")
	}
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.JWT_SECRET), nil
	})
	if err != nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
