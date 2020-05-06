package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/holocycle/holo-back/pkg/model"
)

type TagRepository interface {
	FindAll(tx *gorm.DB, cond *TagCondition) ([]*model.Tag, error)
	FindBy(tx *gorm.DB, cond *TagCondition) (*model.Tag, error)
	Save(tx *gorm.DB, tag *model.Tag) error
}

type TagCondition struct {
	ID   string
	Name string
}

func NewTagRepository() TagRepository {
	return &TagRepositoryImpl{}
}

type TagRepositoryImpl struct {
}

func (r *TagRepositoryImpl) FindAll(tx *gorm.DB, cond *TagCondition) ([]*model.Tag, error) {
	res := make([]*model.Tag, 0)

	tx = tx.Where(&model.Tag{
		ID:   cond.ID,
		Name: cond.Name,
	})
	if err := tx.Where(cond).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *TagRepositoryImpl) FindBy(tx *gorm.DB, cond *TagCondition) (*model.Tag, error) {
	res := &model.Tag{}

	tx = tx.Where(&model.Tag{
		ID:   cond.ID,
		Name: cond.Name,
	})
	if err := tx.First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *TagRepositoryImpl) Save(tx *gorm.DB, tag *model.Tag) error {
	if err := tx.Save(tag).Error; err != nil {
		return err
	}
	return nil
}
