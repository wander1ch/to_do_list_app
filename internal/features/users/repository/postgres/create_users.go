package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/wander1ch/to_do_list_app/internal/core/domain"
)


func (r *UserRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `
		INSERT INTO todolist.users (name, phone)
		VALUES ($1, $2)
		RETURNING id, version, name, phone
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version, 
		&userModel.FullName, 
		&userModel.PhoneNumber)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to scan created user: %w", err)
	}
	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,	
		userModel.PhoneNumber,
	)
	return userDomain, nil
}