package repository

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type SessionRepository struct {
	Tx *gorm.DB
}

func NewSessionRepository(ctx context.Context) *SessionRepository {
	return &SessionRepository{
		Tx: ctx.GetDB(),
	}
}

func (r *SessionRepository) FindBy(cond *model.Session) (*model.Session, error) {
	res := &model.Session{}
	if err := r.Tx.Where(cond).First(res).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err // FIXME
		}
		return nil, err // FIXME
	}
	return res, nil
}

func (r *SessionRepository) Create(session *model.Session) error {
	if err := r.Tx.Create(session).Error; err != nil {
		return err
	}
	return nil
}

func (r *SessionRepository) Delete(session *model.Session) error {
	if err := r.Tx.Delete(session).Error; err != nil {
		return err
	}
	return nil
}
