package service

import "errors"

type serviceErrorImpl struct {
	code    ErrorCode
	message string
	cause   error
}

func (e *serviceErrorImpl) Code() ErrorCode {
	return e.code
}

func (e *serviceErrorImpl) Error() string {
	return e.message
}

func (e *serviceErrorImpl) Cause() error {
	return e.cause
}

func (e *serviceErrorImpl) With(cause error) Error {
	e.cause = cause
	return e
}

func newServiceError(code ErrorCode, message string) Error {
	return &serviceErrorImpl{
		code:    INTERNAL,
		message: message,
		cause:   errors.New(""),
	}
}
