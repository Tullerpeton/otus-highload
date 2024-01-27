package api_user_register

import (
	"encoding/json"
	"fmt"
	"github.com/Tullerpeton/otus-highload/internal/usecase/model"
	"net/http"
)

type RequestBody struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Birthdate  string `json:"birthdate"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	Password   string `json:"password"`
}

type Response struct {
	UserId string `json:"user_id"`
}

type Handler struct {
	repo repository
}

func New(repo repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.repo.Save(r.Context(), model.User{
		FirstName:  reqBody.FirstName,
		SecondName: reqBody.SecondName,
		Birthdate:  reqBody.Birthdate,
		Biography:  reqBody.Biography,
		City:       reqBody.City,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(userID)
	http.Error(w, err.Error(), http.StatusOK)
}
