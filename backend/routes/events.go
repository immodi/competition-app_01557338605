package routes

import (
	"database/sql"
	"encoding/json"
	"immodi/submission-backend/helpers"
	"immodi/submission-backend/repos"
	"immodi/submission-backend/routes/requests"
	"immodi/submission-backend/routes/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func EventsRouter(r chi.Router, db *sql.DB) {
	eventRepository := repos.NewEventRepository(db)

	r.Get("/", getAllEvents(eventRepository))
	r.Post("/", createEvent(eventRepository))
	r.Get("/{id}", getEvent(eventRepository))
}

func getAllEvents(eventRepository *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := eventRepository.GetAllEvents()
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "Failed to get all events")
			return
		}

		helpers.HttpJson(w, http.StatusOK, events)
	}
}

func createEvent(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requests.EventCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid request, likey an invalid schema")
			return
		}

		date, err := time.Parse(time.RFC3339, req.Date)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid date format, only RFC3339 is supported")
			return
		}

		if req.Name == "" || req.Description == "" || req.Category == "" || date.IsZero() || req.Venue == "" || req.Price == 0 {
			helpers.HttpError(w, http.StatusBadRequest, "missing name, description, category, date, venue or price")
			return
		}

		eventId, err := eventRepo.CreateEvent(req.Name, req.Description, req.Category, date.String(), req.Venue, req.Price, req.Image)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "could not create event")
			return
		}

		res := &responses.EventCreateResponse{
			EventId: eventId,
		}

		helpers.HttpJson(w, http.StatusCreated, res)
	}
}

func getEvent(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		eventId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "Invalid eventId, pass a valid one")
			return
		}

		event, err := eventRepo.GetEventById(eventId)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "couldn't get event")
			return
		}
		if event == nil {
			helpers.HttpError(w, http.StatusNotFound, "Event not found")
			return
		}

		helpers.HttpJson(w, http.StatusOK, event)
	}
}
