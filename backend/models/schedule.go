package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Schedule struct {
	gorm.Model
	SlotID uint
	Slot   Slot
	Date   time.Time `json:"date" valid:"rfc3339,required~Shift Date is required,trim~Shift Date is required"`
	Status string    `json:"status"`
}
