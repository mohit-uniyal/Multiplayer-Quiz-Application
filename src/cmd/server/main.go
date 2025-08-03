package main

import (
	"log"
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
	//3. Initialize repo code
	//4. Initialize services
	//5. Initialize handlers
	//6. Initialize routes
	app := fiber.New()
	routes.InitializeRoutes(app)
	//7. Initialize server

	app.Listen(":" + config.PORT)
}
