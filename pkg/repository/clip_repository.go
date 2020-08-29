package repository

import (
	"context"
	"strings"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
)

type ClipRepository interface {
	NewQuery(ctx context.Context) ClipQuery
}

type ClipQuery interface {
	Where(cond *model.Clip) ClipQuery

	Limit(limit int) ClipQuery
	Latest() ClipQuery
	TopRated() ClipQuery
	JoinVideo() ClipQuery
	JoinFavorite() ClipQuery
	WhereContainsTags(tagIDs []string) ClipQuery

	Create(clip *model.Clip) error
	Find() (*model.Clip, error)
	FindAll() ([]*model.Clip, error)
	Save(clip *model.Clip) error
	Delete() (int, error)
}

func NewClipRepository() ClipRepository {
	return &ClipRepositoryImpl{}
}

type ClipRepositoryImpl struct{}

func (r *ClipRepositoryImpl) NewQuery(ctx context.Context) ClipQuery {
	return &ClipQueryImpl{Tx: app_context.GetDB(ctx)}
}

type ClipQueryImpl struct {
	Tx *gorm.DB
}

func (q *ClipQueryImpl) Where(cond *model.Clip) ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Where(cond)}
}

func (q *ClipQueryImpl) Limit(limit int) ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Limit(limit)}
}

func (q *ClipQueryImpl) Latest() ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Order("created_at desc").
		Group("clips.id")}
}

func (q *ClipQueryImpl) TopRated() ClipQuery {
	tx := q.Tx.Table("clips").
		Select(strings.Join([]string{
			"clips.*",
			"COUNT(distinct favorites.user_id) as favorite_count",
		}, ",")).
		Joins("LEFT JOIN favorites ON clips.id = favorites.clip_id").
		Group("clips.id").
		Order("favorite_count desc")
	return &ClipQueryImpl{Tx: tx}
}

func (q *ClipQueryImpl) JoinVideo() ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Preload("Video")}
}

func (q *ClipQueryImpl) JoinFavorite() ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.Preload("Favorites")}
}

func (q *ClipQueryImpl) WhereContainsTags(tagNames []string) ClipQuery {
	return &ClipQueryImpl{Tx: q.Tx.
		Where("clips.id IN (?)",
			q.Tx.Table("clip_tagged").
				Select("clip_tagged.clip_id").
				Joins("INNER JOIN tags ON clip_tagged.tag_id = tags.id").
				Where("tags.name IN (?)", tagNames).
				Group("clip_tagged.clip_id").
				Having("COUNT(distinct clip_tagged.tag_id) = (?)", len(tagNames)).
				SubQuery(),
		)}
}

func (q *ClipQueryImpl) Create(clip *model.Clip) error {
	err := q.Tx.Create(clip).Error
	return newErr(err)
}

func (q *ClipQueryImpl) Find() (*model.Clip, error) {
	res := &model.Clip{}
	if err := q.Tx.First(res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ClipQueryImpl) FindAll() ([]*model.Clip, error) {
	res := make([]*model.Clip, 0)
	if err := q.Tx.Find(&res).Error; err != nil {
		return nil, newErr(err)
	}
	return res, nil
}

func (q *ClipQueryImpl) Save(clip *model.Clip) error {
	err := q.Tx.Save(clip).Error
	return newErr(err)
}

func (q *ClipQueryImpl) Delete() (int, error) {
	res := q.Tx.Delete(&model.Clip{})
	return (int)(res.RowsAffected), newErr(res.Error)
}
