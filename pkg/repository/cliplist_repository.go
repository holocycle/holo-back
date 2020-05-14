package repository

import (
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type CliplistRepository interface {
	NewQuery(tx *gorm.DB) CliplistQuery
}

type CliplistQuery interface {
	Where(cond *model.Cliplist) CliplistQuery
	JoinChannel() CliplistQuery

	Create(Cliplist *model.Cliplist) error
	Find() (*model.Cliplist, error)
	FindAll() ([]*model.Cliplist, error)
	Save(Cliplist *model.Cliplist) error
	Delete() (int, error)
}

func NewCliplistRepository() CliplistRepository {
	return &CliplistRepositoryImpl{}
}

type CliplistRepositoryImpl struct{}

func (r *CliplistRepositoryImpl) NewQuery(tx *gorm.DB) CliplistQuery {
	return &CliplistQueryImpl{Tx: tx}
}

type CliplistQueryImpl struct {
	Tx *gorm.DB
}

func (q *CliplistQueryImpl) Where(cond *model.Cliplist) CliplistQuery {
	return &CliplistQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *CliplistQueryImpl) JoinChannel() CliplistQuery {
	return &CliplistQueryImpl{Tx: q.Tx.Preload("Channel")}
}

func (q *CliplistQueryImpl) Create(Cliplist *model.Cliplist) error {
	err := q.Tx.Create(Cliplist).Error
	return newErr(err)
}

func (q *CliplistQueryImpl) Find() (*model.Cliplist, error) {
	res := &model.Cliplist{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *CliplistQueryImpl) FindAll() ([]*model.Cliplist, error) {
	res := make([]*model.Cliplist, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *CliplistQueryImpl) Save(Cliplist *model.Cliplist) error {
	err := q.Tx.Save(Cliplist).Error
	return newErr(err)
}

func (q *CliplistQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Cliplist{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
