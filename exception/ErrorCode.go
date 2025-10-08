package exception

import "net/http"

type ErrorCode struct {
	Code     int
	Message  string
	HTTPCode int
}

var (
	UncategorizedException = ErrorCode{9999, "Uncategorized exception", http.StatusInternalServerError}
	InvalidKey             = ErrorCode{1001, "Invalid message key", http.StatusBadRequest}
	UserExisted            = ErrorCode{1002, "User existed", http.StatusBadRequest}
	UserNotExisted         = ErrorCode{1005, "User not existed", http.StatusNotFound}
	UsernameInvalid        = ErrorCode{1003, "Username must be at least 3 characters", http.StatusBadRequest}
	InvalidPassword        = ErrorCode{1010, "Password must be at least 8 characters", http.StatusBadRequest}
	InvalidFullname        = ErrorCode{1009, "Fullname must be at least 2 and at most 50", http.StatusBadRequest}
	UserNotFound           = ErrorCode{1004, "User not found", http.StatusNotFound}
	Unauthenticated        = ErrorCode{1006, "Unauthenticated", http.StatusUnauthorized}
	Unauthorized           = ErrorCode{1007, "You do not have permission", http.StatusForbidden}
	InvalidDob             = ErrorCode{1008, "Your age must be at least 16", http.StatusBadRequest}
	PasswordIncorrect      = ErrorCode{1011, "Password incorrect", http.StatusBadRequest}
	AddressNotFound        = ErrorCode{1012, "Address not found", http.StatusNotFound}
	DatabaseError          = ErrorCode{1013, "Database error", http.StatusInternalServerError}
	InvalidID              = ErrorCode{1014, "Invalid ID format", http.StatusBadRequest}
)
