package services

import (
	"github.com/durgaprasad-budhwani/home-automation/backend/models"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type SlotService struct {
	db *gorm.DB
}

func NewSlotService(db *gorm.DB) SlotService {
	return SlotService{db}
}

func (s SlotService) Save(scheduler *models.Slot) error {
	log.Debug().Msg("[Slot - SlotService - Save] - Saving new scheduler")
	return s.db.
		Create(&scheduler).Error
}

func (s SlotService) Update(id uint, scheduler *models.Slot) error {
	existingAgency := &models.Slot{}
	err := s.db.
		Where("id = ?", id).
		First(existingAgency).Error
	if err != nil {
		return err
	}
	log.Debug().Msg("[Slot - SlotService - Update] - Updating scheduler")
	return s.db.Model(&existingAgency).
		Update(&scheduler).Error
}

func (s SlotService) GetAll(limit, page int) ([]models.Slot, error) {
	log.Debug().Msg("[Slot - SlotService - GetAll] - Fetching all schedulers from db")
	var schedulers []models.Slot
	err := s.db.
		Find(&schedulers).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return schedulers, nil
		}
		return nil, err
	}
	return schedulers, nil
}
