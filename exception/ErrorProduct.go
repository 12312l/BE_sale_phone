package exception

import "net/http"


var (
	CategoryNotFound = ErrorCode{1015, "Category not found", http.StatusNotFound}
	ProductNotFound  = ErrorCode{1016, "Product not found", http.StatusNotFound}
)