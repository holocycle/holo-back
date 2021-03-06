package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type CommentRepository interface {
	NewQuery(ctx context.Context) CommentQuery
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

func (r *CommentRepositoryImpl) NewQuery(ctx context.Context) CommentQuery {
	return &CommentQueryImpl{Tx: app_context.GetDB(ctx)}
}

type CommentQueryImpl struct {
	Tx *gorm.DB
}

func (q *CommentQueryImpl) Where(cond *model.Comment) CommentQuery {
	return &CommentQueryImpl{Tx: q.Tx.Where(cond)}
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
	err := q.Tx.Create(Comment).Error
	return newErr(err)
}

func (q *CommentQueryImpl) Find() (*model.Comment, error) {
	res := &model.Comment{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *CommentQueryImpl) FindAll() ([]*model.Comment, error) {
	res := make([]*model.Comment, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *CommentQueryImpl) Save(Comment *model.Comment) error {
	err := q.Tx.Save(Comment).Error
	return newErr(err)
}

func (q *CommentQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Comment{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
