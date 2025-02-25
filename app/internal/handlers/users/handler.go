package users

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/apperror"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/client/user_service"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

const (
	authURL        = "/api/auth"
	signupURL      = "/api/signup"
	assignProducer = "/api/assign-producer"
)

type Handler struct {
	Logger      logging.Logger
	UserService user_service.UserService
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, authURL, apperror.Middleware(h.Auth))
	router.HandlerFunc(http.MethodPut, authURL, apperror.Middleware(h.Auth))
	router.HandlerFunc(http.MethodPost, signupURL, apperror.Middleware(h.Signup))
	router.HandlerFunc(http.MethodPost, assignProducer, apperror.Middleware(h.AssignProducerRole))
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	var token []byte
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()
		var dto user_service.SigninUserDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			return apperror.BadRequestError("failed to decode data")
		}
		u, err := h.UserService.GetByPhoneNumberAndPassword(r.Context(), dto.PhoneNumber, dto.Password)
		if err != nil {
			return err
		}

		token = []byte(u.JWTToken)
		h.Logger.Debug("see token: ", u.JWTToken)

	case http.MethodPut:
		defer r.Body.Close()
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(token)

	return nil
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	var dto user_service.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return apperror.BadRequestError("failed to decode data")
	}
	u, err := h.UserService.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	var response user_service.CreateUserResponse
	response.UUID = u.UUID
	response.JWTToken = u.JWTToken

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	h.Logger.Tracef("token: %s", u.JWTToken)
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBytes)
	return nil
}

func (h *Handler) AssignProducerRole(w http.ResponseWriter, r *http.Request) error {
	return nil
}
