package Events

import (
	"fmt"
	"time"
)

type EventManager struct {
	events map[int]*Event
	lastID int
}

func NewEventManager() *EventManager {
	return &EventManager{make(map[int]*Event), 0}
}

/*func (em *EventManager) findEvent(e *Event) int { //find event by date, user_id and title
	for id, event := range em.events {
		if strings.Compare(e.UserID, event.UserID) == 0 {
			if e.Date.Equal(event.Date) && strings.Compare(e.UserID, event.UserID) == 0 {
				return id
			}
		}
	}
	return -1
}*/

func (em *EventManager) AddEvent(e *Event) {
	e.ID = em.lastID
	em.lastID++
	em.events[e.ID] = e
}

func (em *EventManager) UpdateEvent(e *Event) bool {
	if cEvent, err := em.events[e.ID]; err {
		cEvent.mutex.Lock()
		if e.Title != "" {
			cEvent.Title = e.Title
		}
		if !e.Date.IsZero() {
			cEvent.Date = e.Date
		}
		if e.Description != "" {
			cEvent.Description = e.Description
		}
		if e.UserID != "" {
			cEvent.UserID = e.UserID
		}
		cEvent.mutex.Unlock()
	} else {
		fmt.Println("Event with id", e.ID, "not found")
		return err
	}
	return true
}

func (em *EventManager) DeleteEvent(id int) bool {
	var err bool
	em.events[id].mutex.Lock()
	if _, err = em.events[id]; err {
		delete(em.events, id)
	} else {
		fmt.Println("Event with id", id, "not found")
	}
	return err
}

func (em *EventManager) GetEventsInPeriod(userID string, begin, end time.Time) (result []*Event) {
	for _, event := range em.events {
		if event.UserID == userID {
			if event.Date.After(begin) && event.Date.Before(end) || event.Date.Equal(begin) {
				result = append(result, event)
			}
		}
	}
	return result
}
