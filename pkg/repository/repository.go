package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

func NotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func AlreadyExistsError(err error) bool {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		return false
	}
	return pqErr.Code == "23505"
}
