package user

import (
	"awesomeProject/internal/handlers"
	"awesomeProject/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, h.GetList)
	router.POST(userURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	responseText := []byte("this is list of users")
	if _, err := w.Write(responseText); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)

	responseText := []byte("Create User Route")
	if _, err := w.Write(responseText); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)

	responseText := []byte("Get By UUID User Route")
	if _, err := w.Write(responseText); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)

	responseText := []byte("Update User Route")
	if _, err := w.Write(responseText); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)

	responseText := []byte("Partially Update User Route")
	if _, err := w.Write(responseText); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)

	responseText := []byte("Delete User Route")
	if _, err := w.Write(responseText); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
