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

type CliplistService interface {
	ListCliplists(
		ctx context.Context,
		req *api.ListCliplistsRequest,
	) (*api.ListCliplistsResponse, service.Error)

	GetCliplist(
		ctx context.Context,
		cliplistID string,
		req *api.GetCliplistRequest,
	) (*api.GetCliplistResponse, service.Error)

	PostCliplist(
		ctx context.Context,
		req *api.PostCliplistRequest,
	) (*api.PostCliplistResponse, service.Error)

	PutCliplist(
		ctx context.Context,
		cliplistID string,
		req *api.PutCliplistRequest,
	) (*api.PutCliplistResponse, service.Error)

	DeleteCliplist(
		ctx context.Context,
		cliplistID string,
		req *api.DeleteCliplistRequest,
	) (*api.DeleteCliplistResponse, service.Error)
}

type CliplistServiceImpl struct {
	ClipRepository     repository.ClipRepository
	CliplistRepository repository.CliplistRepository
}

func NewCliplistService() CliplistService {
	return &CliplistServiceImpl{
		ClipRepository:     repository.NewClipRepository(),
		CliplistRepository: repository.NewCliplistRepository(),
	}
}

func (s *CliplistServiceImpl) ListCliplists(
	ctx context.Context,
	req *api.ListCliplistsRequest,
) (*api.ListCliplistsResponse, service.Error) {
	tx := app_context.GetDB(ctx)

	// TODO: use query
	if req.Limit > 0 {
		tx = tx.Limit(req.Limit)
	}

	// TODO: OrderBy

	cliplist, err := s.CliplistRepository.NewQuery(tx).
		JoinClip().
		Where(&model.Cliplist{Status: model.CliplistStatusPublic}).
		FindAll()
	if err != nil {
		return nil, InternalError.With(err)
	}

	return &api.ListCliplistsResponse{
		Cliplists: converter.ConvertToCliplists(cliplist),
	}, nil
}

func (s *CliplistServiceImpl) GetCliplist(
	ctx context.Context,
	cliplistID string,
	req *api.GetCliplistRequest,
) (*api.GetCliplistResponse, service.Error) {
	cliplist, err := s.CliplistRepository.NewQuery(app_context.GetDB(ctx)).
		JoinClip().
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

	pageBegin := req.ItemPerPage * req.Page
	pageEnd := pageBegin + req.ItemPerPage
	if pageEnd > len(cliplist.CliplistContains) {
		pageEnd = len(cliplist.CliplistContains)
	}

	return &api.GetCliplistResponse{
		Cliplist:      converter.ConvertToCliplist(cliplist),
		PageInfo:      converter.ConvertToPageInfo(len(cliplist.CliplistContains), req.Page, req.ItemPerPage),
		CliplistItems: converter.ConvertToCliplistItems(cliplist.CliplistContains[pageBegin:pageEnd]),
	}, nil
}

func (s *CliplistServiceImpl) PostCliplist(
	ctx context.Context,
	req *api.PostCliplistRequest,
) (*api.PostCliplistResponse, service.Error) {
	cliplist := model.NewCliplist(
		app_context.GetSession(ctx).UserID,
		req.Title,
		req.Description,
		model.CliplistStatusPublic,
	)
	err := s.CliplistRepository.NewQuery(app_context.GetDB(ctx)).Create(cliplist)
	if err != nil {
		return nil, InternalError.With(err)
	}

	return &api.PostCliplistResponse{
		CliplistID: cliplist.ID,
	}, nil
}

func (s *CliplistServiceImpl) PutCliplist(
	ctx context.Context,
	cliplistID string,
	req *api.PutCliplistRequest,
) (*api.PutCliplistResponse, service.Error) {
	cliplist, err := s.CliplistRepository.NewQuery(app_context.GetDB(ctx)).
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
		return nil, NoPermissionToCliplist
	}

	cliplist.Title = req.Title
	cliplist.Description = req.Description
	if err = s.CliplistRepository.NewQuery(app_context.GetDB(ctx)).Save(cliplist); err != nil {
		return nil, InternalError.With(err)
	}

	return &api.PutCliplistResponse{
		CliplistID: cliplist.ID,
	}, nil
}

func (s *CliplistServiceImpl) DeleteCliplist(
	ctx context.Context,
	cliplistID string,
	req *api.DeleteCliplistRequest,
) (*api.DeleteCliplistResponse, service.Error) {
	cliplist, err := s.CliplistRepository.NewQuery(app_context.GetDB(ctx)).
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
		return nil, NoPermissionToCliplist
	}

	cliplist.Status = model.CliplistStatusDeleted
	if err := s.CliplistRepository.NewQuery(app_context.GetDB(ctx)).Save(cliplist); err != nil {
		return nil, InternalError.With(err)
	}

	return &api.DeleteCliplistResponse{}, nil
}
