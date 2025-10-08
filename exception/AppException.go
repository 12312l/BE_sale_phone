package exception

type AppException struct {
	ErrorCode ErrorCode
}

func (e *AppException) Error() string {
	return e.ErrorCode.Message
}

func NewAppException(code ErrorCode) *AppException {
	return &AppException{
		ErrorCode: code,
	}
}
