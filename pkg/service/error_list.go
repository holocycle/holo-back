package service

import (
	"github.com/holocycle/holo-back/pkg/core/service"
)

var (
	InternalError = service.NewInternalError("Sorry. Please retry later.")
)

var (
	NoPermissionToClip     = service.NewForbiddenError("have no permission for the clip")
	NoPermissionToCliplist = service.NewForbiddenError("have no permission for the cliplist")
)

var (
	CliplistIndexOutOfRange = service.NewNotFoundError("The index of cliplist is out of range.")
)

var (
	ClipNotFound         = service.NewNotFoundError("The clip was not found.")
	CliplistNotFound     = service.NewNotFoundError("The cliplist was not found.")
	CliplistItemNotFound = service.NewNotFoundError("The item of cliplist was not found.")
)
