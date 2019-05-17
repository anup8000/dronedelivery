package models

import "time"

//Order class denotes an individual order and its attributes
type Order struct {
	ID            int
	OrderID       string
	Coordinates   string
	Timestamp     time.Time
	Distance      float64
	Waitinminutes float64
	IsPromoter    bool
	IsDetractor   bool
	IsNeutral     bool
	ProcessedTime time.Time
}

//Orders is a collection of order
type Orders *[]Order
