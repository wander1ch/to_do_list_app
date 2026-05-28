package users_service

import (
	"context"
	"fmt"

	core_domain "github.com/wander1ch/to_do_list_app/internal/core/domain"
)




func (s *UserService) CreateUser(
	ctx context.Context,
	user core_domain.User,
) (core_domain.User, error) {
	if err := user.Validate(); err != nil {
		return core_domain.User{}, fmt.Errorf("invalid user: %w", err)
	}

	user, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}