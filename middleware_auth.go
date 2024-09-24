package main

import (
	"log"
	"net/http"

	"github.com/hawkaii/rssagg/internal/auth"
	"github.com/hawkaii/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find API key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			log.Println(err)
			respondWithError(w, http.StatusUnauthorized, "Couldn't find user")
			return
		}

		handler(w, r, user)
	}
}
