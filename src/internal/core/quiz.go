package core

import "time"

type Question struct {
	Question string   `json:"question" bson:"question"`
	Options  []string `json:"options" bson:"options"`
	Answer   string   `json:"answer" bson:"answer"`
	Duration int      `json:"duration"` // In seconds
}

type Quiz struct {
	QuizId     string      `json:"quiz_id" bson:"quiz_id"`         // For internal use
	UniqueCode string      `json:"unique_code" bson:"unique_code"` // For users
	Questions  []*Question `json:"questions" bson:"questions"`
	LiveTime   time.Time   `json:"live_time" bson:"live_time"`
	Duration   int         `json:"duration" bson:"duration"` // In seconds
}
