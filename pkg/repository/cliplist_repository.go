package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type CliplistRepository interface {
	NewQuery(ctx context.Context) CliplistQuery
}

type CliplistQuery interface {
	Where(cond *model.Cliplist) CliplistQuery
	JoinClip() CliplistQuery

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

func (r *CliplistRepositoryImpl) NewQuery(ctx context.Context) CliplistQuery {
	return &CliplistQueryImpl{Tx: app_context.GetDB(ctx)}
}

type CliplistQueryImpl struct {
	Tx *gorm.DB
}

func (q *CliplistQueryImpl) Where(cond *model.Cliplist) CliplistQuery {
	return &CliplistQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *CliplistQueryImpl) JoinClip() CliplistQuery {
	return &CliplistQueryImpl{
		Tx: q.Tx.
			Preload("CliplistContains", func(tx *gorm.DB) *gorm.DB {
				return tx.Order("cliplist_contains.index")
			}).
			Preload("CliplistContains.Clip").
			Preload("CliplistContains.Clip.Video").
			Preload("CliplistContains.Clip.Favorites"),
	}
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
