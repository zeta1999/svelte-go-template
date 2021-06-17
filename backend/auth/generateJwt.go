package auth

import (
	config "homefill/backend/config"
	"homefill/backend/errset"
	"homefill/backend/logs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string
	jwt.StandardClaims
}

// GenerateJwtToken : generates jwt token from given user id with expire time 3 hours
func GenerateJwtToken(userId string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = userId
	claims["time"] = time.Now().Add(time.Hour * 3)

	tokenString, err := token.SignedString(config.JWT_KEY)
	if err != nil {
		logs.LogIt(logs.LogWarn, "GenerateJwtToken", "unable to generate token", err)
		return "", errset.ErrInternalServer
	}

	return tokenString, nil
}
