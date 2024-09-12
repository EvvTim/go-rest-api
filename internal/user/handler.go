package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-api/internal/apperror"
	"go-rest-api/internal/handlers"
	"go-rest-api/pkg/logging"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: *logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersUrl, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userUrl, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userUrl, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodDelete, userUrl, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "not found", "test", "US-000004")
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("API error")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "not found", "test", "US-000004")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("update user"))
	h.logger.Info("update user")

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("delete user"))
	h.logger.Info("delete user")

	return nil
}
