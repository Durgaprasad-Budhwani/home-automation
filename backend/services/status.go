package services

import (
	"github.com/durgaprasad-budhwani/home-automation/backend/models"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type StatusService struct {
	db *gorm.DB
}

func NewStatusService(db *gorm.DB) StatusService {
	return StatusService{db}
}

func (s StatusService) Save(status *models.Status) error {
	log.Debug().Msg("[status - StatusService - Save] - Saving new status")
	return s.db.
		Create(&status).Error
}

func (s StatusService) Update(id string, status *models.Status) error {
	existingAgency := &models.Status{}
	err := s.db.
		Where("id = ?", id).
		First(existingAgency).Error
	if err != nil {
		return err
	}
	log.Debug().Msg("[status - StatusService - Update] - Updating status")
	return s.db.Model(&existingAgency).
		Update(&status).Error
}

func (s StatusService) GetAll(limit, page int) ([]models.Status, error) {
	log.Debug().Msg("[Status - StatusService - GetAll] - Fetching all statuses from db")
	var statuses []models.Status
	err := s.db.
		Order("name asc").
		Limit(limit).
		Offset(limit * page).
		Find(&statuses).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return statuses, nil
		}
		return nil, err
	}
	return statuses, nil
}
