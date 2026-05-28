package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/wander1ch/to_do_list_app/internal/core/domain"
	core_domain "github.com/wander1ch/to_do_list_app/internal/core/domain"
	core_errors "github.com/wander1ch/to_do_list_app/internal/core/errors"
)

func (r *UserRepository) GetUser(
	ctx context.Context,
	id int,
) (core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, name, phone
		FROM todolist.users
		WHERE id = $1
	`
	row := r.pool.QueryRow(ctx, query, id)
	
	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version, 
		&userModel.FullName, 
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id %d not found: %w", id, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("failed to scan user: %w", err)
	}
	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,	
		userModel.PhoneNumber,
	)
	return userDomain, nil
}