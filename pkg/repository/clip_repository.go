package repository

import (
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type ClipRepository interface {
	NewQuery(tx *gorm.DB) ClipQuery
}

type ClipQuery interface {
	Where(cond *model.Clip) ClipQuery

	Limit(limit int) ClipQuery
	Latest() ClipQuery

	Create(clip *model.Clip) error
	Find() (*model.Clip, error)
	FindAll() ([]*model.Clip, error)
	Save(clip *model.Clip) error
	Delete() (int, error)
}

func NewClipRepository() ClipRepository {
	return &ClipRepositoryImpl{}
}

type ClipRepositoryImpl struct{}

func (r *ClipRepositoryImpl) NewQuery(tx *gorm.DB) ClipQuery {
	return &ClipQueryImpl{Tx: tx}
}

type ClipQueryImpl struct {
	Tx *gorm.DB
}

func (q *ClipQueryImpl) Where(cond *model.Clip) ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *ClipQueryImpl) Limit(limit int) ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Limit(limit)}
}

func (q *ClipQueryImpl) Latest() ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Order("created_at desc")}
}

func (q *ClipQueryImpl) Create(clip *model.Clip) error {
	return q.Tx.Create(clip).Error
}

func (q *ClipQueryImpl) Find() (*model.Clip, error) {
	res := &model.Clip{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *ClipQueryImpl) FindAll() ([]*model.Clip, error) {
	res := make([]*model.Clip, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *ClipQueryImpl) Save(clip *model.Clip) error {
	return q.Tx.Save(clip).Error
}

func (q *ClipQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Clip{})
	return (int)(res.RowsAffected), res.Error
}
