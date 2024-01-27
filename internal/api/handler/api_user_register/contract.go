package api_user_register

import (
	"context"

	"github.com/Tullerpeton/otus-highload/internal/usecase/model"
)

type repository interface {
	Save(ctx context.Context, user model.User) (int64, error)
}
