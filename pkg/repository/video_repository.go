package repository

import (
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type VideoRepository interface {
	NewQuery(tx *gorm.DB) VideoQuery
}

type VideoQuery interface {
	Where(cond *model.Video) VideoQuery

	Create(video *model.Video) error
	Find() (*model.Video, error)
	FindAll() ([]*model.Video, error)
	Save(video *model.Video) error
	Delete() (int, error)
}

func NewVideoRepository() VideoRepository {
	return &VideoRepositoryImpl{}
}

type VideoRepositoryImpl struct{}

func (r *VideoRepositoryImpl) NewQuery(tx *gorm.DB) VideoQuery {
	return &VideoQueryImpl{Tx: tx}
}

type VideoQueryImpl struct {
	Tx *gorm.DB
}

func (q *VideoQueryImpl) Where(cond *model.Video) VideoQuery {
	return &VideoQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *VideoQueryImpl) Create(video *model.Video) error {
	return q.Tx.Create(video).Error
}

func (q *VideoQueryImpl) Find() (*model.Video, error) {
	res := &model.Video{}
	if err := q.Tx.First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *VideoQueryImpl) FindAll() ([]*model.Video, error) {
	res := make([]*model.Video, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *VideoQueryImpl) Save(video *model.Video) error {
	return q.Tx.Save(video).Error
}

func (q *VideoQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Video{})
	return (int)(res.RowsAffected), res.Error
}
