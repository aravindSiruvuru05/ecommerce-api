package requestresponseutils

import (
	"fmt"
	"haste/pkg/types"
	"net/http"
)

// APIResponse Response format structure.
type APIResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

// SuccessResponse prepares success response format.
// It returns an API response object.
func SuccessResponse(data interface{}, err error) APIResponse {
	return PrepareResponse(data, err, http.StatusOK)
}

// ErrorResponse prepares error response format.
// It returns an API response object.
func ErrorResponse(data interface{}, err error) APIResponse {
	return PrepareResponse(data, err, http.StatusOK)
}

// PrepareResponse Prepares response format.
// It returns APIResponse.
func PrepareResponse(data interface{}, err error, code int, v ...*types.Errors) APIResponse {
	var errors *types.Errors
	if len(v) > 0 {
		errors = v[0]
	}

	r := APIResponse{
		Code: code,
		Data: data,
	}

	if err == nil {
		return r
	}

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
