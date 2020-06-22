package service

type Container struct {
	CliplistService     CliplistService
	CliplistItemService CliplistItemService
}

func NewContainer() *Container {
	return &Container{
		CliplistService:     NewCliplistService(),
		CliplistItemService: NewCliplistItemService(),
	}
}
