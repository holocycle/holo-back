package repository

import (
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type LiverRepository interface {
	NewQuery(tx *gorm.DB) LiverQuery
}

type LiverQuery interface {
	Where(cond *model.Liver) LiverQuery

	Create(liver *model.Liver) error
	Find() (*model.Liver, error)
	FindAll() ([]*model.Liver, error)
	Save(liver *model.Liver) error
	Delete() (int, error)
}

func NewLiverRepository() LiverRepository {
	return &LiverRepositoryImpl{}
}

type LiverRepositoryImpl struct{}

func (r *LiverRepositoryImpl) NewQuery(tx *gorm.DB) LiverQuery {
	return &LiverQueryImpl{Tx: tx}
}

type LiverQueryImpl struct {
	Tx *gorm.DB
}

func (q *LiverQueryImpl) Where(cond *model.Liver) LiverQuery {
	return &LiverQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *LiverQueryImpl) Create(liver *model.Liver) error {
	return q.Tx.Create(liver).Error
}

func (q *LiverQueryImpl) Find() (*model.Liver, error) {
	res := &model.Liver{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *LiverQueryImpl) FindAll() ([]*model.Liver, error) {
	res := make([]*model.Liver, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *LiverQueryImpl) Save(liver *model.Liver) error {
	return q.Tx.Save(liver).Error
}

func (q *LiverQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Liver{})
	return (int)(res.RowsAffected), res.Error
}
