package user

import (
	"context"

	"github.com/Tullerpeton/otus-highload/internal/usecase/model"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Save(ctx context.Context, user model.User) (int64, error) {

	return 0, nil
}

func (r *Repository) GetByID(ctx context.Context, userID int64) (*model.User, error) {

	return nil, nil
}
