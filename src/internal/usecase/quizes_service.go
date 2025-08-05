package usecase

import (
	"context"
	mongo_persistance "multiplayer-quiz-application/src/internal/adapters/persistance/mongo"
	"multiplayer-quiz-application/src/internal/core"
)

type QuizesServiceImpl interface {
	GetQuizByCode(ctx context.Context, code string) (*core.Quiz, error)
}

type QuizesService struct {
	quizRepo mongo_persistance.QuizesRepoImpl
}

func NewQuizesService(quizRepo mongo_persistance.QuizesRepoImpl) *QuizesService {
	return &QuizesService{
		quizRepo: quizRepo,
	}
}

func (qs *QuizesService) GetQuizByCode(ctx context.Context, code string) (*core.Quiz, error) {
	return qs.quizRepo.FetchQuizByCode(ctx, code)
}
