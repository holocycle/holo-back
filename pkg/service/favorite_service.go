package service

import (
	"context"

	"github.com/holocycle/holo-back/pkg/api"
	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/core/service"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
)

const (
	FAVORITE   = true
	UNFAVORITE = false
)

type FavoriteService interface {
	GetFavorite(
		ctx context.Context,
		clipID string,
		request *api.GetFavoriteRequest,
	) (*api.GetFavoriteResponse, service.Error)

	PutFavorite(
		ctx context.Context,
		clipID string,
		request *api.PutFavoriteRequest,
	) (*api.PutFavoriteResponse, service.Error)

	DeleteFavorite(
		ctx context.Context,
		clipID string,
		request *api.DeleteFavoriteRequest,
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

func (s *FavoriteServiceImpl) GetFavorite(
	ctx context.Context,
	clipID string,
	_ *api.GetFavoriteRequest,
) (*api.GetFavoriteResponse, service.Error) {
	_, err := s.RepositoryContainer.ClipRepository.
		NewQuery(ctx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, ClipNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}

	favorite := model.NewFavorite(clipID, app_context.GetSession(ctx).UserID)
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

func (s *FavoriteServiceImpl) PutFavorite(
	ctx context.Context,
	clipID string,
	_ *api.PutFavoriteRequest,
) (*api.PutFavoriteResponse, service.Error) {
	_, err := s.RepositoryContainer.ClipRepository.
		NewQuery(ctx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, ClipNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}

	favorite := model.NewFavorite(clipID, app_context.GetSession(ctx).UserID)
	count := s.RepositoryContainer.FavoriteRepository.
		NewQuery(ctx).
		Where(favorite).
		Count()
	if count < 0 {
		return nil, InternalError
	}

	err = s.RepositoryContainer.FavoriteRepository.
		NewQuery(ctx).
		Save(favorite)
	if err != nil {
		return nil, InternalError.With(err)
	}

	return &api.PutFavoriteResponse{}, nil
}

func (s *FavoriteServiceImpl) DeleteFavorite(
	ctx context.Context,
	clipID string,
	_ *api.DeleteFavoriteRequest,
) (*api.DeleteFavoriteResponse, service.Error) {
	_, err := s.RepositoryContainer.ClipRepository.
		NewQuery(ctx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, ClipNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}

	favorite := model.NewFavorite(clipID, app_context.GetSession(ctx).UserID)
	_, err = s.RepositoryContainer.FavoriteRepository.
		NewQuery(ctx).
		Where(favorite).
		Delete()
	if err != nil {
		return nil, InternalError.With(err)
	}

	return &api.DeleteFavoriteResponse{}, nil
}
