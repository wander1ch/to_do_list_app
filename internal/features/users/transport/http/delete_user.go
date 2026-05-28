package users_http_transport

import (
	"net/http"

	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	core_http_response "github.com/wander1ch/to_do_list_app/internal/core/transport/http/response"
	"github.com/wander1ch/to_do_list_app/internal/core/transport/http/utils"
)

func (h *UserHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()	
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	
	userID, err := utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid user ID in path")
		return
	}
	if err := h.userService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}
	responseHandler.NoContentResponse()
}

