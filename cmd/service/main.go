package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Tullerpeton/otus-highload/internal/api/handler/api_login"
	"github.com/Tullerpeton/otus-highload/internal/api/handler/api_user_get_by_id"
	"github.com/Tullerpeton/otus-highload/internal/api/handler/api_user_register"
	"github.com/Tullerpeton/otus-highload/internal/infrastructure/repository/postgres/user"
)

func main() {
	userRepo := user.New()

	apiLoginHandler := api_login.New(userRepo)
	apiUserGetByIdHandler := api_user_get_by_id.New(userRepo)
	apiUserRegisterHandler := api_user_register.New(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/login", apiLoginHandler.Handle)
	r.Post("/user/register", apiUserRegisterHandler.Handle)
	r.Get("/user/get/{id}", apiUserGetByIdHandler.Handle)
	http.ListenAndServe(":3000", r)
}
