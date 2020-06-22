package converter

import "github.com/holocycle/holo-back/pkg/api"

func ConvertToPageInfo(total, page, itemPerPage int) *api.PageInfo {
	return &api.PageInfo{
		TotalPage:   (total + itemPerPage - 1) / itemPerPage, // round up
		CurrentPage: page,
		ItemPerPage: itemPerPage,
	}
}
