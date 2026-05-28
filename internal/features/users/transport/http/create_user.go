package users_http_transport

import (
	"net/http"

	domain "github.com/wander1ch/to_do_list_app/internal/core/domain"
	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	core_http_request "github.com/wander1ch/to_do_list_app/internal/core/transport/http/request"
	core_http_response "github.com/wander1ch/to_do_list_app/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=2,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserDTOResponse

func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	log.Debug("handling CreateUser request")
	
	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateJSONRequest(r, &request); err != nil {

		responseHandler.ErrorResponse(err, "failed to decode and validate CreateUser request")
			
		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.userService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(http.StatusCreated, response)


}
	
func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}

