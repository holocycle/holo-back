package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/holocycle/holo-back/pkg/model"
)

type TagRepository interface {
	NewQuery(tx *gorm.DB) TagQuery
}

type TagQuery interface {
	Where(cond *model.Tag) TagQuery

	Create(tag *model.Tag) error
	Find() (*model.Tag, error)
	FindAll() ([]*model.Tag, error)
	Save(tag *model.Tag) error
	Delete() (int, error)
}

func NewTagRepository() TagRepository {
	return &TagRepositoryImpl{}
}

type TagRepositoryImpl struct {
}

func (r *TagRepositoryImpl) NewQuery(tx *gorm.DB) TagQuery {
	return &TagQueryImpl{Tx: tx}
}

type TagQueryImpl struct {
	Tx *gorm.DB
}

func (q *TagQueryImpl) Where(cond *model.Tag) TagQuery {
	return &TagQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *TagQueryImpl) Create(tag *model.Tag) error {
	err := q.Tx.Create(tag).Error
	return newErr(err)
}

func (q *TagQueryImpl) Find() (*model.Tag, error) {
	res := &model.Tag{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *TagQueryImpl) FindAll() ([]*model.Tag, error) {
	res := make([]*model.Tag, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *TagQueryImpl) Save(tag *model.Tag) error {
	err := q.Tx.Save(tag).Error
	return newErr(err)
}

func (q *TagQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Tag{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
