package service

import (
	"context"

	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/core/service"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
)

const (
	FAVORITE   = true
	UNFAVORITE = false
)

type FavoriteService interface {
	GetFavoriteItem(
		ctx context.Context,
		cliID string,
		userID string,
	) (*api.GetFavoriteResponse, service.Error)

	PutFavoriteItem(
		ctx context.Context,
		clipID string,
		userID string,
	) (*api.PutFavoriteResponse, service.Error)

	DeleteFavoriteItem(
		ctx context.Context,
		clipID string,
		userID string,
	) (*api.DeleteFavoriteResponse, service.Error)
}

type FavoriteServiceImpl struct {
	RepositoryContainer *repository.Container
}

func NewFavoriteService() FavoriteService {
	return &FavoriteServiceImpl{
		RepositoryContainer: repository.NewContainer(),
	}
}

func (s *FavoriteServiceImpl) GetFavoriteItem(
	ctx context.Context,
	clipID string,
	userID string,
) (*api.GetFavoriteResponse, service.Error) {
	favorite := model.NewFavorite(clipID, userID)
	count := s.RepositoryContainer.FavoriteRepository.
		NewQuery(ctx).
		Where(favorite).
		Count()
	if count == 0 {
		return &api.GetFavoriteResponse{
			Favorite: UNFAVORITE,
		}, nil
	}

	if count == 1 {
		return &api.GetFavoriteResponse{
			Favorite: FAVORITE,
		}, nil
	}

	return nil, InternalError
}

func (s *FavoriteServiceImpl) PutFavoriteItem(
	ctx context.Context,
	clipID string,
	userID string,
) (*api.PutFavoriteResponse, service.Error) {
	if _, err := s.RepositoryContainer.ClipRepository.NewQuery(ctx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find(); err != nil {
		if repository.NotFoundError(err) {
			return nil, ClipNotFound.With(err)
		}
	}

	favorite := model.NewFavorite(clipID, userID)
	count := s.RepositoryContainer.FavoriteRepository.NewQuery(ctx).Where(favorite).Count()
	if count < 0 {
		return nil, InternalError
	} else if count >= 1 {
		// お気に入りにしようとしたが、既にお気に入り登録済の状態。正常終了と扱う
		return &api.PutFavoriteResponse{}, nil
	}

	if err := s.RepositoryContainer.FavoriteRepository.NewQuery(ctx).Create(favorite); err != nil {
		return nil, InternalError.With(err)
	}

	return &api.PutFavoriteResponse{}, nil
}

func (s *FavoriteServiceImpl) DeleteFavoriteItem(
	ctx context.Context,
	clipID string,
	userID string,
) (*api.DeleteFavoriteResponse, service.Error) {
	if _, err := s.RepositoryContainer.ClipRepository.NewQuery(ctx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find(); err != nil {
		if repository.NotFoundError(err) {
			return nil, ClipNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}

	favorite := model.NewFavorite(clipID, userID)
	_, err := s.RepositoryContainer.FavoriteRepository.NewQuery(ctx).Where(favorite).Delete()
	if err != nil {
		return nil, InternalError.With(err)
	}

	return &api.DeleteFavoriteResponse{}, nil
}
