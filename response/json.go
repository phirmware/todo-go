package response

import (
	"encoding/json"
	"net/http"
)

// Response defines the shape of the response object
type Response struct {
	w      http.ResponseWriter
	status int
}

// NewResponse returns the new response
func NewResponse(w http.ResponseWriter, status int) *Response {
	return &Response{
		w:      w,
		status: status,
	}
}

// JSON handles json response
func (r *Response) JSON(data interface{}) {
	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(r.status)
	json.NewEncoder(r.w).Encode(data)
}
