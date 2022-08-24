package Server

import (
	Events "task11/Calendar"
	"time"
)

type eventJSON struct {
	ID          int    `json:"id"`
	Date        string `json:"date"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (e *eventJSON) ParseToModel() *Events.Event {
	event := Events.NewEvent()
	event.ID = e.ID
	event.UserID = e.UserID
	event.Title = e.Title
	event.Description = e.Description
	event.Date, _ = time.Parse("2006-01-02", e.Date)
	return event
}

func (e *eventJSON) isOK() bool {
	if e.UserID != "" && e.Title != "" {
		if _, ok := time.Parse("2006-01-02", e.Date); ok == nil {
			return true
		}
	}
	return false
}
