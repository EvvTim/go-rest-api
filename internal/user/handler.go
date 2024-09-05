package user

import (
	"github.com/EvvTim/go-rest-api/internal/handlers"
	"github.com/EvvTim/go-rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userUrl  = "/user/:uuid"
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
	router.GET(userUrl, h.GetByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userUrl, h.UpdateUser)
	router.PATCH(userUrl, h.PartiallyUpdateUser)
	router.DELETE(userUrl, h.DeleteUser)

}

func (h *handler) GetList(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	h.logger.Info("get user list")
	w.WriteHeader(200)
	w.Write([]byte("user list"))
}

func (h *handler) GetByUUID(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	h.logger.Info("get user by uuid")
	w.WriteHeader(200)
	w.Write([]byte("user"))
}

func (h *handler) CreateUser(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	h.logger.Info("create user")
	w.WriteHeader(201)
	w.Write([]byte("create user"))
}

func (h *handler) UpdateUser(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	h.logger.Info("update user")
	w.WriteHeader(200)
	w.Write([]byte("update user"))
}

func (h *handler) PartiallyUpdateUser(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	h.logger.Info("partially update user")
	w.WriteHeader(200)
	w.Write([]byte("update user"))
}

func (h *handler) DeleteUser(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	h.logger.Info("delete user")
	w.WriteHeader(204)
	w.Write([]byte("delete user"))
}
