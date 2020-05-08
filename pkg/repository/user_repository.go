package repository

import (
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	NewQuery(tx *gorm.DB) UserQuery
}

type UserQuery interface {
	Where(cond *model.User) UserQuery

	Create(user *model.User) error
	Find() (*model.User, error)
	FindAll() ([]*model.User, error)
	Save(user *model.User) error
	Delete() (int, error)
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) NewQuery(tx *gorm.DB) UserQuery {
	return &UserQueryImpl{Tx: tx}
}

type UserQueryImpl struct {
	Tx *gorm.DB
}

func (q *UserQueryImpl) Where(cond *model.User) UserQuery {
	return &UserQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *UserQueryImpl) Create(user *model.User) error {
	return q.Tx.Create(user).Error
}

func (q *UserQueryImpl) Find() (*model.User, error) {
	res := &model.User{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *UserQueryImpl) FindAll() ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *UserQueryImpl) Save(user *model.User) error {
	return q.Tx.Save(user).Error
}

func (q *UserQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.User{})
	return (int)(res.RowsAffected), res.Error
}
