package service

var (
	InternalError = NewNotFoundError("Sorry. Please retry later.")
)

var (
	NoPermissionToClip     = NewForbiddenError("have no permission for the clip")
	NoPermissionToCliplist = NewForbiddenError("have no permission for the cliplist")
)

var (
	CliplistIndexOutOfRange = NewNotFoundError("The index of cliplist is out of range.")
)

var (
	ClipNotFound         = NewNotFoundError("The clip was not found.")
	CliplistNotFound     = NewNotFoundError("The cliplist was not found.")
	CliplistItemNotFound = NewNotFoundError("The item of cliplist was not found.")
)
