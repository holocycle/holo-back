package service

type Container struct {
	CliplistService     CliplistService
	CliplistItemService CliplistItemService
	FavoriteService     FavoriteService
}

func NewContainer() *Container {
	return &Container{
		CliplistService:     NewCliplistService(),
		CliplistItemService: NewCliplistItemService(),
		FavoriteService:     NewFavoriteService(),
	}
}
