package users_http_transport

import (
	"fmt"
	"net/http"

	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	core_http_response "github.com/wander1ch/to_do_list_app/internal/core/transport/http/response"
	"github.com/wander1ch/to_do_list_app/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UserHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	limit, offset, err := getLimitOffsetQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get query parameters")
		return
	}
	usersDomain, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response :=GetUsersResponse(userDTOsFromDomains(usersDomain))
	responseHandler.JSONResponse(http.StatusOK, response)	
}

func getLimitOffsetQueryParam(r *http.Request) (*int, *int, error) {
	limit, err := utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get 'limit' query parameter: %w", err)
	}
	
	offset, err := utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get 'offset' query parameter: %w", err)
	}
	return limit, offset, nil
}