"use strict";
/* Do not change, this code is generated from Golang structs */
exports.__esModule = true;
var ModelBase = /** @class */ (function () {
    function ModelBase() {
    }
    ModelBase.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ModelBase();
        result.type = source["type"];
        result.id = source["id"];
        return result;
    };
    return ModelBase;
}());
exports.ModelBase = ModelBase;
var Time = /** @class */ (function () {
    function Time() {
    }
    Time.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Time();
        return result;
    };
    return Time;
}());
exports.Time = Time;
var Channel = /** @class */ (function () {
    function Channel() {
    }
    Channel.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Channel();
        result.type = source["type"];
        result.id = source["id"];
        result.title = source["title"];
        result.description = source["description"];
        result.smallThumbnailUrl = source["smallThumbnailUrl"];
        result.mediumThumbnailUrl = source["mediumThumbnailUrl"];
        result.largeThumbnailUrl = source["largeThumbnailUrl"];
        result.smallBannerUrl = source["smallBannerUrl"];
        result.mediumBannerUrl = source["mediumBannerUrl"];
        result.largeBannerUrl = source["largeBannerUrl"];
        result.viewCount = source["viewCount"];
        result.commentCount = source["commentCount"];
        result.subscriberCount = source["subscriberCount"];
        result.videoCount = source["videoCount"];
        result.publishedAt = source["publishedAt"] ? Time.createFrom(source["publishedAt"]) : null;
        return result;
    };
    return Channel;
}());
exports.Channel = Channel;
var Liver = /** @class */ (function () {
    function Liver() {
    }
    Liver.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Liver();
        result.type = source["type"];
        result.id = source["id"];
        result.name = source["name"];
        result.mainColor = source["mainColor"];
        result.subColor = source["subColor"];
        result.channel = source["channel"] ? Channel.createFrom(source["channel"]) : null;
        return result;
    };
    return Liver;
}());
exports.Liver = Liver;
var ListLiversRequest = /** @class */ (function () {
    function ListLiversRequest() {
    }
    ListLiversRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListLiversRequest();
        return result;
    };
    return ListLiversRequest;
}());
exports.ListLiversRequest = ListLiversRequest;
var ListLiversResponse = /** @class */ (function () {
    function ListLiversResponse() {
    }
    ListLiversResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListLiversResponse();
        result.livers = source["livers"] ? source["livers"].map(function (element) { return Liver.createFrom(element); }) : null;
        return result;
    };
    return ListLiversResponse;
}());
exports.ListLiversResponse = ListLiversResponse;
var GetLiverRequest = /** @class */ (function () {
    function GetLiverRequest() {
    }
    GetLiverRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetLiverRequest();
        return result;
    };
    return GetLiverRequest;
}());
exports.GetLiverRequest = GetLiverRequest;
var GetLiverResponse = /** @class */ (function () {
    function GetLiverResponse() {
    }
    GetLiverResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetLiverResponse();
        result.liver = source["liver"] ? Liver.createFrom(source["liver"]) : null;
        return result;
    };
    return GetLiverResponse;
}());
exports.GetLiverResponse = GetLiverResponse;
var Video = /** @class */ (function () {
    function Video() {
    }
    Video.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Video();
        result.type = source["type"];
        result.id = source["id"];
        result.channelId = source["channelId"];
        result.title = source["title"];
        result.description = source["description"];
        result.duration = source["duration"];
        result.smallThumnailUrl = source["smallThumnailUrl"];
        result.mediumThumnailUrl = source["mediumThumnailUrl"];
        result.largeThumnailUrl = source["largeThumnailUrl"];
        result.publishedAt = new Date(source["publishedAt"]);
        return result;
    };
    return Video;
}());
exports.Video = Video;
var Clip = /** @class */ (function () {
    function Clip() {
    }
    Clip.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Clip();
        result.type = source["type"];
        result.id = source["id"];
        result.title = source["title"];
        result.description = source["description"];
        result.videoId = source["videoId"];
        result.beginAt = source["beginAt"];
        result.endAt = source["endAt"];
        result.video = source["video"] ? Video.createFrom(source["video"]) : null;
        return result;
    };
    return Clip;
}());
exports.Clip = Clip;
var ListClipsRequest = /** @class */ (function () {
    function ListClipsRequest() {
    }
    ListClipsRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListClipsRequest();
        result.limit = source["limit"];
        result.orderBy = source["orderBy"];
        return result;
    };
    return ListClipsRequest;
}());
exports.ListClipsRequest = ListClipsRequest;
var ListClipsResponse = /** @class */ (function () {
    function ListClipsResponse() {
    }
    ListClipsResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListClipsResponse();
        result.clips = source["clips"] ? source["clips"].map(function (element) { return Clip.createFrom(element); }) : null;
        return result;
    };
    return ListClipsResponse;
}());
exports.ListClipsResponse = ListClipsResponse;
var PostClipRequest = /** @class */ (function () {
    function PostClipRequest() {
    }
    PostClipRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostClipRequest();
        result.videoId = source["videoId"];
        result.title = source["title"];
        result.description = source["description"];
        result.beginAt = source["beginAt"];
        result.endAt = source["endAt"];
        return result;
    };
    return PostClipRequest;
}());
exports.PostClipRequest = PostClipRequest;
var PostClipResponse = /** @class */ (function () {
    function PostClipResponse() {
    }
    PostClipResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostClipResponse();
        result.clipId = source["clipId"];
        return result;
    };
    return PostClipResponse;
}());
exports.PostClipResponse = PostClipResponse;
var GetClipRequest = /** @class */ (function () {
    function GetClipRequest() {
    }
    GetClipRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetClipRequest();
        return result;
    };
    return GetClipRequest;
}());
exports.GetClipRequest = GetClipRequest;
var GetClipResponse = /** @class */ (function () {
    function GetClipResponse() {
    }
    GetClipResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetClipResponse();
        result.clip = source["clip"] ? Clip.createFrom(source["clip"]) : null;
        return result;
    };
    return GetClipResponse;
}());
exports.GetClipResponse = GetClipResponse;
var User = /** @class */ (function () {
    function User() {
    }
    User.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new User();
        result.type = source["type"];
        result.id = source["id"];
        result.name = source["name"];
        result.email = source["email"];
        result.iconUrl = source["iconUrl"];
        return result;
    };
    return User;
}());
exports.User = User;
var Comment = /** @class */ (function () {
    function Comment() {
    }
    Comment.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Comment();
        result.type = source["type"];
        result.id = source["id"];
        result.userId = source["userId"];
        result.clipId = source["clipId"];
        result.content = source["content"];
        result.user = source["user"] ? User.createFrom(source["user"]) : null;
        return result;
    };
    return Comment;
}());
exports.Comment = Comment;
var ListCommentsRequest = /** @class */ (function () {
    function ListCommentsRequest() {
    }
    ListCommentsRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListCommentsRequest();
        result.limit = source["limit"];
        result.orderBy = source["orderBy"];
        return result;
    };
    return ListCommentsRequest;
}());
exports.ListCommentsRequest = ListCommentsRequest;
var ListCommentsResponse = /** @class */ (function () {
    function ListCommentsResponse() {
    }
    ListCommentsResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListCommentsResponse();
        result.comments = source["comments"] ? source["comments"].map(function (element) { return Comment.createFrom(element); }) : null;
        return result;
    };
    return ListCommentsResponse;
}());
exports.ListCommentsResponse = ListCommentsResponse;
var GetCommentRequest = /** @class */ (function () {
    function GetCommentRequest() {
    }
    GetCommentRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetCommentRequest();
        return result;
    };
    return GetCommentRequest;
}());
exports.GetCommentRequest = GetCommentRequest;
var GetCommentResponse = /** @class */ (function () {
    function GetCommentResponse() {
    }
    GetCommentResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetCommentResponse();
        result.comment = source["comment"] ? Comment.createFrom(source["comment"]) : null;
        return result;
    };
    return GetCommentResponse;
}());
exports.GetCommentResponse = GetCommentResponse;
var PostCommentRequest = /** @class */ (function () {
    function PostCommentRequest() {
    }
    PostCommentRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostCommentRequest();
        result.content = source["content"];
        return result;
    };
    return PostCommentRequest;
}());
exports.PostCommentRequest = PostCommentRequest;
var PostCommentResponse = /** @class */ (function () {
    function PostCommentResponse() {
    }
    PostCommentResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostCommentResponse();
        result.commentId = source["commentId"];
        return result;
    };
    return PostCommentResponse;
}());
exports.PostCommentResponse = PostCommentResponse;
var DeleteCommentRequest = /** @class */ (function () {
    function DeleteCommentRequest() {
    }
    DeleteCommentRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteCommentRequest();
        return result;
    };
    return DeleteCommentRequest;
}());
exports.DeleteCommentRequest = DeleteCommentRequest;
var DeleteCommentResponse = /** @class */ (function () {
    function DeleteCommentResponse() {
    }
    DeleteCommentResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteCommentResponse();
        return result;
    };
    return DeleteCommentResponse;
}());
exports.DeleteCommentResponse = DeleteCommentResponse;
var Tag = /** @class */ (function () {
    function Tag() {
    }
    Tag.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Tag();
        result.type = source["type"];
        result.id = source["id"];
        result.name = source["name"];
        result.color = source["color"];
        return result;
    };
    return Tag;
}());
exports.Tag = Tag;
var ListTagsRequest = /** @class */ (function () {
    function ListTagsRequest() {
    }
    ListTagsRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListTagsRequest();
        return result;
    };
    return ListTagsRequest;
}());
exports.ListTagsRequest = ListTagsRequest;
var ListTagsResponse = /** @class */ (function () {
    function ListTagsResponse() {
    }
    ListTagsResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListTagsResponse();
        result.tags = source["tags"] ? source["tags"].map(function (element) { return Tag.createFrom(element); }) : null;
        return result;
    };
    return ListTagsResponse;
}());
exports.ListTagsResponse = ListTagsResponse;
var GetTagRequest = /** @class */ (function () {
    function GetTagRequest() {
    }
    GetTagRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetTagRequest();
        return result;
    };
    return GetTagRequest;
}());
exports.GetTagRequest = GetTagRequest;
var GetTagResponse = /** @class */ (function () {
    function GetTagResponse() {
    }
    GetTagResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetTagResponse();
        result.tag = source["tag"] ? Tag.createFrom(source["tag"]) : null;
        return result;
    };
    return GetTagResponse;
}());
exports.GetTagResponse = GetTagResponse;
var PutTagRequest = /** @class */ (function () {
    function PutTagRequest() {
    }
    PutTagRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutTagRequest();
        result.name = source["name"];
        result.color = source["color"];
        return result;
    };
    return PutTagRequest;
}());
exports.PutTagRequest = PutTagRequest;
var PutTagResponse = /** @class */ (function () {
    function PutTagResponse() {
    }
    PutTagResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutTagResponse();
        result.tagId = source["tagId"];
        return result;
    };
    return PutTagResponse;
}());
exports.PutTagResponse = PutTagResponse;
var ListTagsOnClipRequest = /** @class */ (function () {
    function ListTagsOnClipRequest() {
    }
    ListTagsOnClipRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListTagsOnClipRequest();
        return result;
    };
    return ListTagsOnClipRequest;
}());
exports.ListTagsOnClipRequest = ListTagsOnClipRequest;
var ListTagsOnClipResponse = /** @class */ (function () {
    function ListTagsOnClipResponse() {
    }
    ListTagsOnClipResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListTagsOnClipResponse();
        result.clipId = source["clipId"];
        result.tags = source["tags"] ? source["tags"].map(function (element) { return Tag.createFrom(element); }) : null;
        return result;
    };
    return ListTagsOnClipResponse;
}());
exports.ListTagsOnClipResponse = ListTagsOnClipResponse;
var PutTagOnClipRequest = /** @class */ (function () {
    function PutTagOnClipRequest() {
    }
    PutTagOnClipRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutTagOnClipRequest();
        return result;
    };
    return PutTagOnClipRequest;
}());
exports.PutTagOnClipRequest = PutTagOnClipRequest;
var PutTagOnClipResponse = /** @class */ (function () {
    function PutTagOnClipResponse() {
    }
    PutTagOnClipResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutTagOnClipResponse();
        return result;
    };
    return PutTagOnClipResponse;
}());
exports.PutTagOnClipResponse = PutTagOnClipResponse;
var DeleteTagOnClipRequest = /** @class */ (function () {
    function DeleteTagOnClipRequest() {
    }
    DeleteTagOnClipRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteTagOnClipRequest();
        return result;
    };
    return DeleteTagOnClipRequest;
}());
exports.DeleteTagOnClipRequest = DeleteTagOnClipRequest;
var DeleteTagOnClipResponse = /** @class */ (function () {
    function DeleteTagOnClipResponse() {
    }
    DeleteTagOnClipResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteTagOnClipResponse();
        return result;
    };
    return DeleteTagOnClipResponse;
}());
exports.DeleteTagOnClipResponse = DeleteTagOnClipResponse;
var Favorite = /** @class */ (function () {
    function Favorite() {
    }
    Favorite.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Favorite();
        result.type = source["type"];
        result.id = source["id"];
        result.clipId = source["clipId"];
        result.userId = source["userId"];
        return result;
    };
    return Favorite;
}());
exports.Favorite = Favorite;
var PutFavoriteRequest = /** @class */ (function () {
    function PutFavoriteRequest() {
    }
    PutFavoriteRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutFavoriteRequest();
        return result;
    };
    return PutFavoriteRequest;
}());
exports.PutFavoriteRequest = PutFavoriteRequest;
var PutFavoriteResponse = /** @class */ (function () {
    function PutFavoriteResponse() {
    }
    PutFavoriteResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutFavoriteResponse();
        return result;
    };
    return PutFavoriteResponse;
}());
exports.PutFavoriteResponse = PutFavoriteResponse;
var DeleteFavoriteRequest = /** @class */ (function () {
    function DeleteFavoriteRequest() {
    }
    DeleteFavoriteRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteFavoriteRequest();
        return result;
    };
    return DeleteFavoriteRequest;
}());
exports.DeleteFavoriteRequest = DeleteFavoriteRequest;
var DeleteFavoriteResponse = /** @class */ (function () {
    function DeleteFavoriteResponse() {
    }
    DeleteFavoriteResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteFavoriteResponse();
        return result;
    };
    return DeleteFavoriteResponse;
}());
exports.DeleteFavoriteResponse = DeleteFavoriteResponse;
