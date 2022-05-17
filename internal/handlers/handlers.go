package handlers

import (
	"github.com/izzettinozbektas/go-feedback/internal/repository"
	"github.com/izzettinozbektas/go-feedback/internal/repository/dbrepo"
	"go.mongodb.org/mongo-driver/mongo"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
}

func NewRepo(db *mongo.Database) *Repository {
	return &Repository{
		DB: dbrepo.NewMongoRepo(*db),
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}
