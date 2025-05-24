package usecase

import (
	"context"
)

type UserUsecase interface {
	GetUsernamesByIds(ctx context.Context, ids []int64) (map[int64]string, error)
	GetUsernameById(ctx context.Context, id int64) (string, error)
}