package repo

import (
	"context"
	"fmt"

	"github.com/Ryan-Gosusluging/go-common-forum/postgres"
	"github.com/rs/zerolog"
)

type userRepository struct {
	pg  *postgres.Postgres
	log *zerolog.Logger
}

func New(pg *postgres.Postgres, log *zerolog.Logger) UserRepository {
	return &userRepository{pg, log}
}

func (r *userRepository) GetUsernamesByIds(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		r.log.Warn().Str("op", "UserRepository.GetUsernamesByIds").Msg("empty ids")
		return make(map[int64]string), nil
	}

	rows, err := r.pg.Pool.Query(ctx, "SELECT id, username FROM users WHERE id = ANY($1)", ids)
	if err != nil {
		r.log.Error().Err(err).Str("op", "UserRepository.GetUsernamesByIds").Msg("failed to get usernames by ids")
		return nil, fmt.Errorf("UserRepository - GetUsernamesByIds - r.pg.Pool.Query: %w", err)
	}
	defer rows.Close()

	usernames := make(map[int64]string, len(ids))
	for rows.Next() {
		var id int64
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			return nil, fmt.Errorf("UserRepository - GetUsernamesByIds - rows.Scan: %w", err)
		}
		usernames[id] = username
	}

	return usernames, nil
}

func (r *userRepository) GetUsernameById(ctx context.Context, id int64) (string, error) {
	row := r.pg.Pool.QueryRow(ctx, "SELECT username FROM users WHERE id = $1", id)
	var username string
	if err := row.Scan(&username); err != nil {
		r.log.Error().Err(err).Str("op", "UserRepository.GetUsernameById").Msg("failed to get username by id")
		return "", fmt.Errorf("UserRepository - GetUsernameById - row.Scan: %w", err)
	}

	return username, nil
}