package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/durgaprasad-budhwani/home-automation/backend/routes"
	"github.com/durgaprasad-budhwani/home-automation/backend/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	fmt.Print("Application started")

	configureRuntime()

	viper := utils.NewViper("./config.yaml")
	fmt.Print(viper.GetString("Database"))
	db, err := gorm.Open("sqlite3", viper.GetString("Database"))
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(viper.GetBool("Verbose"))
	defer db.Close()

	r := routes.SetupRouter(db)

	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	err = r.Run(":" + port)
	if err != nil {
		panic("failed to run application")
	}
}

func configureRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(1)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}
