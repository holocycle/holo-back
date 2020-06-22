package service

import (
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
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				0,
				&api.GetCliplistItemRequest{},
			},
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
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				0,
				&api.GetCliplistItemRequest{},
			},
			Res: nil,
			Err: CliplistItemNotFound,
		},
		{
			Name:   "cliplist not found",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
			},
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				0,
				&api.GetCliplistItemRequest{},
			},
			Res: nil,
			Err: CliplistNotFound,
		},
	}
	test.DoServiceTests(t, testcases, NewCliplistItemService().GetCliplistItem)
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
			ExpectCreation: []interface{}{
				test.ModelCliplistContain(1, 0, 1),
			},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				0,
				&api.PostCliplistItemRequest{
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
			ExpectCreation: []interface{}{
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplistContain(1, 1, 2),
			},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				1,
				&api.PostCliplistItemRequest{
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
			ExpectCreation: []interface{}{
				test.ModelCliplistContain(1, 1, 1),
				test.ModelCliplistContain(1, 0, 2),
			},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				0,
				&api.PostCliplistItemRequest{
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
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				2,
				&api.PostCliplistItemRequest{
					ClipID: "clip02",
				},
			},
			Res: nil,
			Err: CliplistIndexOutOfRange,
		},
	}
	test.DoServiceTests(t, testcases, NewCliplistItemService().PostCliplistItem)
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
			ExpectCreation: []interface{}{
				test.ModelCliplistContain(1, 0, 1),
				test.ModelCliplistContain(1, 1, 3),
			},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				1,
				&api.DeleteCliplistItemRequest{},
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
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				0,
				&api.DeleteCliplistItemRequest{},
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
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"cliplist01",
				0,
				&api.DeleteCliplistItemRequest{},
			},
			Res: nil,
			Err: CliplistIndexOutOfRange,
		},
	}
	test.DoServiceTests(t, testcases, NewCliplistItemService().DeleteCliplistItem)
}
