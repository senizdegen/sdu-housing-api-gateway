package property

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/apperror"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/client/property_service"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

const (
	propertysURL = "api/property"       //getAll, create
	propertyURL  = "api/property/:uuid" //getById update delete

)

type Handler struct {
	Logger          logging.Logger
	PropertyService property_service.PropertyService
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, propertysURL, apperror.Middleware(h.GetAll))
	router.HandlerFunc(http.MethodPost, propertysURL, apperror.Middleware(h.Create))
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	return nil
}
