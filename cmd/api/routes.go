package main

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	handler := AddContentType(mux)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"author":     "inuoshios",
			"github_url": "https://github.com/inuoshios",
			"message":    "api is running...",
		}
		marshalResponse, _ := json.MarshalIndent(response, "", "\t")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write(marshalResponse)
		if err != nil {
			http.Error(w, "An error occurred, please try again", http.StatusBadRequest)
			return
		}
	})

	mux.HandleFunc("POST /user/signup", app.handlers.CreateUser)
	mux.HandleFunc("POST /user/signin", app.handlers.SignIn)
	mux.Handle("GET /user/get-users", Authenticate(HandlerFunc(app.handlers.GetUsers)))

	// boards
	mux.Handle("POST /user/board", Authenticate(HandlerFunc(app.handlers.CreateBoard)))
	mux.Handle("POST /user/board/column", Authenticate(HandlerFunc(app.handlers.CreateBoardColumn)))

	return handler
}
