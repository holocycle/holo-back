package repository

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type LiverRepository struct {
	Tx *gorm.DB
}

func NewLiverRepository(ctx context.Context) *LiverRepository {
	return &LiverRepository{
		Tx: ctx.GetDB(),
	}
}

func (r *LiverRepository) FindAll(cond *model.Liver) ([]*model.Liver, error) {
	res := make([]*model.Liver, 0)
	r.Tx.LogMode(true)
	if err := r.Tx.Where(cond).Find(&res).Error; err != nil {
		return nil, err // FIXME
	}

	return res, nil
}

func (r *LiverRepository) FindBy(cond *model.Liver) (*model.Liver, error) {
	res := &model.Liver{}
	if err := r.Tx.Where(cond).First(res).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err // FIXME
		}
		return nil, err // FIXME
	}
	return res, nil
}
