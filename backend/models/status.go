package models

import "time"

type Status struct {
	Scheduler
	Date time.Time `json:"shift_date" valid:"rfc3339,required~Shift Date is required,trim~Shift Date is required"`
}
