package service

import (
	"testing"

	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/test"
)

func Test_GetFavorite(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:   "unfavorite",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelUser(2),
				test.ModelVideo(1),
				test.ModelClip(1, 2, 1, model.ClipStatusPublic),
			},
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"clip01",
				&api.GetFavoriteRequest{},
			},
			Res: &api.GetFavoriteResponse{
				Favorite: UNFAVORITE,
			},
			Err: nil,
		},
		{
			Name:   "favorite",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelUser(2),
				test.ModelVideo(1),
				test.ModelClip(1, 2, 1, model.ClipStatusPublic),
				test.ModelFavorite(1, 1),
			},
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"clip01",
				&api.GetFavoriteRequest{},
			},
			Res: &api.GetFavoriteResponse{
				Favorite: FAVORITE,
			},
			Err: nil,
		},
	}
	test.DoServiceTests(t, testcases, NewFavoriteService().GetFavorite)
}

func Test_PutFavorite(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:   "unfavorite",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelUser(2),
				test.ModelVideo(1),
				test.ModelClip(1, 2, 1, model.ClipStatusPublic),
			},
			ExpectCreation: []interface{}{
				test.ModelFavorite(1, 1),
			},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"clip01",
				&api.PutFavoriteRequest{},
			},
			Res: &api.PutFavoriteResponse{},
			Err: nil,
		},
		{
			Name:   "favorite",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelUser(2),
				test.ModelVideo(1),
				test.ModelClip(1, 2, 1, model.ClipStatusPublic),
				test.ModelFavorite(1, 1),
			},
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"clip01",
				&api.PutFavoriteRequest{},
			},
			Res: &api.PutFavoriteResponse{},
			Err: nil,
		},
	}
	test.DoServiceTests(t, testcases, NewFavoriteService().PutFavorite)
}

func Test_DeleteFavorite(t *testing.T) {
	testcases := []test.ServiceTestcase{
		{
			Name:   "unfavorite",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelUser(2),
				test.ModelVideo(1),
				test.ModelClip(1, 2, 1, model.ClipStatusPublic),
			},
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{},
			Req: []interface{}{
				"clip01",
				&api.DeleteFavoriteRequest{},
			},
			Res: &api.DeleteFavoriteResponse{},
			Err: nil,
		},
		{
			Name:   "favorite",
			UserID: "user01",
			Precondition: []interface{}{
				test.ModelUser(1),
				test.ModelUser(2),
				test.ModelVideo(1),
				test.ModelClip(1, 2, 1, model.ClipStatusPublic),
			},
			ExpectCreation: []interface{}{},
			ExpectDeletion: []interface{}{
				test.ModelFavorite(1, 1),
			},
			Req: []interface{}{
				"clip01",
				&api.DeleteFavoriteRequest{},
			},
			Res: &api.DeleteFavoriteResponse{},
			Err: nil,
		},
	}
	test.DoServiceTests(t, testcases, NewFavoriteService().DeleteFavorite)
}
