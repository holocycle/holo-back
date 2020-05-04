package repository

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type ClipRepository struct {
	Tx *gorm.DB
}

func NewClipRepository(ctx context.Context) *ClipRepository {
	return &ClipRepository{
		Tx: ctx.GetDB(),
	}
}

func (r *ClipRepository) FindBy(cond *model.Clip) (*model.Clip, error) {
	res := &model.Clip{}
	if err := r.Tx.Where(cond).First(res).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err // FIXME
		}
		return nil, err // FIXME
	}
	return res, nil
}

func (r *ClipRepository) Create(clip *model.Clip) error {
	if err := r.Tx.Create(clip).Error; err != nil {
		return err
	}
	return nil
}
