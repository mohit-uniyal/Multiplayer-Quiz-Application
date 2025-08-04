package main

import (
	"log"
	mongo_persistance "multiplayer-quiz-application/src/internal/adapters/persistance/mongo"
	"multiplayer-quiz-application/src/internal/config"
	"multiplayer-quiz-application/src/internal/interface/input/api/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//1. Load Configs
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	//2. Initialize DB connections

	mongoDb, err := mongo_persistance.NewDatabase(config)
	if err != nil {
		log.Fatalf("failed to connect to mongo DB: %v", err)
	}
	defer mongoDb.Close()

	//3. Initialize repo code

	_ = mongo_persistance.NewQuestionsRepo(mongoDb)

	//4. Initialize services
	//5. Initialize handlers
	//6. Initialize routes
	app := fiber.New()
	routes.InitializeRoutes(app)
	//7. Initialize server

	app.Listen(":" + config.PORT)
}
