package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	resp "github.com/inuoshios/little-jira/internal/helpers"
	"github.com/inuoshios/little-jira/internal/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")
			if len(authToken) == 0 {
				resp.ErrorJSON(w, utils.ErrAuthHeader, http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authToken, " ")
			if len(bearerToken) < 2 {
				resp.ErrorJSON(w, utils.ErrInvalidAuthHeader, http.StatusUnauthorized)
				return
			}

			if bearerToken[0] != "Bearer" {
				resp.ErrorJSON(w, utils.ErrUnsupportedAuthType, http.StatusUnauthorized)
				return
			}

			accesssToken := bearerToken[1]
			payload, err := utils.VerifyToken(accesssToken)
			if err != nil {
				resp.ErrorJSON(w, fmt.Errorf("%w", err), http.StatusUnauthorized)
			}

			ctx := context.WithValue(context.Background(), "userId", payload.UserID)
			next.ServeHTTP(w, r.Clone(ctx))
		},
	)
}
