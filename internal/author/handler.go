package author

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-api/internal/apperror"
	"go-rest-api/internal/handlers"
	"go-rest-api/pkg/logging"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	authorsUrl = "/authors"
	authorUrl  = "/authors/:uuid"
)

type handler struct {
	logger     logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     *logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, authorsUrl, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, authorsUrl, apperror.Middleware(h.CreateAuthor))
	router.HandlerFunc(http.MethodGet, authorUrl, apperror.Middleware(h.GetAuthorByUUID))
	router.HandlerFunc(http.MethodPut, authorUrl, apperror.Middleware(h.UpdateAuthor))
	router.HandlerFunc(http.MethodDelete, authorUrl, apperror.Middleware(h.DeleteAuthor))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	list, err := h.repository.GetList(context.TODO())

	if err != nil {
		w.WriteHeader(400)
		return err
	}

	listBytes, err := json.Marshal(list)
	if err != nil {
		return err
	}

	_, err = w.Write(listBytes)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("API error")
}

func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "not found", "test", "US-000004")
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("update author"))
	h.logger.Info("update author")

	return nil
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("delete author"))
	h.logger.Info("delete author")

	return nil
}
