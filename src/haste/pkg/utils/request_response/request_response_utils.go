package requestresponseutils

import (
	"fmt"
	"haste/pkg/types"
)

// APIResponse Response format structure.
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

// PrepareResponse Prepares response format.
// It returns APIResponse.
func PrepareResponse(data interface{}, err error, message string, v ...*types.Errors) APIResponse {
	var errors *types.Errors
	if len(v) > 0 {
		errors = v[0]
	}

	r := APIResponse{
		Message: message,
		Data:    data,
	}

	if err == nil {
		return r
	}
	fmt.Println("errrr", err, err.Error())
	if errStr := err.Error(); errStr != "" {
		r.Errors = errStr
	} else {
		if errors.ErrorMsg == "" {
			errors.ErrorMsg = "Bad request."
		}
		r.Errors = errors
	}

	return r
}

func init() {
	fmt.Println("req res init util")
}
