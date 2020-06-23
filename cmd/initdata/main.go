package main

import (
	"context"
	"fmt"
	"time"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration. err=%+v\n", err)
		return
	}

	db, err := db.NewDB(&config.DB)
	defer db.Close()

	ctx := context.Background()
	ctx = app_context.SetDB(ctx, db)

	user1 := model.NewUser("user1", "test1@test.com", "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo")
	user2 := model.NewUser("user2", "test2@test.com", "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo")
	user3 := model.NewUser("user3", "test3@test.com", "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo")
	userRepository := repository.NewUserRepository()
	_ = userRepository.NewQuery(ctx).Create(user1)
	_ = userRepository.NewQuery(ctx).Create(user2)
	_ = userRepository.NewQuery(ctx).Create(user3)

	v1PublishAt := time.Date(2020, 4, 23, 17, 01, 42, 0, time.UTC).Local()
	v2PublishAt := time.Date(2020, 3, 27, 11, 0, 12, 0, time.UTC).Local()
	v3PublishAt := time.Date(2020, 4, 25, 12, 11, 57, 0, time.UTC).Local()
	video1 := model.NewVideo("xccH7xxG5zc", "UCCzUftO8KOVkV4wQG1vkUvg", "【雑談】ていうか報告することがいろいろある！？【ホロライブ/宝鐘マリン】", "よくよく考えたら告知することがいろいろあるじゃあないですか！ ", 5154, "https://i.ytimg.com/vi/xccH7xxG5zc/default.jpg", "https://i.ytimg.com/vi/xccH7xxG5zc/mqdefault.jpg", "https://i.ytimg.com/vi/xccH7xxG5zc/hqdefault.jpg", &v1PublishAt)
	video2 := model.NewVideo("X9zw0QF12Kc", "UC-hM6YJuNYVAmUWxeIr9FeA", "サクラカゼ / さくらみこ【オリジナル曲】", "さくらみこ 2ndオリジナル楽曲", 231, "https://i.ytimg.com/vi/X9zw0QF12Kc/default.jpg", "https://i.ytimg.com/vi/X9zw0QF12Kc/mqdefault.jpg", "https://i.ytimg.com/vi/X9zw0QF12Kc/hqdefault.jpg", &v2PublishAt)
	video3 := model.NewVideo("lwbGo-O6buc", "UC1opHUrw8rvnsadT-iGp7Cg", "【コラボ耐久】あくシオ💓わちゃわちゃ二人組がゼロからエンドラ倒すまで終われない！？【湊あくあ/紫咲シオン】", "隣人戦争、遂に宇宙へ―…。", 23961, "https://i.ytimg.com/vi/lwbGo-O6buc/default.jpg", "https://i.ytimg.com/vi/lwbGo-O6buc/mqdefault.jpg", "https://i.ytimg.com/vi/lwbGo-O6buc/hqdefault.jpg", &v3PublishAt)
	videoRepository := repository.NewVideoRepository()
	_ = videoRepository.NewQuery(ctx).Create(video1)
	_ = videoRepository.NewQuery(ctx).Create(video2)
	_ = videoRepository.NewQuery(ctx).Create(video3)

	clip1 := model.NewClip(user1.ID, "タイトル", "動画の説明", "xccH7xxG5zc", 2799, 2871, model.ClipStatusPublic)
	clip2 := model.NewClip(user1.ID, "サクラカゼ", "動画の説明", "X9zw0QF12Kc", 0, 100, model.ClipStatusPublic)
	clip3 := model.NewClip(user2.ID, "切り抜き動画を見る紫咲シオン", "動画の説明", "lwbGo-O6buc", 1320, 1420, model.ClipStatusPublic)
	clip4 := model.NewClip(user3.ID, "切り抜き動画を見る紫咲シオン", "動画の説明", "lwbGo-O6buc", 1320, 1420, model.ClipStatusPublic)
	clipRepository := repository.NewClipRepository()
	_ = clipRepository.NewQuery(ctx).Create(clip1)
	_ = clipRepository.NewQuery(ctx).Create(clip2)
	_ = clipRepository.NewQuery(ctx).Create(clip3)
	_ = clipRepository.NewQuery(ctx).Create(clip4)

	tags := []model.Tag{
		*model.NewTag("ときのそら", "#4374FF"),
		*model.NewTag("ロボ子さん", "#4374FF"),
		*model.NewTag("さくらみこ", "#D252FF"),
		*model.NewTag("夜空メル", "#FEA5D1"),
		*model.NewTag("白上フブキ", "#FFD200"),
		*model.NewTag("夏色まつり", "#49E5FF"),
		*model.NewTag("赤井はあと", "#FF7608"),
		*model.NewTag("アキ・ローゼンタール", "#FD00AE"),
		*model.NewTag("湊あくあ", "#8045FF"),
		*model.NewTag("百鬼あやめ", "#D252FF"),
		*model.NewTag("癒月ちょこ", "#F7002F"),
		*model.NewTag("紫咲シオン", "#FD00AE"),
		*model.NewTag("大空スバル", "#D252FF"),
		*model.NewTag("大神ミオ", "#FFD200"),
		*model.NewTag("猫又おかゆ", "#3BE898"),
		*model.NewTag("戌神ころね", "#D252FF"),
		*model.NewTag("不知火フレア", "#FFD200"),
		*model.NewTag("白銀ノエル", "#FF7608"),
		*model.NewTag("宝鐘マリン", "#FFFFFF"),
		*model.NewTag("兎田ぺこら", "#F7002F"),
		*model.NewTag("潤羽るしあ", "#49E5FF"),
		*model.NewTag("星街すいせい", "#8045FF"),
		*model.NewTag("天音かなた", "#4374FF"),
		*model.NewTag("桐生ココ", "#FFFFFF"),
		*model.NewTag("角巻わため", "#FF7608"),
		*model.NewTag("常闇トワ", "#FFD200"),
		*model.NewTag("姫森ルーナ", "#D252FF"),
	}
	tagRepository := repository.NewTagRepository()
	for i := 0; i < len(tags); i++ {
		_ = tagRepository.NewQuery(ctx).Create(&tags[i])
	}

	clipTagged1 := model.NewClipTagged(clip1.ID, tags[0].ID, clip1.UserID)
	clipTagged2 := model.NewClipTagged(clip1.ID, tags[1].ID, clip1.UserID)
	clipTagged3 := model.NewClipTagged(clip2.ID, tags[2].ID, clip2.UserID)
	clipTagged4 := model.NewClipTagged(clip3.ID, tags[3].ID, clip3.UserID)
	clipTaggedRepository := repository.NewClipTaggedRepository()
	_ = clipTaggedRepository.NewQuery(ctx).Create(clipTagged1)
	_ = clipTaggedRepository.NewQuery(ctx).Create(clipTagged2)
	_ = clipTaggedRepository.NewQuery(ctx).Create(clipTagged3)
	_ = clipTaggedRepository.NewQuery(ctx).Create(clipTagged4)

	favorite1 := model.NewFavorite(clip1.ID, user1.ID)
	favorite2 := model.NewFavorite(clip2.ID, user1.ID)
	favorite3 := model.NewFavorite(clip3.ID, user1.ID)
	favorite4 := model.NewFavorite(clip1.ID, user2.ID)
	favoriteRepository := repository.NewFavoriteRepository()
	_ = favoriteRepository.NewQuery(ctx).Create(favorite1)
	_ = favoriteRepository.NewQuery(ctx).Create(favorite2)
	_ = favoriteRepository.NewQuery(ctx).Create(favorite3)
	_ = favoriteRepository.NewQuery(ctx).Create(favorite4)

	comments1 := model.NewComment(user1.ID, clip1.ID, "1get")
	comments2 := model.NewComment(user2.ID, clip1.ID, "2get")
	comments3 := model.NewComment(user3.ID, clip1.ID, "3get")
	comments4 := model.NewComment(user1.ID, clip2.ID, "てえてえ")
	commentsRepository := repository.NewCommentRepository()
	_ = commentsRepository.NewQuery(ctx).Create(comments1)
	_ = commentsRepository.NewQuery(ctx).Create(comments2)
	_ = commentsRepository.NewQuery(ctx).Create(comments3)
	_ = commentsRepository.NewQuery(ctx).Create(comments4)

	clipList := []model.Cliplist{
		*model.NewCliplist(user1.ID, "クリップリスト1", "詳細", model.CliplistStatusPublic),
		*model.NewCliplist(user1.ID, "クリップリスト2", "詳細", model.CliplistStatusPublic),
		*model.NewCliplist(user1.ID, "クリップリスト3", "詳細", model.CliplistStatusPublic),
		*model.NewCliplist(user1.ID, "1件クリップリスト", "詳細", model.CliplistStatusPublic),
		*model.NewCliplist(user1.ID, "未登録クリップリスト", "詳細", model.CliplistStatusPublic),
	}
	clipListRepository := repository.NewCliplistRepository()
	for i := 0; i < len(clipList); i++ {
		_ = clipListRepository.NewQuery(ctx).Create(&clipList[i])
	}

	clipListContains := []model.CliplistContain{
		*model.NewCliplistContain(clipList[0].ID, 0, clip1.ID),
		*model.NewCliplistContain(clipList[0].ID, 1, clip2.ID),
		*model.NewCliplistContain(clipList[0].ID, 2, clip3.ID),
		*model.NewCliplistContain(clipList[0].ID, 3, clip4.ID),
		*model.NewCliplistContain(clipList[1].ID, 0, clip2.ID),
		*model.NewCliplistContain(clipList[1].ID, 1, clip3.ID),
		*model.NewCliplistContain(clipList[1].ID, 2, clip4.ID),
		*model.NewCliplistContain(clipList[2].ID, 0, clip1.ID),
		*model.NewCliplistContain(clipList[2].ID, 1, clip2.ID),
		*model.NewCliplistContain(clipList[2].ID, 2, clip4.ID),
		*model.NewCliplistContain(clipList[3].ID, 0, clip1.ID),
	}
	clipListContainRepository := repository.NewCliplistContainRepository()
	for i := 0; i < len(clipList); i++ {
		_ = clipListContainRepository.NewQuery(ctx).Create(&clipListContains[i])
	}

	bookmarks1 := model.NewBookMark(clipList[0].ID, user1.ID)
	bookmarks2 := model.NewBookMark(clipList[0].ID, user2.ID)
	bookmarks3 := model.NewBookMark(clipList[0].ID, user3.ID)
	bookmarks4 := model.NewBookMark(clipList[1].ID, user1.ID)
	bookmarkRepository := repository.NewBookmarkRepository()
	_ = bookmarkRepository.NewQuery(ctx).Create(bookmarks1)
	_ = bookmarkRepository.NewQuery(ctx).Create(bookmarks2)
	_ = bookmarkRepository.NewQuery(ctx).Create(bookmarks3)
	_ = bookmarkRepository.NewQuery(ctx).Create(bookmarks4)
}
