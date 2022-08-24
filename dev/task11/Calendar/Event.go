package Events

import (
	"sync"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	mutex       sync.Mutex
}

func NewEvent() *Event {
	return new(Event)
}
