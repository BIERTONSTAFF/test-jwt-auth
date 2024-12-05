package routes

import (
	"time"

	"desq.com.ru/testjwtauth/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func setRTCookie(c *fiber.Ctx, RT string, expires time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    RT,
		HTTPOnly: true,
		Expires:  expires,
	})
}

func CreateToken(c *fiber.Ctx) error {
	DB := c.Locals("db").(*gorm.DB)

	uID := c.Params("id")
	userID, err := uuid.Parse(uID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	expires := time.Now().Add(time.Hour * 24)

	T, RT, err := handlers.CreateToken(DB, userID, c.IP(), expires.Unix())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	setRTCookie(c, RT, expires)

	return c.Status(201).JSON(fiber.Map{
		"token": T,
	})
}

func RefreshToken(c *fiber.Ctx) error {
	DB := c.Locals("db").(*gorm.DB)
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	RT := c.Cookies("refreshToken")
	if RT == "" {
		return c.Status(400).SendString("Refresh cookie must be set")
	}

	uID := claims["userId"].(string)

	userID, err := uuid.Parse(uID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	err = handlers.RefreshToken(DB, userID, claims, RT, c.IP())
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}

	expires := time.Now().Add(time.Hour * 24)

	t, RT, err := handlers.CreateToken(DB, userID, c.IP(), expires.Unix())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	setRTCookie(c, RT, expires)

	return c.Status(200).JSON(fiber.Map{
		"token": t,
	})
}
