package auth

import (
	"fmt"
	config "homefill/backend/config"
	"homefill/backend/errset"
	"homefill/backend/logs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func VerifyJwt(tokenString string) (string, error) {

	tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			logs.LogIt(logs.LogWarn, "VerifyJwt", "error in jwt verfication", fmt.Errorf("false"))
			return nil, errset.ErrBadRequest
		}
		return config.JWT_KEY, nil
	})

	if err != nil {
		logs.LogIt(logs.LogWarn, "VerifyJwt", "error in jwt parsing", err)
		return "", errset.ErrUnAuthorized
	}

	claims := tkn.Claims.(jwt.MapClaims)
	s := claims["time"].(string)
	t, err := time.Parse(time.RFC3339, s)

	if err != nil {
		logs.LogIt(logs.LogWarn, "VerifyJwt", "error in time parsing", err)
		return "", errset.ErrBadRequest
	}

	if time.Now().Before(t) {
		return claims["userid"].(string), nil
	} else {
		return "", errset.ErrUnAuthorized
	}
}
