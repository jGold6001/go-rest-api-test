package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapi-lesson/internal/handlers"
	"restapi-lesson/pkg/logging"
)

var _ handlers.Handler = &handler{}

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
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.ParticularUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("First log")
	writer.WriteHeader(200)
	_, err := writer.Write([]byte("list users"))
	if err != nil {
		return
	}
}

func (h *handler) GetUserByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(200)
	_, err := writer.Write([]byte("get user by id"))
	if err != nil {
		return
	}
}

func (h *handler) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(201)
	_, err := writer.Write([]byte("create user"))
	if err != nil {
		return
	}
}

func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(204)
	_, err := writer.Write([]byte("update user"))
	if err != nil {
		return
	}
}

func (h *handler) ParticularUpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(204)
	_, err := writer.Write([]byte("part update user"))
	if err != nil {
		return
	}
}

func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(204)
	_, err := writer.Write([]byte("delete user"))
	if err != nil {
		return
	}
}
