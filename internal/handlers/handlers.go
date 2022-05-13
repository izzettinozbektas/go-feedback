package handlers

import (
	"github.com/izzettinozbektas/go-feedback/internal/repository"
	"github.com/izzettinozbektas/go-feedback/internal/repository/dbrepo"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectionToDB(db *mongo.Database) *Repository {
	return &Repository{
		DB: dbrepo.NewMongoRepo(*db),
	}
}

type Repository struct {
	DB repository.DatabaseRepo
}
