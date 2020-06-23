package service

import (
	"context"

	"github.com/holocycle/holo-back/pkg/api"
	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/core/service"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
)

type CliplistItemService interface {
	GetCliplistItem(
		ctx context.Context,
		cliplistID string,
		index int,
		req *api.GetCliplistItemRequest,
	) (*api.GetCliplistItemResponse, service.Error)

	PostCliplistItem(
		ctx context.Context,
		cliplistID string,
		index int,
		req *api.PostCliplistItemRequest,
	) (*api.PostCliplistItemResponse, service.Error)

	DeleteCliplistItem(
		ctx context.Context,
		cliplistID string,
		index int,
		req *api.DeleteCliplistItemRequest,
	) (*api.DeleteCliplistItemResponse, service.Error)
}

type CliplistItemServiceImpl struct {
	ClipRepository            repository.ClipRepository
	CliplistRepository        repository.CliplistRepository
	CliplistContainRepository repository.CliplistContainRepository
}

func NewCliplistItemService() CliplistItemService {
	return &CliplistItemServiceImpl{
		ClipRepository:            repository.NewClipRepository(),
		CliplistRepository:        repository.NewCliplistRepository(),
		CliplistContainRepository: repository.NewCliplistContainRepository(),
	}
}

func (s *CliplistItemServiceImpl) GetCliplistItem(
	ctx context.Context,
	cliplistID string,
	index int,
	req *api.GetCliplistItemRequest,
) (*api.GetCliplistItemResponse, service.Error) {
	_, err := s.CliplistRepository.NewQuery(ctx).
		Where(&model.Cliplist{
			ID:     cliplistID,
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, CliplistNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}

	cliplistContain, err := s.CliplistContainRepository.NewQuery(ctx).
		JoinClip().
		Where(&model.CliplistContain{
			CliplistID: cliplistID,
			Index:      index,
		}).Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, CliplistItemNotFound
		}
		return nil, InternalError.With(err)
	}

	return &api.GetCliplistItemResponse{
		CliplistItem: converter.ConvertToCliplistItem(cliplistContain),
	}, nil
}

func (s *CliplistItemServiceImpl) PostCliplistItem(
	ctx context.Context,
	cliplistID string,
	index int,
	req *api.PostCliplistItemRequest,
) (*api.PostCliplistItemResponse, service.Error) {
	cliplist, err := s.CliplistRepository.NewQuery(ctx).
		Where(&model.Cliplist{
			ID:     cliplistID,
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, CliplistItemNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}
	if cliplist.UserID != app_context.GetSession(ctx).UserID {
		return nil, NoPermissionToCliplist
	}

	cliplistContains, err := s.CliplistContainRepository.NewQuery(ctx).
		Where(&model.CliplistContain{
			CliplistID: cliplistID,
		}).FindAll()
	if err != nil {
		return nil, InternalError.With(err)
	}
	if index > len(cliplistContains) {
		return nil, CliplistIndexOutOfRange
	}

	_, err = s.ClipRepository.NewQuery(ctx).Where(&model.Clip{
		ID:     req.ClipID,
		Status: model.ClipStatusPublic,
	}).Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, ClipNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}

	cliplistContain := model.NewCliplistContain(
		cliplistID,
		index,
		req.ClipID,
	)
	err = s.CliplistContainRepository.InsertToList(ctx, cliplistContain)
	if err != nil {
		return nil, InternalError.With(err)
	}

	return &api.PostCliplistItemResponse{
		CliplistID: cliplistID,
	}, nil
}

func (s *CliplistItemServiceImpl) DeleteCliplistItem(
	ctx context.Context,
	cliplistID string,
	index int,
	req *api.DeleteCliplistItemRequest,
) (*api.DeleteCliplistItemResponse, service.Error) {
	cliplist, err := s.CliplistRepository.NewQuery(ctx).
		Where(&model.Cliplist{
			ID:     cliplistID,
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return nil, CliplistNotFound.With(err)
		}
		return nil, InternalError.With(err)
	}
	if cliplist.UserID != app_context.GetSession(ctx).UserID {
		return nil, NoPermissionToCliplist.With(err)
	}

	cliplistContains, err := s.CliplistContainRepository.NewQuery(ctx).
		Where(&model.CliplistContain{
			CliplistID: cliplistID,
		}).FindAll()
	if err != nil {
		return nil, InternalError.With(err)
	}
	if index >= len(cliplistContains) {
		return nil, CliplistIndexOutOfRange
	}

	err = s.CliplistContainRepository.DeleteFromList(ctx, cliplistContains[index])
	if err != nil {
		return nil, InternalError.With(err)
	}

	return &api.DeleteCliplistItemResponse{}, nil
}
