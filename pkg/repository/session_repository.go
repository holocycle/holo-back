package repository

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type SessionRepository interface {
	NewQuery(ctx context.Context) SessionQuery
}

type SessionQuery interface {
	Where(cond *model.Session) SessionQuery

	Create(session *model.Session) error
	Find() (*model.Session, error)
	FindAll() ([]*model.Session, error)
	Save(session *model.Session) error
	Delete() (int, error)
}

func NewSessionRepository() SessionRepository {
	return &SessionRepositoryImpl{}
}

type SessionRepositoryImpl struct{}

func (r *SessionRepositoryImpl) NewQuery(ctx context.Context) SessionQuery {
	return &SessionQueryImpl{Tx: app_context.GetDB(ctx)}
}

type SessionQueryImpl struct {
	Tx *gorm.DB
}

func (q *SessionQueryImpl) Where(cond *model.Session) SessionQuery {
	return &SessionQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *SessionQueryImpl) Create(session *model.Session) error {
	err := q.Tx.Create(session).Error
	return newErr(err)
}

func (q *SessionQueryImpl) Find() (*model.Session, error) {
	res := &model.Session{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *SessionQueryImpl) FindAll() ([]*model.Session, error) {
	res := make([]*model.Session, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *SessionQueryImpl) Save(session *model.Session) error {
	err := q.Tx.Save(session).Error
	return newErr(err)
}

func (q *SessionQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Session{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
