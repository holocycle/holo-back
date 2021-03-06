package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type ChannelRepository interface {
	NewQuery(ctx context.Context) ChannelQuery
}

type ChannelQuery interface {
	Where(cond *model.Channel) ChannelQuery

	Create(channel *model.Channel) error
	Find() (*model.Channel, error)
	FindAll() ([]*model.Channel, error)
	Save(channel *model.Channel) error
	Delete() (int, error)
}

func NewChannelRepository() ChannelRepository {
	return &ChannelRepositoryImpl{}
}

type ChannelRepositoryImpl struct{}

func (r *ChannelRepositoryImpl) NewQuery(ctx context.Context) ChannelQuery {
	return &ChannelQueryImpl{Tx: app_context.GetDB(ctx)}
}

type ChannelQueryImpl struct {
	Tx *gorm.DB
}

func (q *ChannelQueryImpl) Where(cond *model.Channel) ChannelQuery {
	return &ChannelQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *ChannelQueryImpl) Create(channel *model.Channel) error {
	err := q.Tx.Create(channel).Error
	return newErr(err)
}

func (q *ChannelQueryImpl) Find() (*model.Channel, error) {
	res := &model.Channel{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ChannelQueryImpl) FindAll() ([]*model.Channel, error) {
	res := make([]*model.Channel, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ChannelQueryImpl) Save(channel *model.Channel) error {
	err := q.Tx.Save(channel).Error
	return newErr(err)
}

func (q *ChannelQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Channel{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
