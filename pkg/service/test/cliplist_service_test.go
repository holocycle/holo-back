package test

import (
	"testing"

	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestListCliplists(t *testing.T) {
	service := service.NewCliplistService()

	testcases := []struct {
		Name string
		Req  *api.ListCliplistsRequest
		Res  *api.ListCliplistsResponse
		Err  error
	}{
		{
			Name: "test1",
			Req: &api.ListCliplistsRequest{
				Limit:   10,
				OrderBy: "any",
			},
			Res: &api.ListCliplistsResponse{
				Cliplists: []*api.Cliplist{},
			},
			Err: nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx, rollback := GetTestHelper().NewContext()
			defer rollback()

			res, err := service.ListCliplists(ctx, tc.Req)
			assert.Equal(t, tc.Res, res)
			assert.Equal(t, tc.Err, err)
		})
	}
}
