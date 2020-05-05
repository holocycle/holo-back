package repository

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type VideoRepository struct {
	Tx *gorm.DB
}

func NewVideoRepository(ctx context.Context) *VideoRepository {
	return &VideoRepository{
		Tx: ctx.GetDB(),
	}
}

func (r *VideoRepository) FindAll(cond *model.Video) ([]*model.Video, error) {
	res := make([]*model.Video, 0)
	if err := r.Tx.Where(cond).Find(&res).Error; err != nil {
		return nil, err // FIXME
	}

	return res, nil
}

func (r *VideoRepository) FindBy(cond *model.Video) (*model.Video, error) {
	res := &model.Video{}
	if err := r.Tx.Where(cond).First(res).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err // FIXME
		}
		return nil, err // FIXME
	}
	return res, nil
}

func (r *VideoRepository) Save(video *model.Video) error {
	if err := r.Tx.Save(video).Error; err != nil {
		return err
	}
	return nil
}
