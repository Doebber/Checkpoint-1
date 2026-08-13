package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		loc, err := time.LoadLocation("America/Sao_Paulo")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Não foi possível carregar o fuso horário.",
			})
		}

		now := time.Now().In(loc)

		return c.JSON(fiber.Map{
			"timezone": "America/Sao_Paulo",
			"data":     now.Format("02/01/2006"),
			"hora":     now.Format("15:04:05"),
		})
	})

	log.Fatal(app.Listen(":8080"))
}
