package models

import "time"

// Custom event type(struct)
type Event struct {
	ID int
	Title string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

func (e Event) Save() {
	// task : add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}