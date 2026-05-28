package utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/wander1ch/to_do_list_app/internal/core/errors"
)




func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	params := r.URL.Query().Get(key)
	if params == "" {
		return nil, nil
	}

	value, err := strconv.Atoi(params)
	if err != nil {
		return nil, fmt.Errorf("params = %s: invalid query parameter '%s': %v : %w ", params, key, err, core_errors.ErrInvalidArgument)
	}
	return &value, nil
}