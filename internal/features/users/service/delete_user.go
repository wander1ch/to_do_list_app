package users_service

import (
	"context"
	"fmt"

)

func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	if err := s.userRepo.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}