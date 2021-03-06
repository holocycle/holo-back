package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type VideoRepository interface {
	NewQuery(ctx context.Context) VideoQuery
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

func (r *VideoRepositoryImpl) NewQuery(ctx context.Context) VideoQuery {
	return &VideoQueryImpl{Tx: app_context.GetDB(ctx)}
}

type VideoQueryImpl struct {
	Tx *gorm.DB
}

func (q *VideoQueryImpl) Where(cond *model.Video) VideoQuery {
	return &VideoQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *VideoQueryImpl) Create(video *model.Video) error {
	err := q.Tx.Create(video).Error
	return newErr(err)
}

func (q *VideoQueryImpl) Find() (*model.Video, error) {
	res := &model.Video{}
	if err := q.Tx.First(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *VideoQueryImpl) FindAll() ([]*model.Video, error) {
	res := make([]*model.Video, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *VideoQueryImpl) Save(video *model.Video) error {
	err := q.Tx.Save(video).Error
	return newErr(err)
}

func (q *VideoQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Video{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
