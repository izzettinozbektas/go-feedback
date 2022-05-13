package handlers

import (
	"github.com/izzettinozbektas/go-feedback/internal/models"
	"github.com/izzettinozbektas/go-feedback/internal/response"
	"net/http"
)

func(m *Repository) PostFeedback(w http.ResponseWriter, r *http.Request){
	var FeedbackRequestData  models.Feedback
	FeedbackRequestData.Feedback = r.FormValue("feedback")

	_, err := m.DB.FeedbackCreate(FeedbackRequestData)
	if err != nil {
		response.Write(w, response.Error("Kayıt Yapılamadı", map[string]string{"error": err.Error()}),response.Code(http.StatusBadRequest))
	}
	response.Write(w,response.Success("işlem başarılı",nil),response.Code(http.StatusOK))
}

func(m *Repository) GetFeedback(w http.ResponseWriter, r *http.Request){
	var feedbacks []models.Feedback
	feedbacks, err := m.DB.Feedbacks()
	if err != nil {
		response.Write(w,response.Error("İşlem başarısız", map[string]string{"error": err.Error()}),response.Code(http.StatusBadRequest))
	}
	response.Write(w,response.Success("",feedbacks),response.Code(http.StatusOK))
}