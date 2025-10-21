package exception

import "net/http"

// Common errors
var (
	UncategorizedException = ErrorCode{9999, "Uncategorized exception", http.StatusInternalServerError}
	InvalidKey             = ErrorCode{1001, "Invalid message key", http.StatusBadRequest}
	DatabaseError          = ErrorCode{1013, "Database error", http.StatusInternalServerError}
	InvalidID              = ErrorCode{1014, "Invalid ID format", http.StatusBadRequest}
)