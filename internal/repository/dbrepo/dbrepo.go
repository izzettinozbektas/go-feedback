package dbrepo

import (
	"github.com/izzettinozbektas/go-feedback/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBRepo struct {
	DB mongo.Database
}

func NewMongoRepo(conn mongo.Database) repository.DatabaseRepo {
	return &mongoDBRepo{
		DB: conn,
	}
}