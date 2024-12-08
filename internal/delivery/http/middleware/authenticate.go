package middleware

import (
	"context"
	"net/http"
	tokenservice "snapp-food/internal/service/token"
	"strings"
)

const authorizationKey = "Authorization"

func Authenticate(tokenSvc tokenservice.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get(authorizationKey)
			if header == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			tokenString := strings.Split(header, " ")
			if len(tokenString) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// TODO: claim token
			claims, err := tokenSvc.Claim(r.Context(), tokenString[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", claims[tokenservice.UserID])

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
