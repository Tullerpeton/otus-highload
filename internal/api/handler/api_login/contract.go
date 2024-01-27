package api_login

import (
	"context"

	"github.com/Tullerpeton/otus-highload/internal/usecase/model"
)

type repository interface {
	GetByID(ctx context.Context, userID int64) (*model.User, error)
}
