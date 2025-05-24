package repo

import (
	"context"
)

type UserRepository interface {
	GetUsernamesByIds(ctx context.Context, ids []int64) (map[int64]string, error)
	GetUsernameById(ctx context.Context, id int64) (string, error)
}