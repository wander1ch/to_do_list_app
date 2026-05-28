package users_service

import (
	"context"
	"fmt"

	"github.com/wander1ch/to_do_list_app/internal/core/domain"
)
func (s *UserService) PatchUser(
	ctx context.Context,
	userID int,
	patch domain.UserPatch,
) (domain.User, error) {
	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to patch user: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("failed to apply patch: %w", err)
	}
	patchedUser, err :=s.userRepo.PatchUser(ctx, userID, user) 
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to patch user: %w", err)
	}

	return patchedUser, err
}
