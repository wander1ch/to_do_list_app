package users_http_transport

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/wander1ch/to_do_list_app/internal/core/domain"
	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	core_http_request "github.com/wander1ch/to_do_list_app/internal/core/transport/http/request"
	core_http_response "github.com/wander1ch/to_do_list_app/internal/core/transport/http/response"
	core_http_types "github.com/wander1ch/to_do_list_app/internal/core/transport/http/types"
	"github.com/wander1ch/to_do_list_app/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	Name core_http_types.Nullable[string] `json:"full_name"`
	Phone core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchUserRequest) Validate() error {
	if r.Name.Set {
		if r.Name.Value == nil {
			return fmt.Errorf("Full_name cannot be null")
		}
		NameLen := len([]rune(*r.Name.Value))
		if NameLen < 2 || NameLen > 100 {
			return fmt.Errorf("Full_name must be between 2 and 100 characters")
		}
	}

	if r.Phone.Set {
		if r.Phone.Value != nil {
			phoneLen := len([]rune(*r.Phone.Value))
			if phoneLen < 10 || phoneLen > 15 {
				return fmt.Errorf("Phone number must be between 10 and 15 characters")
			}
			re := regexp.MustCompile(`^\+[0-9]{10,15}$`)
			if !re.MatchString(*r.Phone.Value) {
				return fmt.Errorf("Invalid phone number format")
			}
		}
	}
	return nil
}


type PatchUserResponse UserDTOResponse

func (h *UserHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid user ID in path")
		return
	}

	var req PatchUserRequest
	if err := core_http_request.DecodeAndValidateJSONRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "invalid request body")
		return
	}

	UserPatch := userPatchFromRequest(req)  

	userDomain, err := h.userService.PatchUser(ctx, userID, UserPatch) 
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(http.StatusOK, response)

	log.Debug(fmt.Sprintf("PatchUser request fields: Name=%v, Phone=%v", req.Name, req.Phone))
	
	rw.WriteHeader(http.StatusOK)
}

func userPatchFromRequest(req PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		Name: req.Name.ToDomain(),
		Phone: req.Phone.ToDomain(),
	}
}