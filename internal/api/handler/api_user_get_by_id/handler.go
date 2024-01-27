package api_user_get_by_id

import (
	"net/http"
)

type Response struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Birthdate  string `json:"birthdate"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
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

}
