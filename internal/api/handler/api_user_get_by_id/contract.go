package api_user_get_by_id

import (
	"context"

	"github.com/Tullerpeton/otus-highload/internal/usecase/model"
)

type repository interface {
	GetByID(ctx context.Context, userID int64) (*model.User, error)
}
