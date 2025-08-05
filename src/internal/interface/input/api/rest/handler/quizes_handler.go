package handler

import (
	"context"
	"multiplayer-quiz-application/src/internal/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckAPIHandler(c *fiber.Ctx) error {
	return c.SendString("API is running")
}

type QuizesHandler struct {
	quizesService usecase.QuizesServiceImpl
}

func NewQuizesHandler(quizesService usecase.QuizesServiceImpl) *QuizesHandler {
	return &QuizesHandler{quizesService: quizesService}
}

func (qh *QuizesHandler) GetQuizByUniqueCodeHandler(c *fiber.Ctx) error {
	quizCode := c.Query("quiz_code")
	if quizCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no quiz code given",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	quiz, err := qh.quizesService.GetQuizByCode(ctx, quizCode)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"quiz": quiz,
	})
}
