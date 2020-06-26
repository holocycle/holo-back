package test

import (
	"fmt"
	"reflect"

	"github.com/holocycle/holo-back/pkg/model"
)

func ModelUser(id int) *model.User {
	user := model.NewUser(
		fmt.Sprintf("user%02d", id),
		fmt.Sprintf("user%02d@example.com", id),
		fmt.Sprintf("http://icon"),
	)
	user.ID = fmt.Sprintf("user%02d", id)
	return user
}

func ModelVideo(id int) *model.Video {
	video := model.NewVideo(
		fmt.Sprintf("video%02d", id),
		fmt.Sprintf("channel%02d", id),
		fmt.Sprintf("video %02d", id),
		fmt.Sprintf("video %02d", id),
		100,
		"http://small",
		"http://medium",
		"http://large",
		tm("2020/1/1 00:00:00"),
	)
	return video
}

func ModelClip(id, user, video int, status model.ClipStatus) *model.Clip {
	clip := model.NewClip(
		fmt.Sprintf("user%02d", user),
		fmt.Sprintf("clip %02d", id),
		fmt.Sprintf("clip %02d", id),
		fmt.Sprintf("video%02d", video),
		1,
		10,
		status,
	)
	clip.ID = fmt.Sprintf("clip%02d", id)
	return clip
}

func ModelCliplist(id, user int, status model.CliplistStatus) *model.Cliplist {
	cliplist := model.NewCliplist(
		fmt.Sprintf("user%02d", user),
		fmt.Sprintf("cliplist %02d", id),
		fmt.Sprintf("cliplist %02d", id),
		status,
	)
	cliplist.ID = fmt.Sprintf("cliplist%02d", id)
	return cliplist
}

func ModelCliplistContain(cliplist, index, clip int) *model.CliplistContain {
	cliplistContain := model.NewCliplistContain(
		fmt.Sprintf("cliplist%02d", cliplist),
		index,
		fmt.Sprintf("clip%02d", clip),
	)
	return cliplistContain
}

func NewModelSameTypeWith(model interface{}) interface{} {
	modelType := reflect.TypeOf(model).Elem()
	return reflect.New(modelType).Interface()
}

func ModelFavorite(clip, user int) *model.Favorite {
	favorite := model.NewFavorite(
		fmt.Sprintf("clip%02d", clip),
		fmt.Sprintf("user%02d", user),
	)
	return favorite
}
