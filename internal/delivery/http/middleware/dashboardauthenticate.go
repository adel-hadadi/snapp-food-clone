package middleware

import (
	"context"
	"net/http"
	"strings"

	tokenservice "snapp-food/internal/service/token"
)

const SellerCtxKey = "user"

func DashboardAuthenticate(tokenSvc tokenservice.Service) func(http.Handler) http.Handler {
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

			claims, err := tokenSvc.GetClaims(tokenString[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// TODO: check token is for an storage
			if claims.User.Role != "seller" {
				w.WriteHeader(http.StatusForbidden)
			}

			ctx := context.WithValue(r.Context(), SellerCtxKey, claims.User)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
