package repository

import (
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type FavoriteRepository interface {
	NewQuery(tx *gorm.DB) FavoriteQuery
}

type FavoriteQuery interface {
	Where(cond *model.Favorite) FavoriteQuery

	Create(Favorite *model.Favorite) error
	Find() (*model.Favorite, error)
	FindAll() ([]*model.Favorite, error)
	Save(Favorite *model.Favorite) error
	Delete() (int, error)
}

func NewFavoriteRepository() FavoriteRepository {
	return &FavoriteRepositoryImpl{}
}

type FavoriteRepositoryImpl struct{}

func (r *FavoriteRepositoryImpl) NewQuery(tx *gorm.DB) FavoriteQuery {
	return &FavoriteQueryImpl{Tx: tx}
}

type FavoriteQueryImpl struct {
	Tx *gorm.DB
}

func (q *FavoriteQueryImpl) Where(cond *model.Favorite) FavoriteQuery {
	return &FavoriteQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *FavoriteQueryImpl) Create(Favorite *model.Favorite) error {
	err := q.Tx.Create(Favorite).Error
	return newErr(err)
}

func (q *FavoriteQueryImpl) Find() (*model.Favorite, error) {
	res := &model.Favorite{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *FavoriteQueryImpl) FindAll() ([]*model.Favorite, error) {
	res := make([]*model.Favorite, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *FavoriteQueryImpl) Save(Favorite *model.Favorite) error {
	err := q.Tx.Save(Favorite).Error
	return newErr(err)
}

func (q *FavoriteQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Favorite{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
