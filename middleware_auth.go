package main

import (
	"fmt"
	"github.com/anmol-pal/go-tutorial/internal/auth"
	"github.com/anmol-pal/go-tutorial/internal/database"
	"net/http"
)

type authedhandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error Authenticating: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserFromAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Could not fetch User: %v", err))
			return
		}
		handler(w, r, user)
	}
}
