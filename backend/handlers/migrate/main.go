package main

import (
	"fmt"
	models2 "github.com/durgaprasad-budhwani/home-automation/backend/models"

	"github.com/durgaprasad-budhwani/home-automation/backend/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	fmt.Print("Migration started")
	viper := utils.NewViper("./config.yaml")
	db, err := gorm.Open("sqlite3", viper.GetString("Database"))
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(viper.GetBool("Verbose"))
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models2.Scheduler{}, &models2.Status{})
}