package users_http_transport

import (
	"context"
	"net/http"

	core_http_server "github.com/wander1ch/to_do_list_app/internal/core/transport/http/server"
	domain "github.com/wander1ch/to_do_list_app/internal/core/domain"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit, offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		userID int,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		userID int,
	) error
	PatchUser(
		ctx context.Context,
		userID int,
		patch domain.UserPatch,
	) (domain.User, error)
}

func NewUserHTTPHandler(userService UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/users",
			Handler: h.CreateUser,
		},
		{
			Method: http.MethodGet,
			Path: "/users",
			Handler: h.GetUsers,
		},
		{
			Method: http.MethodGet,
			Path: "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method: http.MethodDelete,
			Path: "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method: http.MethodPatch,
			Path: "/users/{id}",
			Handler: h.PatchUser,
		},


	}

}