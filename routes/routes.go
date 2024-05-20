package routes

import (
	"github.com/Mydao57/budget-api/controllers"
	"github.com/Mydao57/logtofile"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	logtofile.INFO("Setting up routes")

	router.POST("/budget", controllers.BudgetCreate)
	router.GET("/budget", controllers.BudgetList)
	router.GET("/budget/:id", controllers.BudgetShow)
	router.PUT("/budget/:id", controllers.BudgetUpdate)
	router.DELETE("/budget/:id", controllers.BudgetDelete)
	router.GET("/budget/:id/remaining", controllers.GetRemainingBudget)

	router.POST("/expense", controllers.ExpenseCreate)
	router.GET("/expense", controllers.ExpenseList)
	router.GET("/expense/:id", controllers.ExpenseShow)
	router.PUT("/expense/:id", controllers.ExpenseUpdate)
	router.DELETE("/expense/:id", controllers.ExpenseDelete)

	routes := router.Routes()
	for _, route := range routes {
		logtofile.INFO("Route loaded: " + route.Path + ` (` + route.Method + `)`)
	}

	if err := router.Run(); err != nil {
		logtofile.ERROR("Error loading routes: " + err.Error())
	}

}
