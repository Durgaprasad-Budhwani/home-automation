package models

import "github.com/jinzhu/gorm"

type Hours int

type Scheduler struct {
	gorm.Model
	StartHours   Hours `json:"start_hours" valid:"numeric~Hours must be number,length(0|23)~Hours must be 0-23,required~Hours are required"`
	StartMinutes Hours `json:"start_minutes" valid:"numeric~Minutes must be number,length(0|59)~Hours must be 0-59,required~Minutes are required"`
	EndHours     Hours `json:"end_hours" valid:"numeric~Hours must be number,length(0|23)~Hours must be 0-23,required~Hours are required"`
	EndMinutes   Hours `json:"end_minutes" valid:"numeric~Minutes must be number,length(0|59)~Hours must be 0-59,required~Minutes are required"`
}
