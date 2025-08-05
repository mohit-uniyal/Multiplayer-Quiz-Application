package routes

import (
	"log"
	"multiplayer-quiz-application/src/internal/interface/input/api/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App,
	quizesHandler *handler.QuizesHandler,
) {
	api := app.Group("/api")

	api.Get("/", handler.CheckAPIHandler)

	api.Get("/quiz", quizesHandler.GetQuizByUniqueCodeHandler)

	log.Printf("routes initialized")
}
