package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"immodi/submission-backend/helpers"
	"immodi/submission-backend/repos"
	"immodi/submission-backend/routes/requests"
	"immodi/submission-backend/routes/responses"
	helper_structs "immodi/submission-backend/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func EventsRouter(r chi.Router, db *sql.DB, api *helper_structs.API) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, getAllEvents(api.EventRepo, r))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			return api.UserRepo.IsAdmin(username)
		}, createEvent(api.EventRepo))
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, getEvent(api.EventRepo))
	})

	r.Get("/category/{category}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, getEventsByCategory(api.EventRepo))
	})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			return api.UserRepo.IsAdmin(username)
		}, updateEvent(api.EventRepo))
	})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			return api.UserRepo.IsAdmin(username)
		}, deleteEvent(api.EventRepo))
	})

	r.Get("/upcoming", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, getUpcomingEvents(api.EventRepo))
	})
	r.Get("/search/{keyword}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, searchEvents(api.EventRepo))
	})

	r.Post("/assign/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			userId, err := helpers.ParseTheUserIdFromRequest(r)
			if err != nil {
				return false
			}
			return api.UserRepo.IsSameUser(username, userId) || api.UserRepo.IsAdmin(username)
		}, assignEvent(api.EventRepo, api.UserRepo))
	})
}

func getAllEvents(eventRepository *repos.EventRepository, r *http.Request) http.HandlerFunc {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	return func(w http.ResponseWriter, r *http.Request) {
		events, err := eventRepository.GetAllEvents()
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "Failed to get all events")
			return
		}

		startIndex := (page - 1) * limit
		endIndex := min(page*limit, len(events))
		events = events[startIndex:endIndex]

		helpers.HttpJson(w, http.StatusOK, events)
	}
}

func createEvent(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requests.EventRequest
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

		eventId, err := eventRepo.CreateEvent(req.Name, req.Description, req.Category, date.String(), req.Venue, req.Price, req.Image, req.Translations)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "could not create event")
			return
		}

		res := &responses.EventResponse{
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
			helpers.HttpError(w, http.StatusBadRequest, "Invalid id, pass a valid one")
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

func getEventsByCategory(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := chi.URLParam(r, "category")
		events, err := eventRepo.GetEventsByCategory(category)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "this category has no events")
			return
		}

		if len(events) == 0 || events == nil {
			events = []repos.Event{}
		}

		helpers.HttpJson(w, http.StatusOK, events)
	}
}

func updateEvent(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		eventId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "Invalid id, pass a valid one")
			return
		}

		var req requests.EventRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid request, likey an invalid schema")
			return
		}

		date, err := time.Parse(time.RFC3339, req.Date)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid date format, only RFC3339 is supported")
			return
		}

		if eventId == 0 || req.Name == "" || req.Description == "" || req.Category == "" || date.IsZero() || req.Venue == "" || req.Price == 0 {
			helpers.HttpError(w, http.StatusBadRequest, "missing id, name, description, category, date, venue or price")
			return
		}

		err = eventRepo.UpdateEvent(eventId, req.Name, req.Description, req.Category, date.String(), req.Venue, req.Price, req.Image, req.Translations)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "could not update the event")
			return
		}

		res := &responses.EventResponse{
			EventId: eventId,
		}

		helpers.HttpJson(w, http.StatusOK, res)
	}
}

func deleteEvent(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		eventId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "Invalid id, pass a valid one")
			return
		}

		err = eventRepo.DeleteEvent(eventId)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "could not delete the event")
			return
		}

		res := &responses.EventDeletionResponse{
			Id:      eventId,
			Message: "the event with the above id was deleted successfully",
		}

		helpers.HttpJson(w, http.StatusOK, res)
	}
}

func getUpcomingEvents(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := eventRepo.GetUpcomingEvents()
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "wasn't able to get upcoming events, please try again later")
			return
		}
		helpers.HttpJson(w, http.StatusOK, events)
	}
}

func searchEvents(eventRepo *repos.EventRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyword := r.URL.Query().Get("keyword")
		events, err := eventRepo.SearchEvents(keyword)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "couldn't search events, please try again later")
			return
		}
		helpers.HttpJson(w, http.StatusOK, events)
	}

}

func assignEvent(eventRepo *repos.EventRepository, userRepo *repos.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		eventId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid id, pass a valid one")
			return
		}

		var req requests.EventAssignRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid request, likey an invalid schema")
			return
		}

		if eventId == 0 || req.UserID == 0 {
			helpers.HttpError(w, http.StatusBadRequest, "missing event id or user id")
			return
		}

		user, err := userRepo.GetUserById(req.UserID)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, fmt.Sprintf("could not find the user with id '%d'", req.UserID))
			return
		}
		if user == nil {
			helpers.HttpError(w, http.StatusNotFound, fmt.Sprintf("user with id '%d' not found", req.UserID))
			return
		}

		err = eventRepo.RegisterUserToEvent(user.ID, eventId)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, fmt.Sprintf("could not assign the event to the user with id '%d'", req.UserID))
			return
		}

		res := &responses.EventResponse{
			EventId: eventId,
		}

		helpers.HttpJson(w, http.StatusOK, res)
	}
}
