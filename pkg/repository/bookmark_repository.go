package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type BookmarkRepository interface {
	NewQuery(ctx context.Context) BookmarkQuery
}

type BookmarkQuery interface {
	Where(cond *model.Bookmark) BookmarkQuery

	Create(Bookmark *model.Bookmark) error
	Find() (*model.Bookmark, error)
	FindAll() ([]*model.Bookmark, error)
	Save(Bookmark *model.Bookmark) error
	Delete() (int, error)
}

func NewBookmarkRepository() BookmarkRepository {
	return &BookmarkRepositoryImpl{}
}

type BookmarkRepositoryImpl struct{}

func (r *BookmarkRepositoryImpl) NewQuery(ctx context.Context) BookmarkQuery {
	return &BookmarkQueryImpl{Tx: app_context.GetDB(ctx)}
}

type BookmarkQueryImpl struct {
	Tx *gorm.DB
}

func (q *BookmarkQueryImpl) Where(cond *model.Bookmark) BookmarkQuery {
	return &BookmarkQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *BookmarkQueryImpl) Create(Bookmark *model.Bookmark) error {
	err := q.Tx.Create(Bookmark).Error
	return newErr(err)
}

func (q *BookmarkQueryImpl) Find() (*model.Bookmark, error) {
	res := &model.Bookmark{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *BookmarkQueryImpl) FindAll() ([]*model.Bookmark, error) {
	res := make([]*model.Bookmark, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *BookmarkQueryImpl) Save(Bookmark *model.Bookmark) error {
	err := q.Tx.Save(Bookmark).Error
	return newErr(err)
}

func (q *BookmarkQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Bookmark{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
