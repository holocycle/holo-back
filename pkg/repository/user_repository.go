package repository

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Tx *gorm.DB
}

func NewUserRepository(ctx context.Context) *UserRepository {
	return &UserRepository{
		Tx: ctx.GetDB(),
	}
}

func (r *UserRepository) FindBy(cond *model.User) (*model.User, error) {
	res := &model.User{}
	if err := r.Tx.Where(cond).First(res).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err // FIXME
		}
		return nil, err // FIXME
	}
	return res, nil
}

func (r *UserRepository) Create(user *model.User) error {
	if err := r.Tx.Create(user).Error; err != nil {
		return err
	}
	return nil
}
