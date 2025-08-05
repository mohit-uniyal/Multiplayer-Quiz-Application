package mongo_persistance

import (
	"context"
	"fmt"
	"multiplayer-quiz-application/src/internal/core"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	quizCollectionName = "quizes"
)

type QuizesRepoImpl interface {
	FetchQuizByCode(ctx context.Context, code string) (*core.Quiz, error)
}

type QuizesRepo struct {
	db *Database
}

func NewQuizesRepo(db *Database) *QuizesRepo {
	return &QuizesRepo{db: db}
}

func (qr *QuizesRepo) FetchQuizByCode(ctx context.Context, code string) (*core.Quiz, error) {

	filter := bson.M{
		"unique_code": code,
	}

	var quiz core.Quiz

	err := qr.db.DB.Collection(quizCollectionName).FindOne(ctx, filter).Decode(&quiz)
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz with code: %w", err)
	}

	return &quiz, nil
}
