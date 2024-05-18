package models

import (
	"time"

	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	DateStart time.Time
	DateEnd   time.Time
	Amount    float64
}