package users_http_transport

import (
	"net/http"

	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	core_http_response "github.com/wander1ch/to_do_list_app/internal/core/transport/http/response"
	"github.com/wander1ch/to_do_list_app/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

func (h *UserHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"fatal error occurred while parsing id query param",
		)
		return
	}
	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"fatal error occurred while fetching user",
		)
		return
	}
	response := GetUserResponse(userDTOFromDomain(user))
	responseHandler.JSONResponse(http.StatusOK, response)
}