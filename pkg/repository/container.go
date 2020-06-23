package repository

type Container struct {
	LiverRepository           LiverRepository
	ChannelRepository         ChannelRepository
	VideoRepository           VideoRepository
	UserRepository            UserRepository
	SessionRepository         SessionRepository
	ClipRepository            ClipRepository
	FavoriteRepository        FavoriteRepository
	CommentRepository         CommentRepository
	TagRepository             TagRepository
	ClipTaggedRepository      ClipTaggedRepository
	CliplistRepository        CliplistRepository
	CliplistContainRepository CliplistContainRepository
	BookmarkRepository        BookmarkRepository
}

func NewContainer() *Container {
	return &Container{
		LiverRepository:           NewLiverRepository(),
		ChannelRepository:         NewChannelRepository(),
		VideoRepository:           NewVideoRepository(),
		UserRepository:            NewUserRepository(),
		SessionRepository:         NewSessionRepository(),
		ClipRepository:            NewClipRepository(),
		FavoriteRepository:        NewFavoriteRepository(),
		CommentRepository:         NewCommentRepository(),
		TagRepository:             NewTagRepository(),
		ClipTaggedRepository:      NewClipTaggedRepository(),
		CliplistRepository:        NewCliplistRepository(),
		CliplistContainRepository: NewCliplistContainRepository(),
		BookmarkRepository:        NewBookmarkRepository(),
	}
}
