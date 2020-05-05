package repository

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type ClipRepository struct {
	Tx *gorm.DB
}

type ClipOrder int

const (
	Any ClipOrder = iota
	RecentlyCreated
)

type ClipCondition struct {
	model.Clip
	Limit   int
	OrderBy ClipOrder
}

func NewClipRepository(ctx context.Context) *ClipRepository {
	return &ClipRepository{
		Tx: ctx.GetDB(),
	}
}

func (r *ClipRepository) FindAll(cond *ClipCondition) ([]*model.Clip, error) {
	tx := r.Tx
	if cond.Limit > 0 {
		tx = tx.Limit(cond.Limit)
	}
	if cond.OrderBy == RecentlyCreated {
		tx = tx.Order("created_at desc")
	}

	res := make([]*model.Clip, 0)
	if err := tx.Where(cond.Clip).Find(&res).Error; err != nil {
		return nil, err // FIXME
	}

	return res, nil
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
