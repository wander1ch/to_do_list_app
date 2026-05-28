package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/wander1ch/to_do_list_app/internal/core/errors"
)

func (r *UserRepository) DeleteUser(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	
	query := `DELETE FROM todolist.users WHERE id = $1`
	comTag, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to execute delete user query: %w", err)
	}
	if comTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found with ID %d:%w", userID, core_errors.ErrNotFound)
	}
	return nil
}
