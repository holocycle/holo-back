package service

import (
	"context"

	"github.com/holocycle/holo-back/pkg/api"
	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
)

type CliplistService interface {
	ListCliplists(ctx context.Context, req *api.ListCliplistsRequest) (*api.ListCliplistsResponse, error)
}

type CliplistServiceImpl struct {
	CliplistRepository repository.CliplistRepository
}

func NewCliplistService() CliplistService {
	return &CliplistServiceImpl{
		CliplistRepository: repository.NewCliplistRepository(),
	}
}

func (s *CliplistServiceImpl) ListCliplists(ctx context.Context, req *api.ListCliplistsRequest) (*api.ListCliplistsResponse, error) {
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
		return nil, err
	}

	return &api.ListCliplistsResponse{
		Cliplists: converter.ConvertToCliplists(cliplist),
	}, nil
}
