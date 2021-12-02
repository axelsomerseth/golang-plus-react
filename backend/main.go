package backend

import (
	"fmt"
	"log"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/config"
	"github.com/axelsomerseth/varsity-dev-challenge/backend/db"
	"github.com/axelsomerseth/varsity-dev-challenge/backend/router"
)

func Start() {
	// Load environment variables
	config.Load()

	// Connect Database
	db.Connect()

	// Define routes
	r := router.Setup()
	port := fmt.Sprintf(":%s", config.ENV.Port)

	// Handle errors
	if err := r.Run(port); err != nil {
		log.Fatal("failed to start the app. %w", err)
	}
}
