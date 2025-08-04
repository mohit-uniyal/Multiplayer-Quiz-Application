package mongo_persistance

type QuestionsRepoImpl interface {
}

type QuestionsRepo struct {
	db *Database
}

func NewQuestionsRepo(db *Database) *QuestionsRepo {
	return &QuestionsRepo{db: db}
}
