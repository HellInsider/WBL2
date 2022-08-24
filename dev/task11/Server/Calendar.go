package Server

import (
	"encoding/json"
	"log"
	"net/http"
	Events "task11/Calendar"
	"time"
)

type Calendar struct {
	events  *Events.EventManager
	server  *http.Server
	routing *http.ServeMux
}

func NewServer() *Calendar {
	return &Calendar{Events.NewEventManager(), nil, http.NewServeMux()}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}

func (calendar *Calendar) Run() {
	calendar.server = &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: calendar.routing,
	}
	calendar.routing.HandleFunc("/create_event", Logger(calendar.CreateEvent))
	calendar.routing.HandleFunc("/update_event", Logger(calendar.UpdateEvent))
	calendar.routing.HandleFunc("/delete_event", Logger(calendar.DeleteEvent))
	calendar.routing.HandleFunc("/events_for_day", Logger(calendar.EventsForDay))
	calendar.routing.HandleFunc("/events_for_week", Logger(calendar.EventsForWeek))
	calendar.routing.HandleFunc("/events_for_month", Logger(calendar.EventsForMonth))
	log.Println("Server started")
	defer log.Println("Shutting down server")
	log.Fatal(calendar.server.ListenAndServe())

}

func (calendar *Calendar) CreateEvent(w http.ResponseWriter, r *http.Request) {
	jsonE := eventJSON{}
	if isValid(w, r, http.MethodPost) {
		if err := json.NewDecoder(r.Body).Decode(&jsonE); err != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		} else if !jsonE.isOK() {
			jsonResponse(true, w, http.StatusServiceUnavailable, "Incorrect user_id or date")
		} else {
			calendar.events.AddEvent(jsonE.ParseToModel())
			jsonResponse(false, w, http.StatusOK, "Create complete")
		}
	}
}

func (calendar *Calendar) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if isValid(w, r, http.MethodPost) {
		jsonE := eventJSON{}
		if ok := json.NewDecoder(r.Body).Decode(&jsonE); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else if ok := calendar.events.UpdateEvent(jsonE.ParseToModel()); !ok {
			jsonResponse(true, w, http.StatusServiceUnavailable, "Update error")
		} else {
			jsonResponse(false, w, http.StatusOK, "Update complete")
		}
	}
}

func (calendar *Calendar) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if isValid(w, r, http.MethodPost) {
		jsonE := eventJSON{}
		if err := json.NewDecoder(r.Body).Decode(&jsonE); err != nil || jsonE.ID == -1 {
			jsonResponse(true, w, http.StatusServiceUnavailable, "Incorrect ID")
			return
		}
		if ok := calendar.events.DeleteEvent(jsonE.ID); !ok {
			jsonResponse(true, w, http.StatusServiceUnavailable, "Event not found")
			return
		}
		jsonResponse(false, w, http.StatusOK, "Delete complete")
	}
}
func (calendar *Calendar) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if isValid(w, r, http.MethodGet) {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			userID := r.URL.Query().Get("user_id")
			evs := calendar.events.GetEventsInPeriod(userID, date, date.AddDate(0, 0, 1))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}

}

func (calendar *Calendar) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if isValid(w, r, http.MethodGet) {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			userID := r.URL.Query().Get("user_id")
			evs := calendar.events.GetEventsInPeriod(userID, date, date.AddDate(0, 0, 7))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}

}

func (calendar *Calendar) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if isValid(w, r, http.MethodGet) {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusBadRequest, ok.Error())
		} else {
			userID := r.URL.Query().Get("user_id")
			evs := calendar.events.GetEventsInPeriod(userID, date, date.AddDate(0, 1, 0))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}
