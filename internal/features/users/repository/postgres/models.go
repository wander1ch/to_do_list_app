package users_postgres_repository

import "github.com/wander1ch/to_do_list_app/internal/core/domain"

type UserModel struct {
	ID int
	Version int
	FullName string
	PhoneNumber *string
}

func userDomainFromModel(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,	
			user.PhoneNumber,
		)
	}
	return userDomains
}