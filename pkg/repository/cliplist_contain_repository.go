package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type CliplistContainRepository interface {
	NewQuery(ctx context.Context) CliplistContainQuery

	InsertToList(ctx context.Context, cliplistContain *model.CliplistContain) error
	DeleteFromList(ctx context.Context, cliplistContain *model.CliplistContain) error
}

type CliplistContainQuery interface {
	Where(cond *model.CliplistContain) CliplistContainQuery
	JoinClip() CliplistContainQuery

	Create(CliplistContain *model.CliplistContain) error
	Find() (*model.CliplistContain, error)
	FindAll() ([]*model.CliplistContain, error)
	Save(CliplistContain *model.CliplistContain) error
	Delete() (int, error)
}

func NewCliplistContainRepository() CliplistContainRepository {
	return &CliplistContainRepositoryImpl{}
}

type CliplistContainRepositoryImpl struct{}

func (r *CliplistContainRepositoryImpl) NewQuery(ctx context.Context) CliplistContainQuery {
	return &CliplistContainQueryImpl{Tx: app_context.GetDB(ctx)}
}

type CliplistContainQueryImpl struct {
	Tx *gorm.DB
}

func (q *CliplistContainQueryImpl) Where(cond *model.CliplistContain) CliplistContainQuery {
	return &CliplistContainQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *CliplistContainQueryImpl) JoinClip() CliplistContainQuery {
	return &CliplistContainQueryImpl{
		Tx: q.Tx.
			Preload("Clip").
			Preload("Clip.Video").
			Preload("Clip.Favorites"),
	}
}

func (q *CliplistContainQueryImpl) Create(CliplistContain *model.CliplistContain) error {
	err := q.Tx.Create(CliplistContain).Error
	return newErr(err)
}

func (q *CliplistContainQueryImpl) Find() (*model.CliplistContain, error) {
	res := &model.CliplistContain{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *CliplistContainQueryImpl) FindAll() ([]*model.CliplistContain, error) {
	res := make([]*model.CliplistContain, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *CliplistContainQueryImpl) Save(CliplistContain *model.CliplistContain) error {
	err := q.Tx.Save(CliplistContain).Error
	return newErr(err)
}

func (q *CliplistContainQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.CliplistContain{})
	return (int)(res.RowsAffected), newErr(res.Error)
}

func (r *CliplistContainRepositoryImpl) InsertToList(ctx context.Context, cliplistContain *model.CliplistContain) error {
	tx := app_context.GetDB(ctx)
	err := tx.Model(&model.CliplistContain{}).
		Where(&model.CliplistContain{CliplistID: cliplistContain.CliplistID}).
		Where("index >= ?", cliplistContain.Index).
		Update("index", gorm.Expr("index + 1")).
		Error
	if err != nil {
		return newErr(err)
	}

	err = tx.Create(cliplistContain).Error
	if err != nil {
		return newErr(err)
	}
	return nil
}

func (r *CliplistContainRepositoryImpl) DeleteFromList(ctx context.Context, cliplistContain *model.CliplistContain) error {
	tx := app_context.GetDB(ctx)
	err := tx.Delete(cliplistContain).Error
	if err != nil {
		return newErr(err)
	}

	err = tx.Model(&model.CliplistContain{}).
		Where(&model.CliplistContain{CliplistID: cliplistContain.CliplistID}).
		Where("index >= ?", cliplistContain.Index).
		Update("index", gorm.Expr("index - 1")).
		Error
	if err != nil {
		return newErr(err)
	}
	return nil
}
