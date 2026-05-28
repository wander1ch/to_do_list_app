package utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/wander1ch/to_do_list_app/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	passValue := r.PathValue(key)
	if passValue == "" {
		return 0, fmt.Errorf("path value for key %s is empty, err: %w", key, core_errors.ErrInvalidArgument)
	}

	intValue, err := strconv.Atoi(passValue)
	if err != nil {
		return 0, fmt.Errorf("failed to convert path value  = %s for key %s to int: %w", passValue, key, core_errors.ErrInvalidArgument)
	}
	return intValue, nil
}