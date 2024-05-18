package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Mydao57/budget-api/initializers"
	"github.com/Mydao57/budget-api/models"
	"github.com/Mydao57/logtofile"
	"github.com/gin-gonic/gin"
)

func BudgetCreate(c *gin.Context) {
	var body struct {
		DateStart string
		DateEnd   string
		Amount    float64
	}

	c.Bind(&body)

	layout := "2006-01-02"

	dateStart, err := time.Parse(layout, body.DateStart)
	if err != nil {
		logtofile.ERROR("Error while parsing " + body.DateStart)
		c.JSON(400, gin.H{"error": "Erreur lors de la conversion de DateStart", "details": err.Error()})
		return
	}

	dateEnd, err := time.Parse(layout, body.DateEnd)
	if err != nil {
		logtofile.ERROR("Error while parsing " + body.DateEnd)
		c.JSON(400, gin.H{"error": "Erreur lors de la conversion de DateEnd", "details": err.Error()})
		return
	}

	budget := models.Budget{
		DateStart: dateStart,
		DateEnd:   dateEnd,
		Amount:    body.Amount,
	}

	result := initializers.DB.Create(&budget)

	if result.Error != nil {
		logtofile.ERROR("Error while creating budget")
		c.JSON(400, gin.H{"error": "Erreur lors de la création du budget", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Budget created successfully")
	c.JSON(200, gin.H{"budget": budget})
}

func BudgetList(c *gin.Context) {

	var budgets []models.Budget
	result := initializers.DB.Find(&budgets)

	if result.Error != nil {
		logtofile.ERROR("Error while fetching budgets")
		c.JSON(400, gin.H{"error": "Erreur lors de la récupération des budgets", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Budgets fetched successfully")

	c.JSON(200, gin.H{"budgets": budgets})

}

func BudgetShow(c *gin.Context) {
	id := c.Param("id")

	var budget models.Budget
	result := initializers.DB.First(&budget, id)

	if result.Error != nil {

		logtofile.ERROR("Error while fetching budget")

		c.JSON(400, gin.H{"error": "Erreur lors de la récupération du budget", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Budget fetched successfully")

	c.JSON(200, gin.H{"budget": budget})
}

func BudgetUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		DateStart string
		DateEnd   string
		Amount    float64
	}

	c.Bind(&body)

	layout := "2006-01-02"

	dateStart, err := time.Parse(layout, body.DateStart)
	if err != nil {

		logtofile.ERROR("Error while parsing " + body.DateStart)

		c.JSON(400, gin.H{"error": "Erreur lors de la conversion de DateStart", "details": err.Error()})
		return
	}

	dateEnd, err := time.Parse(layout, body.DateEnd)
	if err != nil {
		logtofile.ERROR("Error while parsing " + body.DateEnd)

		c.JSON(400, gin.H{"error": "Erreur lors de la conversion de DateEnd", "details": err.Error()})
		return
	}

	var budget models.Budget
	result := initializers.DB.First(&budget, id)

	if result.Error != nil {
		logtofile.ERROR("Error while fetching budget")

		c.JSON(400, gin.H{"error": "Erreur lors de la récupération du budget", "details": result.Error.Error()})
		return
	}

	budget.DateStart = dateStart
	budget.DateEnd = dateEnd
	budget.Amount = body.Amount

	result = initializers.DB.Save(&budget)

	if result.Error != nil {
		logtofile.ERROR("Error while updating budget")
		c.JSON(400, gin.H{"error": "Erreur lors de la mise à jour du budget", "details": result.Error.Error()})
		return
	}
	logtofile.INFO("Budget updated successfully")

	c.JSON(200, gin.H{"budget": budget})

}

func BudgetDelete(c *gin.Context) {

	id := c.Param("id")

	var budget models.Budget
	result := initializers.DB.First(&budget, id)

	if result.Error != nil {
		logtofile.ERROR("Error while fetching budget")
		c.JSON(400, gin.H{"error": "Erreur lors de la récupération du budget", "details": result.Error.Error()})
		return
	}

	result = initializers.DB.Delete(&budget)

	if result.Error != nil {
		logtofile.ERROR("Error while deleting budget")

		c.JSON(400, gin.H{"error": "Erreur lors de la suppression du budget", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Budget deleted successfully")

	c.JSON(200, gin.H{"budget": budget})

}

func GetRemainingBudget(c *gin.Context) {
    budgetID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid budget ID"})
        return
    }

    var budget models.Budget
    if err := initializers.DB.First(&budget, budgetID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
        return
    }

    var expenses []models.Expense
    if err := initializers.DB.Where("date >= ? AND date <= ?", budget.DateStart, budget.DateEnd).Find(&expenses).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching expenses"})
        return
    }

    var totalExpenses float64
    for _, expense := range expenses {
        totalExpenses += expense.Amount
    }

    remainingBudget := budget.Amount - totalExpenses

    c.JSON(http.StatusOK, gin.H{
        "budget_id":       budget.ID,
        "remaining_budget": remainingBudget,
    })
}