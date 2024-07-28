package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/client/user_service"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/jwt"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

const (
	authURL   = "/api/auth"
	signupURL = "/api/signup"
)

type Handler struct {
	Logger      logging.Logger
	UserService user_service.UserService
	JWTHelper   jwt.Helper
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, authURL, h.Auth)
	router.HandlerFunc(http.MethodPut, authURL, h.Auth)
	router.HandlerFunc(http.MethodPost, signupURL, h.Signup)
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	h.UserService.GetByPhoneNumberAndPassword(r.Context(), "8778", "pass")
	//TODO: implement
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
}
