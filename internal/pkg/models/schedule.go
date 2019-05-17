package models

import "time"

//Schedule class denotes the schedule for the order
type Schedule struct {
	Orders            *[]Order
	ScheduleStartTime time.Time
	ScheduleEndTime   time.Time
	Schedules         *[]Order
	NetPromoterScore  float64
	IsGreedy          bool
	InputFilePath     string
	OutputFilePath    string
}
