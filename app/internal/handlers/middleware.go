package handlers

import (
	"net/http"

	"github.com/cristalhq/jwt/v3"
)

type UserClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

func AuthRoleMiddleware(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// logger := logging.GetLogger()
		// authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		// if len(authHeader) != 2 {
		// 	logger.Error("Malformer token")
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte("malformed token"))
		// 	return
		// }
		// logger.Debug("create jwt verifier")
		// jwtToken := authHeader[1]
		// key := []byte(config.GetConfig().JWT.Secret)
		// verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
		// if err != nil {
		// 	unauthorized(w, err)
		// 	return
		// }
		// logger.Debug("parse and verify token")
		// token, err := jwt.ParseAndVerifyString(jwtToken, verifier)
		// if err != nil {
		// 	unauthorized(w, err)
		// 	return
		// }
		// logger.Debug("parse user claims")
		// var uc UserClaims
		// err = json.Unmarshal(token.RawClaims(), &uc)
		// if err != nil {
		// 	unauthorized(w, err)
		// 	return
		// }
		// if valid := uc.IsValidAt(time.Now()); !valid {
		// 	logger.Error("token has been expired")
		// 	unauthorized(w, err)
		// 	return
		// }

		// if uc.Role != requiredRole {
		// 	logger.Error("user do not have a permission to this action")
		// 	w.WriteHeader(http.StatusForbidden)
		// 	return
		// }
		next(w, r)
	}
}

// func unauthorized(w http.ResponseWriter, err error) {
// 	logging.GetLogger().Error(err)
// 	w.WriteHeader(http.StatusUnauthorized)
// 	w.Write([]byte("unauthorized"))
// }
