package repository

import "github.com/jinzhu/gorm"

func NotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
