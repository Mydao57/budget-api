package main

import (
	"log"

	"github.com/Mydao57/budget-api/initializers"
	"github.com/Mydao57/budget-api/routes"
	"github.com/Mydao57/logtofile"
)

func init() {
	initializers.LoadEnv()

	err := logtofile.NewLogger("./logs")
	if err != nil {
		log.Fatalf("Error creating log file: %v", err)
	}

	initializers.ConnectToDB()
}

func main() {
	defer logtofile.Close()

	logtofile.INFO("Application started successfully")

	routes.SetupRoutes()

}
