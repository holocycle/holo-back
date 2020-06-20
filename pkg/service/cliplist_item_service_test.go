package service

import (
	"context"
	"testing"

	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/test"
)

func Test_GetCliplistItem(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:   "normal",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
			},
			Postcondition: []interface{}{},
			Req:           &api.GetCliplistItemRequest{},
			Res: &api.GetCliplistItemResponse{
				CliplistItem: test.APICliplistItem(1, 0, test.APIVideo(1), true),
			},
			Err: nil,
		},
		{
			Name:   "cliplist item not found",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Postcondition: []interface{}{},
			Req:           &api.GetCliplistItemRequest{},
			Res:           nil,
			Err:           CliplistItemNotFound,
		},
		{
			Name:   "cliplist not found",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
			},
			Postcondition: []interface{}{},
			Req:           &api.GetCliplistItemRequest{},
			Res:           nil,
			Err:           CliplistNotFound,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistItemService()
		return service.GetCliplistItem(ctx, "cliplist01", 0, req.(*api.GetCliplistItemRequest)) // FIXME
	})
}

func Test_PostCliplistItem(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:   "normal (empty)",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Postcondition: []interface{}{
				test.ModelCliplistContain(1, 0, 1),
			},
			Req: map[string]interface{}{
				"index": 0,
				"req": &api.PostCliplistItemRequest{
					ClipID: "clip01",
				},
			},
			Res: &api.PostCliplistItemResponse{
				CliplistID: "cliplist01",
			},
			Err: nil,
		},
		{
			Name:   "normal (tail)",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelVideo(2),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelClip(2, 1, 2, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
			},
			Postcondition: []interface{}{
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplistContain(1, 1, 2),
			},
			Req: map[string]interface{}{
				"index": 1,
				"req": &api.PostCliplistItemRequest{
					ClipID: "clip02",
				},
			},
			Res: &api.PostCliplistItemResponse{
				CliplistID: "cliplist01",
			},
			Err: nil,
		},
		{
			Name:   "normal (first)",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelVideo(2),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelClip(2, 1, 2, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
			},
			Postcondition: []interface{}{
				test.ModelCliplistContain(1, 1, 1),
				test.ModelCliplistContain(1, 0, 2),
			},
			Req: map[string]interface{}{
				"index": 0,
				"req": &api.PostCliplistItemRequest{
					ClipID: "clip02",
				},
			},
			Res: &api.PostCliplistItemResponse{
				CliplistID: "cliplist01",
			},
			Err: nil,
		},
		{
			Name:   "out of range",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelVideo(2),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelClip(2, 1, 2, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Postcondition: []interface{}{},
			Req: map[string]interface{}{
				"index": 2,
				"req": &api.PostCliplistItemRequest{
					ClipID: "clip02",
				},
			},
			Res: nil,
			Err: CliplistIndexOutOfRange,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistItemService()
		m := req.(map[string]interface{})
		return service.PostCliplistItem(ctx, "cliplist01", m["index"].(int), m["req"].(*api.PostCliplistItemRequest)) // FIXME
	})
}

func Test_DeleteCliplistItem(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:   "normal",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelVideo(2),
				test.ModelVideo(3),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelClip(2, 1, 2, model.ClipStatusPublic),
				test.ModelClip(3, 1, 3, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplistContain(1, 1, 2),
				test.ModelCliplistContain(1, 2, 3),
			},
			Postcondition: []interface{}{
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplistContain(1, 1, 3),
			},
			Req: map[string]interface{}{
				"index": 1,
				"req":   &api.DeleteCliplistItemRequest{},
			},
			Res: &api.DeleteCliplistItemResponse{},
			Err: nil,
		},
		{
			Name:   "normal (empty)",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelVideo(1),
				test.ModelClip(1, 1, 1, model.ClipStatusPublic),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
				test.ModelCliplistContain(1, 0, 1),
			},
			Postcondition: []interface{}{},
			Req: map[string]interface{}{
				"index": 0,
				"req":   &api.DeleteCliplistItemRequest{},
			},
			Res: &api.DeleteCliplistItemResponse{},
			Err: nil,
		},
		{
			Name:   "out of range",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelCliplist(1, 1, model.CliplistStatusPublic),
			},
			Postcondition: []interface{}{},
			Req: map[string]interface{}{
				"index": 0,
				"req":   &api.DeleteCliplistItemRequest{},
			},
			Res: nil,
			Err: CliplistIndexOutOfRange,
		},
	}
	test.DoServiceTests(t, testcases, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		service := NewCliplistItemService()
		m := req.(map[string]interface{})
		return service.DeleteCliplistItem(ctx, "cliplist01", m["index"].(int), m["req"].(*api.DeleteCliplistItemRequest)) // FIXME
	})
}
