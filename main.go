package main

import (
	"log"

	"desq.com.ru/testjwtauth/config"
	"desq.com.ru/testjwtauth/routes"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	DB, err := initDB()
	if err != nil {
		log.Fatalf("Failed to estabilish database connection: %v", err)
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", DB)

		return c.Next()
	})

	app.Post("/me/:id/token", routes.CreateToken)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.JWTSecret)},
	}))
	app.Post("/me/token/refresh", routes.RefreshToken)

	app.Listen(":8080")
}
