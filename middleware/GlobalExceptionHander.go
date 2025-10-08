package middleware

import (
	"net/http"

	"daoduy.com/hoc-golang/exception"
	"daoduy.com/hoc-golang/models/response"
	"github.com/gin-gonic/gin"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypeAny).Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *exception.AppException:
				c.JSON(e.ErrorCode.HTTPCode, response.ApiResponse{
					Code:    e.ErrorCode.Code,
					Message: e.ErrorCode.Message,
				})
			default:
				c.JSON(http.StatusInternalServerError, response.ApiResponse{
					Code:    exception.UncategorizedException.Code,
					Message: err.Error(),
				})
			}
			c.Abort()
		}
	}
}
