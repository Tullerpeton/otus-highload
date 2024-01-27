package api_login

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
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

	fmt.Println(reqBody)
}
