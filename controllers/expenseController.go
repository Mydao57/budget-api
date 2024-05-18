package controllers

import (
	"time"

	"github.com/Mydao57/budget-api/initializers"
	"github.com/Mydao57/budget-api/models"
	"github.com/Mydao57/logtofile"
	"github.com/gin-gonic/gin"
)

func ExpenseCreate(c *gin.Context) {
	var body struct {
		Date   string
		Amount float64
	}

	c.Bind(&body)

	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		logtofile.ERROR("Error while parsing " + body.Date)
		c.JSON(400, gin.H{"error": "Erreur lors de la conversion de Date", "details": err.Error()})
		return
	}

	expense := models.Expense{
		Date:   date,
		Amount: body.Amount,
	}

	result := initializers.DB.Create(&expense)

	if result.Error != nil {
		logtofile.ERROR("Error while creating expense")
		c.JSON(400, gin.H{"error": "Erreur lors de la création de la dépense", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Expenses created successfully")
	c.JSON(200, gin.H{"expense": expense})
}

func ExpenseList(c *gin.Context) {

	var expenses []models.Expense
	result := initializers.DB.Find(&expenses)

	if result.Error != nil {
		logtofile.ERROR("Error while fetching expenses")
		c.JSON(400, gin.H{"error": "Erreur lors de la récupération des dépenses", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Expsenses fetched successfully")

	c.JSON(200, gin.H{"expenses": expenses})

}

func ExpenseShow(c *gin.Context) {
	id := c.Param("id")

	var expense models.Expense
	result := initializers.DB.First(&expense, id)

	if result.Error != nil {

		logtofile.ERROR("Error while fetching expense")

		c.JSON(400, gin.H{"error": "Erreur lors de la récupération de la dépense", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Expense fetched successfully")

	c.JSON(200, gin.H{"expense": expense})
}

func ExpenseUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Date   string
		Amount float64
	}

	c.Bind(&body)

	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		logtofile.ERROR("Error while parsing " + body.Date)

		c.JSON(400, gin.H{"error": "Erreur lors de la conversion de Date", "details": err.Error()})
		return
	}

	var expense models.Expense
	result := initializers.DB.First(&expense, id)

	if result.Error != nil {
		logtofile.ERROR("Error while fetching expense")

		c.JSON(400, gin.H{"error": "Erreur lors de la récupération de la dépense", "details": result.Error.Error()})
		return
	}

	expense.Date = date
	expense.Amount = body.Amount

	result = initializers.DB.Save(&expense)

	if result.Error != nil {
		logtofile.ERROR("Error while updating expense")
		c.JSON(400, gin.H{"error": "Erreur lors de la mise à jour de la dépense", "details": result.Error.Error()})
		return
	}
	logtofile.INFO("Expense updated successfully")

	c.JSON(200, gin.H{"expense": expense})

}

func ExpenseDelete(c *gin.Context) {

	id := c.Param("id")

	var expense models.Expense
	result := initializers.DB.First(&expense, id)

	if result.Error != nil {
		logtofile.ERROR("Error while fetching expense")
		c.JSON(400, gin.H{"error": "Erreur lors de la récupération de la dépense", "details": result.Error.Error()})
		return
	}

	result = initializers.DB.Delete(&expense)

	if result.Error != nil {
		logtofile.ERROR("Error while deleting expense")

		c.JSON(400, gin.H{"error": "Erreur lors de la suppression de la dépense", "details": result.Error.Error()})
		return
	}

	logtofile.INFO("Expense deleted successfully")

	c.JSON(200, gin.H{"expense": expense})

}
