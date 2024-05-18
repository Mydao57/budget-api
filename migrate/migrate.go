package main

import (
	"github.com/Mydao57/budget-api/initializers"
	"github.com/Mydao57/budget-api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Budget{})
	initializers.DB.AutoMigrate(&models.Expense{})
}