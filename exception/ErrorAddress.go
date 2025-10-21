package exception

import "net/http"

type ErrorCode struct {
	Code     int
	Message  string
	HTTPCode int
}

var (
	AddressNotFound = ErrorCode{1012, "Address not found", http.StatusNotFound}
)