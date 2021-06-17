package service

import (
	"homefill/backend/auth"
	"homefill/backend/db"
	"homefill/backend/model"

	"github.com/gofiber/fiber/v2"
)

func GenerateJwtTokenService(user *model.User) (string, error) {

	err := db.DB.Repo.InsertUser(user)
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateJwtToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtMiddleWare(c *fiber.Ctx) (string, error) {
	jwtToken := c.Get("Authorization")[7:]
	id, err := auth.VerifyJwt(jwtToken)

	if err != nil {
		return "", err
	}
	return id, nil
}
