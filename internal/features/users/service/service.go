package users_service

import (
	"context"

	"github.com/wander1ch/to_do_list_app/internal/core/domain"
	core_domain "github.com/wander1ch/to_do_list_app/internal/core/domain"
)

type UserService struct {
	userRepo UserRepository
}

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		user core_domain.User,
	) (core_domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]core_domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
	) (core_domain.User, error) 
	DeleteUser(
		ctx context.Context,
		id int,
	) error
	PatchUser(
		ctx context.Context,
		id int,
		user domain.User,
	) (domain.User, error)
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}