package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	Db *gorm.DB
}

var Database *Db

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(2)
	}

	log.Println("connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// todo: add migrations
	Database = DbInstanse{Db: db}
}
