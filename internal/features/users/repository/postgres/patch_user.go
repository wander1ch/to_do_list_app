package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/wander1ch/to_do_list_app/internal/core/domain"
	core_errors "github.com/wander1ch/to_do_list_app/internal/core/errors"
)

func (r *UserRepository) PatchUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todolist.users
		SET 
			name = $1, 
			phone = $2,
			version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING 
			id, 
			version, 
			name, 
			phone
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		user.FullName,
		user.PhoneNumber,
		id,
		user.Version,
	)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version, 
		&userModel.FullName, 
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id %d conflict: %w", id, core_errors.ErrConflict)

		}
		return domain.User{}, fmt.Errorf("failed to scan patched user: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,	
		userModel.PhoneNumber,
	)
	return userDomain, nil
}