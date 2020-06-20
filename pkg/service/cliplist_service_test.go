package service

import (
	"context"
	"testing"

	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/test"
)

func Test_ListCliplists(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:          "no data",
			UserID:        "user01",
			Precondition:  []interface{}{},
			Postcondition: []interface{}{},
			Req: &api.ListCliplistsRequest{
				Limit:   10,
				OrderBy: "any",
			},
			Res: &api.ListCliplistsResponse{
				Cliplists: []*api.Cliplist{},
			},
			Err: nil,
		},
		{
			Name:   "normal",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplist(2, 1, model.CliplistStatusPublic),
				test.ModelCliplist(3, 1, model.CliplistStatusDeleted),
			},
			Postcondition: []interface{}{},
			Req: &api.ListCliplistsRequest{
				Limit:   10,
				OrderBy: "any",
			},
			Res: &api.ListCliplistsResponse{
				Cliplists: []*api.Cliplist{
					test.APICliplist(1, 1, test.APICliplistItem(1, 0, test.APIVideo(1), true)),
					test.APICliplist(2, 0, nil),
				},
			},
			Err: nil,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistService()
		return service.ListCliplists(ctx, req.(*api.ListCliplistsRequest))
	})
}

func Test_GetCliplist(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:        "normal",
			UserID:      "user01",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelVideo(2),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelClip(2, 1, 1, model.ClipStatusDeleted),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplistContain(1, 1, 2),
			},
			Postcondition: []interface{}{},
			Req: &api.GetCliplistRequest{
				Page:        0,
				ItemPerPage: 10,
			},
			Res: &api.GetCliplistResponse{
				Cliplist: test.APICliplist(1, 2, test.APICliplistItem(1, 0, test.APIVideo(1), true)),
				PageInfo: &api.PageInfo{
					TotalPage:   1,
					CurrentPage: 0,
					ItemPerPage: 10,
				},
				CliplistItems: []*api.CliplistItem{
					test.APICliplistItem(1, 0, test.APIVideo(1), true),
					test.APICliplistItem(2, 0, test.APIVideo(2), false),
				},
			},
			Err: nil,
		},
		{
			Name:        "paging",
			UserID:      "user01",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelVideo(2),
				test.ModelVideo(3),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelClip(2, 1, 1, model.ClipStatusPublic),
				test.ModelClip(3, 1, 1, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplistContain(1, 1, 2),
				test.ModelCliplistContain(1, 2, 3),
			},
			Postcondition: []interface{}{},
			Req: &api.GetCliplistRequest{
				Page:        1,
				ItemPerPage: 2,
			},
			Res: &api.GetCliplistResponse{
				Cliplist: test.APICliplist(1, 3, test.APICliplistItem(1, 0, test.APIVideo(1), true)),
				PageInfo: &api.PageInfo{
					TotalPage:   2,
					CurrentPage: 1,
					ItemPerPage: 2,
				},
				CliplistItems: []*api.CliplistItem{
					test.APICliplistItem(3, 0, test.APIVideo(1), true),
				},
			},
			Err: nil,
		},
		{
			Name:          "not found",
			UserID:        "user01",
			IDGenerator:   test.NewIDGenerator(),
			Precondition:  []interface{}{},
			Postcondition: []interface{}{},
			Req: &api.GetCliplistRequest{
				Page:        0,
				ItemPerPage: 10,
			},
			Res: nil,
			Err: CliplistNotFound,
		},
		{
			Name:        "deleted",
			UserID:      "user01",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelCliplist(1, 1, model.CliplistStatusDeleted),
			},
			Postcondition: []interface{}{},
			Req: &api.GetCliplistRequest{
				Page:        0,
				ItemPerPage: 10,
			},
			Res: nil,
			Err: CliplistNotFound,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistService()
		return service.GetCliplist(ctx, "cliplist01", req.(*api.GetCliplistRequest)) // FIXME
	})
}

func Test_PostCliplist(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:   "normal",
			UserID: "user01",
			IDGenerator: test.NewIDGenerator(
				"cliplist01",
			),
			Precondition: []interface{}{
				test.ModelUser(1),
			},
			Postcondition: []interface{}{
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Req: &api.PostCliplistRequest{
				Title:       "cliplist 01",
				Description: "cliplist 01",
			},
			Res: &api.PostCliplistResponse{
				CliplistID: "cliplist01",
			},
			Err: nil,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistService()
		return service.PostCliplist(ctx, req.(*api.PostCliplistRequest))
	})
}

func Test_PutCliplist(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:        "normal",
			UserID:      "user01",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Postcondition: []interface{}{
				&model.Cliplist{
					ID:          test.ModelCliplist(1, 1, model.CliplistStatusPublic).ID,
					UserID:      test.ModelCliplist(1, 1, model.CliplistStatusPublic).UserID,
					Title:       "modified-title",
					Description: "modified-description",
					Status:      model.CliplistStatusPublic,
				},
			},
			Req: &api.PutCliplistRequest{
				Title:       "modified-title",
				Description: "modified-description",
			},
			Res: &api.PutCliplistResponse{
				CliplistID: "cliplist01",
			},
			Err: nil,
		},
		{
			Name:        "not-found",
			UserID:      "user01",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
			},
			Postcondition: []interface{}{},
			Req: &api.PutCliplistRequest{
				Title:       "modified-title",
				Description: "modified-description",
			},
			Res: nil,
			Err: CliplistNotFound,
		},
		{
			Name:        "deleted",
			UserID:      "user01",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelCliplist(1, 1, model.CliplistStatusDeleted),
			},
			Postcondition: []interface{}{},
			Req: &api.PutCliplistRequest{
				Title:       "modified-title",
				Description: "modified-description",
			},
			Res: nil,
			Err: CliplistNotFound,
		},
		{
			Name:        "forbidden",
			UserID:      "user02",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Postcondition: []interface{}{},
			Req: &api.PutCliplistRequest{
				Title:       "modified-title",
				Description: "modified-description",
			},
			Res: nil,
			Err: NoPermissionToCliplist,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistService()
		return service.PutCliplist(ctx, "cliplist01", req.(*api.PutCliplistRequest)) // FIXME
	})
}

func Test_DeleteCliplist(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:        "normal",
			UserID:      "user01",
			IDGenerator: test.NewIDGenerator(),
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Postcondition: []interface{}{
				test.ModelCliplist(1, 1, model.CliplistStatusDeleted),
			},
			Req: &api.DeleteCliplistRequest{},
			Res: &api.DeleteCliplistResponse{},
			Err: nil,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistService()
		return service.DeleteCliplist(ctx, "cliplist01", req.(*api.DeleteCliplistRequest)) // FIXME
	})
}
