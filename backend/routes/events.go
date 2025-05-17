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
		helpers.ProtectedHandler(w, r, nil, GetAllEvents(api.EventRepo, r))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			return api.UserRepo.IsAdmin(username)
		}, CreateEvent(api.EventRepo))
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, GetEvent(api.EventRepo))
	})

	r.Get("/category/{category}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, GetEventsByCategory(api.EventRepo, r))
	})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			return api.UserRepo.IsAdmin(username)
		}, UpdateEvent(api.EventRepo))
	})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			return api.UserRepo.IsAdmin(username)
		}, DeleteEvent(api.EventRepo))
	})

	// r.Get("/upcoming", func(w http.ResponseWriter, r *http.Request) {
	// 	helpers.ProtectedHandler(w, r, nil, getUpcomingEvents(api.EventRepo))
	// })

	r.Get("/search/{keyword}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, nil, SearchEvents(api.EventRepo, r))
	})

	r.Post("/assign/{id}", func(w http.ResponseWriter, r *http.Request) {
		helpers.ProtectedHandler(w, r, func(username string) bool {
			userId, err := helpers.ParseTheUserIdFromRequest(r)
			if err != nil {
				return false
			}
			return api.UserRepo.IsSameUser(username, userId) || api.UserRepo.IsAdmin(username)
		}, AssignEvent(api.EventRepo, api.UserRepo))
	})
}

func GetAllEvents(eventRepository repos.EventInterface, r *http.Request) http.HandlerFunc {
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

		if len(events) < limit*(page-1) {
			helpers.HttpError(w, http.StatusBadRequest, "requested page does not exist")
			return
		}

		eventsCount := len(events)
		startIndex := (page - 1) * limit
		endIndex := min(page*limit, len(events))
		events = events[startIndex:endIndex]
		resp := &responses.EventsResponse{
			Events: events,
			Count:  eventsCount,
		}

		helpers.HttpJson(w, http.StatusOK, resp)
	}
}

func CreateEvent(eventRepo repos.EventInterface) http.HandlerFunc {
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

func GetEvent(eventRepo repos.EventInterface) http.HandlerFunc {
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

func GetEventsByCategory(eventRepo repos.EventInterface, r *http.Request) http.HandlerFunc {
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
		category := chi.URLParam(r, "category")
		println(category)
		events, err := eventRepo.GetEventsByCategory(category)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "this category has no events")
			return
		}

		if len(events) < limit*(page-1) {
			helpers.HttpError(w, http.StatusBadRequest, "requested page does not exist")
			return
		}

		eventsCount := len(events)
		startIndex := (page - 1) * limit
		endIndex := min(page*limit, len(events))
		events = events[startIndex:endIndex]

		if len(events) == 0 || events == nil {
			events = []repos.Event{}
		}

		resp := &responses.EventsResponse{
			Events: events,
			Count:  eventsCount,
		}

		helpers.HttpJson(w, http.StatusOK, resp)
	}
}

func UpdateEvent(eventRepo repos.EventInterface) http.HandlerFunc {
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

func DeleteEvent(eventRepo repos.EventInterface) http.HandlerFunc {
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

// func getUpcomingEvents(eventRepo repos.EventInterface) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		events, err := eventRepo.GetUpcomingEvents()
// 		if err != nil {
// 			helpers.HttpError(w, http.StatusInternalServerError, "wasn't able to get upcoming events, please try again later")
// 			return
// 		}
// 		helpers.HttpJson(w, http.StatusOK, events)
// 	}
// }

func SearchEvents(eventRepo repos.EventInterface, r *http.Request) http.HandlerFunc {
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
		keyword := chi.URLParam(r, "keyword")
		events, err := eventRepo.SearchEvents(keyword)

		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "couldn't search events, please try again later")
			return
		}

		if len(events) < limit*(page-1) {
			helpers.HttpError(w, http.StatusBadRequest, "requested page does not exist")
			return
		}

		eventsCount := len(events)
		startIndex := (page - 1) * limit
		endIndex := min(page*limit, len(events))
		events = events[startIndex:endIndex]

		if len(events) == 0 || events == nil {
			events = []repos.Event{}
		}

		resp := &responses.EventsResponse{
			Events: events,
			Count:  eventsCount,
		}

		helpers.HttpJson(w, http.StatusOK, resp)
	}

}

func AssignEvent(eventRepo repos.EventInterface, userRepo *repos.UserRepository) http.HandlerFunc {
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

		err = userRepo.RemoveOneTicketFromUser(user.ID)
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
