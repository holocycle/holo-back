package repository

import (
	"context"

	"github.com/jinzhu/gorm"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
)

type TagRepository interface {
	NewQuery(ctx context.Context) TagQuery
}

type TagQuery interface {
	Where(cond *model.Tag) TagQuery
	Like(name string) TagQuery

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

func (r *TagRepositoryImpl) NewQuery(ctx context.Context) TagQuery {
	return &TagQueryImpl{Tx: app_context.GetDB(ctx)}
}

type TagQueryImpl struct {
	Tx *gorm.DB
}

func (q *TagQueryImpl) Where(cond *model.Tag) TagQuery {
	return &TagQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *TagQueryImpl) Like(name string) TagQuery {
	tx := q.Tx.Table("tags").
		Where("name LIKE ?", "%"+name+"%")
	return &TagQueryImpl{Tx: tx}
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
