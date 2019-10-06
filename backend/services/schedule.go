package services

import (
	"github.com/durgaprasad-budhwani/home-automation/backend/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type SchedulerService struct {
	db *gorm.DB
}

func NewSchedulerService(db *gorm.DB) SchedulerService {
	return SchedulerService{db}
}

func (s SchedulerService) Save(status *models.Schedule) error {
	log.Debug().Msg("[status - SchedulerService - Save] - Saving new status")
	return s.db.
		Create(&status).Error
}

func (s SchedulerService) Update(id uint, status *models.Schedule) error {
	existingAgency := &models.Schedule{}
	err := s.db.
		Where("id = ?", id).
		First(existingAgency).Error
	if err != nil {
		return err
	}
	log.Debug().Msg("[status - SchedulerService - Update] - Updating status")
	return s.db.Model(&existingAgency).
		Update(&status).Error
}

func (s SchedulerService) GetAll(limit, page int) ([]models.Schedule, error) {
	log.Debug().Msg("[Schedule - SchedulerService - GetAll] - Fetching all statuses from db")
	var statuses []models.Schedule
	err := s.db.
		Find(&statuses).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return statuses, nil
		}
		return nil, err
	}
	return statuses, nil
}
