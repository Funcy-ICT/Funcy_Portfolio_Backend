package middleware

import (
	"backend/app/interfaces/response"
	"backend/app/packages/utils/auth"

	"github.com/mileusna/useragent"

	"context"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authz := r.Header.Get("Authorization")
		if authz == "" || !strings.HasPrefix(authz, "Bearer ") {
			_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "authentication failure")
			return
		}

		token := strings.TrimPrefix(authz, "Bearer ")

		ua := useragent.Parse(r.UserAgent())
		switch {
		case ua.Mobile == true:
			userId, err := auth.VerifyMobileUserToken(token)
			if err != nil {
				_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "authentication failure")
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), "user_id", userId))
			r = r.WithContext(context.WithValue(r.Context(), "user_id", userId))
		case ua.Desktop == true || ua.Name == "PostmanRuntime":
			userId, err := auth.VerifyUserToken(token)
			if err != nil {
				_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "authentication failure")
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), "user_id", userId))
		default:
			_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "authentication failure")
			return
		}
		next.ServeHTTP(w, r)
	})
}
