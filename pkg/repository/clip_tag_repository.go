package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/holocycle/holo-back/pkg/model"
)

type ClipTagRepository interface {
	NewQuery(tx *gorm.DB) ClipTagQuery
}

type ClipTagQuery interface {
	Where(cond *model.ClipTag) ClipTagQuery

	JoinUser() ClipTagQuery
	JoinClip() ClipTagQuery
	JoinTag() ClipTagQuery

	Create(clipTag *model.ClipTag) error
	Find() (*model.ClipTag, error)
	FindAll() ([]*model.ClipTag, error)
	Save(clipTag *model.ClipTag) error
	Delete() (int, error)
}

func NewClipTagRepository() ClipTagRepository {
	return &ClipTagRepositoryImpl{}
}

type ClipTagRepositoryImpl struct {
}

func (r *ClipTagRepositoryImpl) NewQuery(tx *gorm.DB) ClipTagQuery {
	return &ClipTagQueryImpl{Tx: tx}
}

type ClipTagQueryImpl struct {
	Tx *gorm.DB
}

func (q *ClipTagQueryImpl) Where(cond *model.ClipTag) ClipTagQuery {
	return &ClipTagQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *ClipTagQueryImpl) JoinUser() ClipTagQuery {
	return &ClipTagQueryImpl{Tx: q.Tx.Preload("User")}
}

func (q *ClipTagQueryImpl) JoinClip() ClipTagQuery {
	return &ClipTagQueryImpl{Tx: q.Tx.Preload("Clip")}
}

func (q *ClipTagQueryImpl) JoinTag() ClipTagQuery {
	return &ClipTagQueryImpl{Tx: q.Tx.Preload("Tag")}
}

func (q *ClipTagQueryImpl) Create(clipTag *model.ClipTag) error {
	err := q.Tx.Create(clipTag).Error
	return newErr(err)
}

func (q *ClipTagQueryImpl) Find() (*model.ClipTag, error) {
	res := &model.ClipTag{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ClipTagQueryImpl) FindAll() ([]*model.ClipTag, error) {
	res := make([]*model.ClipTag, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ClipTagQueryImpl) Save(clipTag *model.ClipTag) error {
	err := q.Tx.Save(clipTag).Error
	return newErr(err)
}

func (q *ClipTagQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.ClipTag{})
	return int(res.RowsAffected), newErr(res.Error)
}
