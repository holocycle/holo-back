package service

type ErrorCode int64

const (
	INTERNAL ErrorCode = iota
	NOTFOUND
	OUTOFRANGE
	FORBIDDEN
)

type Error interface {
	Code() ErrorCode
	Error() string
	Cause() error
	With(cause error) Error
}

func NewInternalError(message string) Error {
	return newServiceError(INTERNAL, message)
}

func NewNotFoundError(message string) Error {
	return newServiceError(NOTFOUND, message)
}

func NewOutOfRanegError(message string) Error {
	return newServiceError(OUTOFRANGE, message)
}

func NewForbiddenError(message string) Error {
	return newServiceError(FORBIDDEN, message)
}
