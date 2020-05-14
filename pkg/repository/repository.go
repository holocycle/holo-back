package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func NotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(errors.Cause(err))
}

func AlreadyExistsError(err error) bool {
	pqErr, ok := errors.Cause(err).(*pq.Error)
	if !ok {
		return false
	}
	return pqErr.Code == "23505"
}

func newErr(err error) error {
	if err == nil {
		return nil
	}
	return errors.WithStack(err)
}
