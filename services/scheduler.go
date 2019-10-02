package services

import (
	"github.com/durgaprasad-budhwani/home-automation/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type SchedulerService struct {
	db *gorm.DB
}

func NewSchedulerService(db *gorm.DB) SchedulerService {
	return SchedulerService{db}
}

func (s SchedulerService) Save(scheduler *models.Scheduler) error {
	log.Debug().Msg("[Scheduler - SchedulerService - Save] - Saving new scheduler")
	return s.db.
		Create(&scheduler).Error
}

func (s SchedulerService) Update(id uint, scheduler *models.Scheduler) error {
	existingAgency := &models.Scheduler{}
	err := s.db.
		Where("id = ?", id).
		First(existingAgency).Error
	if err != nil {
		return err
	}
	log.Debug().Msg("[Scheduler - SchedulerService - Update] - Updating scheduler")
	return s.db.Model(&existingAgency).
		Update(&scheduler).Error
}

func (s SchedulerService) GetAll(limit, page int) ([]models.Scheduler, error) {
	log.Debug().Msg("[Scheduler - SchedulerService - GetAll] - Fetching all schedulers from db")
	var schedulers []models.Scheduler
	err := s.db.
		Order("name asc").
		Limit(limit).
		Offset(limit * page).
		Find(&schedulers).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return schedulers, nil
		}
		return nil, err
	}
	return schedulers, nil
}
