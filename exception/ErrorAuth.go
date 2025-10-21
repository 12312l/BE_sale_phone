package exception

import "net/http"


var (
	Unauthenticated   = ErrorCode{1006, "Unauthenticated", http.StatusUnauthorized}
	Unauthorized      = ErrorCode{1007, "You do not have permission", http.StatusForbidden}
	PasswordIncorrect = ErrorCode{1011, "Password incorrect", http.StatusBadRequest}
	InCorrectPassword = ErrorCode{1017, "Incorrect password", http.StatusBadRequest}
)