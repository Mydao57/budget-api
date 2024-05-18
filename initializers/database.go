package initializers

import (
	"os"

	"github.com/Mydao57/logtofile"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logtofile.FATAL("Error connecting to database. Closing application.")
		os.Exit(1)
	} else {
		logtofile.INFO("Connected to database")
	}
}
