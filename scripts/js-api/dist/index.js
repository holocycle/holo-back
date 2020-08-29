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
var PageInfo = /** @class */ (function () {
    function PageInfo() {
    }
    PageInfo.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PageInfo();
        result.totalPage = source["totalPage"];
        result.currentPage = source["currentPage"];
        result.itemPerPage = source["itemPerPage"];
        return result;
    };
    return PageInfo;
}());
exports.PageInfo = PageInfo;
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
        result.beginAt = source["beginAt"];
        result.endAt = source["endAt"];
        result.favoriteCount = source["favoriteCount"];
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
        result.tags = source["tags"];
        result.createdBy = source["createdBy"];
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
var PutClipRequest = /** @class */ (function () {
    function PutClipRequest() {
    }
    PutClipRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutClipRequest();
        result.title = source["title"];
        result.description = source["description"];
        result.beginAt = source["beginAt"];
        result.endAt = source["endAt"];
        return result;
    };
    return PutClipRequest;
}());
exports.PutClipRequest = PutClipRequest;
var PutClipResponse = /** @class */ (function () {
    function PutClipResponse() {
    }
    PutClipResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutClipResponse();
        result.clipId = source["clipId"];
        return result;
    };
    return PutClipResponse;
}());
exports.PutClipResponse = PutClipResponse;
var DeleteClipRequest = /** @class */ (function () {
    function DeleteClipRequest() {
    }
    DeleteClipRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteClipRequest();
        return result;
    };
    return DeleteClipRequest;
}());
exports.DeleteClipRequest = DeleteClipRequest;
var DeleteClipResponse = /** @class */ (function () {
    function DeleteClipResponse() {
    }
    DeleteClipResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteClipResponse();
        return result;
    };
    return DeleteClipResponse;
}());
exports.DeleteClipResponse = DeleteClipResponse;
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
        result.key = source["key"];
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
var LoginUser = /** @class */ (function () {
    function LoginUser() {
    }
    LoginUser.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new LoginUser();
        result.type = source["type"];
        result.id = source["id"];
        result.name = source["name"];
        result.iconUrl = source["iconUrl"];
        result.email = source["email"];
        return result;
    };
    return LoginUser;
}());
exports.LoginUser = LoginUser;
var ListUserRequest = /** @class */ (function () {
    function ListUserRequest() {
    }
    ListUserRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListUserRequest();
        result.limit = source["limit"];
        result.orderBy = source["orderBy"];
        return result;
    };
    return ListUserRequest;
}());
exports.ListUserRequest = ListUserRequest;
var ListUserResponse = /** @class */ (function () {
    function ListUserResponse() {
    }
    ListUserResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListUserResponse();
        result.users = source["users"] ? source["users"].map(function (element) { return User.createFrom(element); }) : null;
        return result;
    };
    return ListUserResponse;
}());
exports.ListUserResponse = ListUserResponse;
var GetUserRequest = /** @class */ (function () {
    function GetUserRequest() {
    }
    GetUserRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetUserRequest();
        return result;
    };
    return GetUserRequest;
}());
exports.GetUserRequest = GetUserRequest;
var GetUserResponse = /** @class */ (function () {
    function GetUserResponse() {
    }
    GetUserResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetUserResponse();
        result.user = source["user"] ? User.createFrom(source["user"]) : null;
        return result;
    };
    return GetUserResponse;
}());
exports.GetUserResponse = GetUserResponse;
var GetLoginUserRequest = /** @class */ (function () {
    function GetLoginUserRequest() {
    }
    GetLoginUserRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetLoginUserRequest();
        return result;
    };
    return GetLoginUserRequest;
}());
exports.GetLoginUserRequest = GetLoginUserRequest;
var GetLoginUserResponse = /** @class */ (function () {
    function GetLoginUserResponse() {
    }
    GetLoginUserResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetLoginUserResponse();
        result.loginUser = source["loginUser"] ? LoginUser.createFrom(source["loginUser"]) : null;
        return result;
    };
    return GetLoginUserResponse;
}());
exports.GetLoginUserResponse = GetLoginUserResponse;
var GetUserFavoritesRequest = /** @class */ (function () {
    function GetUserFavoritesRequest() {
    }
    GetUserFavoritesRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetUserFavoritesRequest();
        return result;
    };
    return GetUserFavoritesRequest;
}());
exports.GetUserFavoritesRequest = GetUserFavoritesRequest;
var GetUserFavoritesResponse = /** @class */ (function () {
    function GetUserFavoritesResponse() {
    }
    GetUserFavoritesResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetUserFavoritesResponse();
        result.favoriteClips = source["favoriteClips"] ? source["favoriteClips"].map(function (element) { return Clip.createFrom(element); }) : null;
        return result;
    };
    return GetUserFavoritesResponse;
}());
exports.GetUserFavoritesResponse = GetUserFavoritesResponse;
var GetFavoriteRequest = /** @class */ (function () {
    function GetFavoriteRequest() {
    }
    GetFavoriteRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetFavoriteRequest();
        return result;
    };
    return GetFavoriteRequest;
}());
exports.GetFavoriteRequest = GetFavoriteRequest;
var GetFavoriteResponse = /** @class */ (function () {
    function GetFavoriteResponse() {
    }
    GetFavoriteResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetFavoriteResponse();
        result.favorite = source["favorite"];
        return result;
    };
    return GetFavoriteResponse;
}());
exports.GetFavoriteResponse = GetFavoriteResponse;
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
var CliplistItem = /** @class */ (function () {
    function CliplistItem() {
    }
    CliplistItem.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new CliplistItem();
        result.type = source["type"];
        result.id = source["id"];
        result.title = source["title"];
        result.description = source["description"];
        result.beginAt = source["beginAt"];
        result.endAt = source["endAt"];
        result.favoriteCount = source["favoriteCount"];
        result.video = source["video"] ? Video.createFrom(source["video"]) : null;
        result.available = source["available"];
        return result;
    };
    return CliplistItem;
}());
exports.CliplistItem = CliplistItem;
var Cliplist = /** @class */ (function () {
    function Cliplist() {
    }
    Cliplist.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new Cliplist();
        result.type = source["type"];
        result.id = source["id"];
        result.title = source["title"];
        result.description = source["description"];
        result.length = source["length"];
        result.firstItem = source["firstItem"] ? CliplistItem.createFrom(source["firstItem"]) : null;
        return result;
    };
    return Cliplist;
}());
exports.Cliplist = Cliplist;
var ListCliplistsRequest = /** @class */ (function () {
    function ListCliplistsRequest() {
    }
    ListCliplistsRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListCliplistsRequest();
        result.limit = source["limit"];
        result.orderBy = source["orderBy"];
        return result;
    };
    return ListCliplistsRequest;
}());
exports.ListCliplistsRequest = ListCliplistsRequest;
var ListCliplistsResponse = /** @class */ (function () {
    function ListCliplistsResponse() {
    }
    ListCliplistsResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new ListCliplistsResponse();
        result.cliplists = source["cliplists"] ? source["cliplists"].map(function (element) { return Cliplist.createFrom(element); }) : null;
        return result;
    };
    return ListCliplistsResponse;
}());
exports.ListCliplistsResponse = ListCliplistsResponse;
var GetCliplistRequest = /** @class */ (function () {
    function GetCliplistRequest() {
    }
    GetCliplistRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetCliplistRequest();
        result.page = source["page"];
        result.itemPerPage = source["itemPerPage"];
        return result;
    };
    return GetCliplistRequest;
}());
exports.GetCliplistRequest = GetCliplistRequest;
var GetCliplistResponse = /** @class */ (function () {
    function GetCliplistResponse() {
    }
    GetCliplistResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetCliplistResponse();
        result.cliplist = source["cliplist"] ? Cliplist.createFrom(source["cliplist"]) : null;
        result.pageInfo = source["pageInfo"] ? PageInfo.createFrom(source["pageInfo"]) : null;
        result.cliplistItems = source["cliplistItems"] ? source["cliplistItems"].map(function (element) { return CliplistItem.createFrom(element); }) : null;
        return result;
    };
    return GetCliplistResponse;
}());
exports.GetCliplistResponse = GetCliplistResponse;
var PostCliplistRequest = /** @class */ (function () {
    function PostCliplistRequest() {
    }
    PostCliplistRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostCliplistRequest();
        result.title = source["title"];
        result.description = source["description"];
        return result;
    };
    return PostCliplistRequest;
}());
exports.PostCliplistRequest = PostCliplistRequest;
var PostCliplistResponse = /** @class */ (function () {
    function PostCliplistResponse() {
    }
    PostCliplistResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostCliplistResponse();
        result.cliplistId = source["cliplistId"];
        return result;
    };
    return PostCliplistResponse;
}());
exports.PostCliplistResponse = PostCliplistResponse;
var PutCliplistRequest = /** @class */ (function () {
    function PutCliplistRequest() {
    }
    PutCliplistRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutCliplistRequest();
        result.title = source["title"];
        result.description = source["description"];
        return result;
    };
    return PutCliplistRequest;
}());
exports.PutCliplistRequest = PutCliplistRequest;
var PutCliplistResponse = /** @class */ (function () {
    function PutCliplistResponse() {
    }
    PutCliplistResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PutCliplistResponse();
        result.cliplistId = source["cliplistId"];
        return result;
    };
    return PutCliplistResponse;
}());
exports.PutCliplistResponse = PutCliplistResponse;
var DeleteCliplistRequest = /** @class */ (function () {
    function DeleteCliplistRequest() {
    }
    DeleteCliplistRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteCliplistRequest();
        return result;
    };
    return DeleteCliplistRequest;
}());
exports.DeleteCliplistRequest = DeleteCliplistRequest;
var DeleteCliplistResponse = /** @class */ (function () {
    function DeleteCliplistResponse() {
    }
    DeleteCliplistResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteCliplistResponse();
        return result;
    };
    return DeleteCliplistResponse;
}());
exports.DeleteCliplistResponse = DeleteCliplistResponse;
var GetCliplistItemRequest = /** @class */ (function () {
    function GetCliplistItemRequest() {
    }
    GetCliplistItemRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetCliplistItemRequest();
        return result;
    };
    return GetCliplistItemRequest;
}());
exports.GetCliplistItemRequest = GetCliplistItemRequest;
var GetCliplistItemResponse = /** @class */ (function () {
    function GetCliplistItemResponse() {
    }
    GetCliplistItemResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new GetCliplistItemResponse();
        result.cliplistItem = source["cliplistItem"] ? CliplistItem.createFrom(source["cliplistItem"]) : null;
        return result;
    };
    return GetCliplistItemResponse;
}());
exports.GetCliplistItemResponse = GetCliplistItemResponse;
var PostCliplistItemRequest = /** @class */ (function () {
    function PostCliplistItemRequest() {
    }
    PostCliplistItemRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostCliplistItemRequest();
        result.clipId = source["clipId"];
        return result;
    };
    return PostCliplistItemRequest;
}());
exports.PostCliplistItemRequest = PostCliplistItemRequest;
var PostCliplistItemResponse = /** @class */ (function () {
    function PostCliplistItemResponse() {
    }
    PostCliplistItemResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new PostCliplistItemResponse();
        result.cliplistId = source["cliplistId"];
        return result;
    };
    return PostCliplistItemResponse;
}());
exports.PostCliplistItemResponse = PostCliplistItemResponse;
var DeleteCliplistItemRequest = /** @class */ (function () {
    function DeleteCliplistItemRequest() {
    }
    DeleteCliplistItemRequest.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteCliplistItemRequest();
        return result;
    };
    return DeleteCliplistItemRequest;
}());
exports.DeleteCliplistItemRequest = DeleteCliplistItemRequest;
var DeleteCliplistItemResponse = /** @class */ (function () {
    function DeleteCliplistItemResponse() {
    }
    DeleteCliplistItemResponse.createFrom = function (source) {
        if ('string' === typeof source)
            source = JSON.parse(source);
        var result = new DeleteCliplistItemResponse();
        result.cliplistId = source["cliplistId"];
        return result;
    };
    return DeleteCliplistItemResponse;
}());
exports.DeleteCliplistItemResponse = DeleteCliplistItemResponse;
