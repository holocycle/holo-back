package service

type Container struct {
	CliplistService CliplistService
}

func NewContainer() *Container {
	return &Container{
		CliplistService: NewCliplistService(),
	}
}
