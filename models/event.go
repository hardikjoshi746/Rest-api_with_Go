package models

import "time"

// Event struct with corrected struct tags
type Event struct {
	ID          int       `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserId      int       `json:"user_id"`
}

// Global slice to store events
var events = []Event{}

// Save method to add an event to the events slice
func (e *Event) Save() {
	events = append(events, *e)
}

// GetAllEvent returns all saved events
func GetAllEvent() []Event {
	return events
}
