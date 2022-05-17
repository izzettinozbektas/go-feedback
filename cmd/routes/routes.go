package routes

import (
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/go-feedback/internal/handlers"
	"net/http"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/feedback", handlers.Repo.PostFeedback)
	mux.Get("/feedback", handlers.Repo.GetFeedback)

	return mux
}
