package repository

import (
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type CommentRepository interface {
	NewQuery(tx *gorm.DB) CommentQuery
}

type CommentQuery interface {
	Where(cond *model.Comment) CommentQuery

	JoinUser() CommentQuery

	Limit(limit int) CommentQuery
	Latest() CommentQuery

	Create(Comment *model.Comment) error
	Find() (*model.Comment, error)
	FindAll() ([]*model.Comment, error)
	Save(Comment *model.Comment) error
	Delete() (int, error)
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

type CommentRepositoryImpl struct{}

func (r *CommentRepositoryImpl) NewQuery(tx *gorm.DB) CommentQuery {
	return &CommentQueryImpl{Tx: tx}
}

type CommentQueryImpl struct {
	Tx *gorm.DB
}

func (q *CommentQueryImpl) Where(cond *model.Comment) CommentQuery {
	return &CommentQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *CommentQueryImpl) JoinClip() CommentQuery {
	return &CommentQueryImpl{Tx: q.Tx.Preload("Clip")}
}

func (q *CommentQueryImpl) JoinUser() CommentQuery {
	return &CommentQueryImpl{Tx: q.Tx.Preload("User")}
}

func (q *CommentQueryImpl) Limit(limit int) CommentQuery {
	return &CommentQueryImpl{Tx: q.Tx.Limit(limit)}
}

func (q *CommentQueryImpl) Latest() CommentQuery {
	return &CommentQueryImpl{Tx: q.Tx.Order("created_at desc")}
}

func (q *CommentQueryImpl) Create(Comment *model.Comment) error {
	return q.Tx.Create(Comment).Error
}

func (q *CommentQueryImpl) Find() (*model.Comment, error) {
	res := &model.Comment{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *CommentQueryImpl) FindAll() ([]*model.Comment, error) {
	res := make([]*model.Comment, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (q *CommentQueryImpl) Save(Comment *model.Comment) error {
	return q.Tx.Save(Comment).Error
}

func (q *CommentQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Comment{})
	return (int)(res.RowsAffected), res.Error
}
