package main

import (
	"fmt"
	"time"

	"github.com/durgaprasad-budhwani/home-automation/backend/models"
	"github.com/durgaprasad-budhwani/home-automation/backend/utils"

	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func addSchedule(db *gorm.DB) {
	fmt.Println("Checking running tasks")

	now := time.Now()
	currentDate := utils.RoundToDay(now)

	var slots []models.Slot
	err := db.
		Find(&slots).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		panic(err)
	}

	var existingSchedules []models.Schedule
	err = db.
		Where("date = ?", currentDate).
		Preload("Slot").
		Find(&existingSchedules).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			for _, slot := range slots {
				schedule := models.Schedule{
					Slot: slot,
					Date: currentDate,
				}
				fmt.Print("saving schedule")
				if saveErr := db.Save(&schedule).Error; saveErr != nil {
					panic(saveErr)
				}
			}
			return
		}
	}

	for _, slot := range slots {
		exists := false
		for _, existingSchedule := range existingSchedules {
			if slot.ID == existingSchedule.Slot.ID {
				exists = true
				break
			}
		}
		if !exists {
			schedule := models.Schedule{
				Slot: slot,
				Date: currentDate,
			}
			fmt.Print("saving schedule")
			if saveErr := db.Save(&schedule).Error; saveErr != nil {
				panic(saveErr)
			}
		}
	}
}

func startStopMotor(db *gorm.DB) {
	var schedules []models.Schedule
	now := time.Now()
	currentDate := utils.RoundToDay(now)
	timeInUTC := currentDate.Add(time.Hour*time.Duration(now.Hour()) +
		time.Minute*time.Duration(now.Minute()) +
		time.Second*time.Duration(now.Second()))
	err := db.
		Where("date = ?", currentDate).
		Preload("Slot").
		First(&schedules).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		panic(err)
	}

	var motorRunning bool
	for _, schedule := range schedules {
		startTime := currentDate.Add(time.Hour*time.Duration(int(schedule.Slot.StartHours)) + time.Minute*time.Duration(int(schedule.Slot.StartMinutes)))
		endTime := currentDate.Add(time.Hour*time.Duration(int(schedule.Slot.EndHours)) + time.Minute*time.Duration(int(schedule.Slot.EndMinutes)))
		if startTime.Before(timeInUTC) && endTime.After(timeInUTC) {
			motorRunning = true
			break
		}
	}

	if motorRunning {
		// start motor
		fmt.Print("motor is running")
		return
	}

	fmt.Print("motor not running")
}

func main() {
	fmt.Print("Cron started")

	viper := utils.NewViper("./config.yaml")
	fmt.Print(viper.GetString("Database"))
	db, err := gorm.Open("sqlite3", viper.GetString("Database"))
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(viper.GetBool("Verbose"))
	defer db.Close()
	// Do jobs with params
	gocron.Every(5).Seconds().Do(addSchedule, db)
	gocron.Every(1).Second().Do(startStopMotor, db)

	<-gocron.Start()
}
