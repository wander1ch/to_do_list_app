package users_postgres_repository

import (
	"context"
	"fmt"

	core_domain "github.com/wander1ch/to_do_list_app/internal/core/domain"
)

func (r *UserRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, name, phone
		FROM todolist.users
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.pool.Query(
		ctx, 
		query, 
		limit, 
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel
		err := rows.Scan(
			&userModel.ID,
			&userModel.Version, 
			&userModel.FullName, 
			&userModel.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		userModels = append(userModels, userModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user rows: %w", err)
	}

	userDomains := userDomainFromModel(userModels)
	return userDomains, nil
}