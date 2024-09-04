package models

import (
	"time"

	"example.com/rest-api/db"
)

// Event struct with corrected struct tags
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserId      int       `json:"user_id"`
}

// Global slice to store events
var events = []Event{}

// Save method to add an event to the events slice
func (e *Event) Save() error {
	query :=
		`INSERT INTO events(name, description, location, dateTime, user_id) 
	 VALUES(?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// GetAllEvent returns all saved events
func GetAllEvent() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
