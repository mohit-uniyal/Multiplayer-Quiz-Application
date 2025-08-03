package routes

import (
	"log"
	"multiplayer-quiz-application/src/internal/interface/input/api/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", handler.CheckAPIHandler)

	log.Printf("routes initialized")
}
