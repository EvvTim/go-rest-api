package user

import (
	"github.com/EvvTim/go-rest-api/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersUrl, h.GetList)
	router.POST(usersUrl, h.CreateUser)
	router.GET(userUrl, h.GetUserByUUID)
	router.PUT(userUrl, h.UpdateUser)
	router.DELETE(userUrl, h.DeleteUser)
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user list"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("add user"))
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("find all users"))

}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("delete user"))
}
