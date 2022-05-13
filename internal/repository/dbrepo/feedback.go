package dbrepo

import (
	"context"
	"github.com/izzettinozbektas/go-feedback/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *mongoDBRepo) FeedbackCreate(res models.Feedback) (bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res.Created_at = time.Now()
	collection := m.DB.Collection("feedback")

	_, err := collection.InsertOne(ctx,res)
	if err != nil {
		return false,err
	}
	return true, nil
}

func (m *mongoDBRepo) Feedbacks() ([]models.Feedback,error)  {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var feedbacks []models.Feedback
	collection := m.DB.Collection("feedback")

	rows, err := collection.Find(ctx, bson.D{})
	rows.All(ctx,&feedbacks)
	return feedbacks, err

}
