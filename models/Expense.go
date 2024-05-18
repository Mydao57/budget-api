package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Date 	  time.Time
	Amount    float64
}