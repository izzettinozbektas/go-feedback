package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/go-feedback/internal/driver"
	"github.com/izzettinozbektas/go-feedback/internal/handlers"
	"net/http"
	"os"
)


func Routes() http.Handler {
	mux := chi.NewRouter()
	connection, err := driver.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	mux.Post("/feedback" , handlers.ConnectionToDB(connection).PostFeedback)
	mux.Get("/feedback" , handlers.ConnectionToDB(connection).GetFeedback)

	return mux
}
