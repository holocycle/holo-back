package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type ClipTaggedRepository interface {
	NewQuery(ctx context.Context) ClipTaggedQuery
}

type ClipTaggedQuery interface {
	Where(cond *model.ClipTagged) ClipTaggedQuery
	JoinTag() ClipTaggedQuery

	Create(ClipTagged *model.ClipTagged) error
	Find() (*model.ClipTagged, error)
	FindAll() ([]*model.ClipTagged, error)
	Save(ClipTagged *model.ClipTagged) error
	Delete() (int, error)
}

func NewClipTaggedRepository() ClipTaggedRepository {
	return &ClipTaggedRepositoryImpl{}
}

type ClipTaggedRepositoryImpl struct{}

func (r *ClipTaggedRepositoryImpl) NewQuery(ctx context.Context) ClipTaggedQuery {
	return &ClipTaggedQueryImpl{Tx: app_context.GetDB(ctx)}
}

type ClipTaggedQueryImpl struct {
	Tx *gorm.DB
}

func (q *ClipTaggedQueryImpl) Where(cond *model.ClipTagged) ClipTaggedQuery {
	return &ClipTaggedQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *ClipTaggedQueryImpl) JoinTag() ClipTaggedQuery {
	return &ClipTaggedQueryImpl{Tx: q.Tx.Preload("Tag")}
}

func (q *ClipTaggedQueryImpl) Create(ClipTagged *model.ClipTagged) error {
	err := q.Tx.Create(ClipTagged).Error
	return newErr(err)
}

func (q *ClipTaggedQueryImpl) Find() (*model.ClipTagged, error) {
	res := &model.ClipTagged{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ClipTaggedQueryImpl) FindAll() ([]*model.ClipTagged, error) {
	res := make([]*model.ClipTagged, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ClipTaggedQueryImpl) Save(ClipTagged *model.ClipTagged) error {
	err := q.Tx.Save(ClipTagged).Error
	return newErr(err)
}

func (q *ClipTaggedQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.ClipTagged{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
