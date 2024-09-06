package property

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/apperror"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/client/property_service"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/handlers"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

const (
	propertysURL = "/api/property"       //getAll, create
	propertyURL  = "/api/property/:uuid" //getById update delete
	producerRole = "producer"
)

type Handler struct {
	Logger          logging.Logger
	PropertyService property_service.PropertyService
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, propertysURL, apperror.Middleware(h.GetAllProperty))
	router.HandlerFunc(http.MethodGet, propertyURL, apperror.Middleware(h.GetPropertyByUUID))
	router.HandlerFunc(http.MethodPost, propertysURL, handlers.AuthRoleMiddleware(producerRole, apperror.Middleware(h.CreateProperty)))
}

func (h *Handler) GetAllProperty(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	property, err := h.PropertyService.GetAll(r.Context())
	if err != nil {
		return err
	}

	response, err := json.Marshal(property)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

	return nil
}

func (h *Handler) GetPropertyByUUID(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	params := httprouter.ParamsFromContext(r.Context())
	uuid := params.ByName("uuid")

	property, err := h.PropertyService.GetByUUID(r.Context(), uuid)
	if err != nil {
		return err
	}

	response, err := json.Marshal(property)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

	return nil
}

func (h *Handler) CreateProperty(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	var dto property_service.CreatePropertyDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return apperror.BadRequestError("failed to decode data")
	}
	propertyUUID, err := h.PropertyService.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	var response property_service.CreatePropertyResponse
	response.UUID = propertyUUID

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	h.Logger.Tracef("created property with uuid: %s", propertyUUID)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)

	return nil
}
