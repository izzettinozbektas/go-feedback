package repository

import "github.com/izzettinozbektas/go-feedback/internal/models"

type DatabaseRepo interface {
	FeedbackCreate(req models.Feedback) (bool, error)
	Feedbacks() ([]models.Feedback,error)
}
