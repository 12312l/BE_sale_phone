package exception

import "net/http"


var (
	UserExisted      = ErrorCode{1002, "User existed", http.StatusBadRequest}
	UserNotExisted   = ErrorCode{1005, "User not existed", http.StatusNotFound}
	UserNotFound     = ErrorCode{1004, "User not found", http.StatusNotFound}
	UsernameInvalid  = ErrorCode{1003, "Username must be at least 3 characters", http.StatusBadRequest}
	InvalidPassword  = ErrorCode{1010, "Password must be at least 8 characters", http.StatusBadRequest}
	InvalidFullname  = ErrorCode{1009, "Fullname must be at least 2 and at most 50", http.StatusBadRequest}
	InvalidDob       = ErrorCode{1008, "Your age must be at least 16", http.StatusBadRequest}
)