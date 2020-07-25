/* Do not change, this code is generated from Golang structs */


export class ModelBase {
  type: string;
  id: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ModelBase();
    result.type = source["type"];
    result.id = source["id"];
    return result;
  }

}
export class PageInfo {
  totalPage: number;
  currentPage: number;
  itemPerPage: number;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PageInfo();
    result.totalPage = source["totalPage"];
    result.currentPage = source["currentPage"];
    result.itemPerPage = source["itemPerPage"];
    return result;
  }

}
export class Time {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Time();
    return result;
  }

}
export class Channel {
  type: string;
  id: string;
  title: string;
  description: string;
  smallThumbnailUrl: string;
  mediumThumbnailUrl: string;
  largeThumbnailUrl: string;
  smallBannerUrl: string;
  mediumBannerUrl: string;
  largeBannerUrl: string;
  viewCount: number;
  commentCount: number;
  subscriberCount: number;
  videoCount: number;
  publishedAt: Time;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Channel();
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
  }

}
export class Liver {
  type: string;
  id: string;
  name: string;
  mainColor: string;
  subColor: string;
  channel: Channel;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Liver();
    result.type = source["type"];
    result.id = source["id"];
    result.name = source["name"];
    result.mainColor = source["mainColor"];
    result.subColor = source["subColor"];
    result.channel = source["channel"] ? Channel.createFrom(source["channel"]) : null;
    return result;
  }

}

export class ListLiversRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListLiversRequest();
    return result;
  }

}
export class ListLiversResponse {
  livers: Liver[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListLiversResponse();
    result.livers = source["livers"] ? source["livers"].map(function(element: any) { return Liver.createFrom(element); }) : null;
    return result;
  }

}
export class GetLiverRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetLiverRequest();
    return result;
  }

}
export class GetLiverResponse {
  liver: Liver;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetLiverResponse();
    result.liver = source["liver"] ? Liver.createFrom(source["liver"]) : null;
    return result;
  }

}
export class Video {
  type: string;
  id: string;
  channelId: string;
  title: string;
  description: string;
  duration: number;
  smallThumnailUrl: string;
  mediumThumnailUrl: string;
  largeThumnailUrl: string;
  publishedAt: Date;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Video();
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
  }

}
export class Clip {
  type: string;
  id: string;
  title: string;
  description: string;
  beginAt: number;
  endAt: number;
  favoriteCount: number;
  video: Video;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Clip();
    result.type = source["type"];
    result.id = source["id"];
    result.title = source["title"];
    result.description = source["description"];
    result.beginAt = source["beginAt"];
    result.endAt = source["endAt"];
    result.favoriteCount = source["favoriteCount"];
    result.video = source["video"] ? Video.createFrom(source["video"]) : null;
    return result;
  }

}
export class ClipFilter {
  tags: string[];
  createdBy: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ClipFilter();
    result.tags = source["tags"];
    result.createdBy = source["createdBy"];
    return result;
  }

}
export class ListClipsRequest {
  limit: number;
  orderBy: string;
  filter: ClipFilter;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListClipsRequest();
    result.limit = source["limit"];
    result.orderBy = source["orderBy"];
    result.filter = source["filter"] ? ClipFilter.createFrom(source["filter"]) : null;
    return result;
  }

}
export class ListClipsResponse {
  clips: Clip[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListClipsResponse();
    result.clips = source["clips"] ? source["clips"].map(function(element: any) { return Clip.createFrom(element); }) : null;
    return result;
  }

}
export class PostClipRequest {
  videoId: string;
  title: string;
  description: string;
  beginAt: number;
  endAt: number;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostClipRequest();
    result.videoId = source["videoId"];
    result.title = source["title"];
    result.description = source["description"];
    result.beginAt = source["beginAt"];
    result.endAt = source["endAt"];
    return result;
  }

}
export class PostClipResponse {
  clipId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostClipResponse();
    result.clipId = source["clipId"];
    return result;
  }

}
export class GetClipRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetClipRequest();
    return result;
  }

}
export class GetClipResponse {
  clip: Clip;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetClipResponse();
    result.clip = source["clip"] ? Clip.createFrom(source["clip"]) : null;
    return result;
  }

}
export class PutClipRequest {
  title: string;
  description: string;
  beginAt: number;
  endAt: number;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutClipRequest();
    result.title = source["title"];
    result.description = source["description"];
    result.beginAt = source["beginAt"];
    result.endAt = source["endAt"];
    return result;
  }

}
export class PutClipResponse {
  clipId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutClipResponse();
    result.clipId = source["clipId"];
    return result;
  }

}
export class DeleteClipRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteClipRequest();
    return result;
  }

}
export class DeleteClipResponse {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteClipResponse();
    return result;
  }

}
export class User {
  type: string;
  id: string;
  name: string;
  iconUrl: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new User();
    result.type = source["type"];
    result.id = source["id"];
    result.name = source["name"];
    result.iconUrl = source["iconUrl"];
    return result;
  }

}
export class Comment {
  type: string;
  id: string;
  userId: string;
  clipId: string;
  content: string;
  user: User;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Comment();
    result.type = source["type"];
    result.id = source["id"];
    result.userId = source["userId"];
    result.clipId = source["clipId"];
    result.content = source["content"];
    result.user = source["user"] ? User.createFrom(source["user"]) : null;
    return result;
  }

}
export class ListCommentsRequest {
  limit: number;
  orderBy: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListCommentsRequest();
    result.limit = source["limit"];
    result.orderBy = source["orderBy"];
    return result;
  }

}
export class ListCommentsResponse {
  comments: Comment[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListCommentsResponse();
    result.comments = source["comments"] ? source["comments"].map(function(element: any) { return Comment.createFrom(element); }) : null;
    return result;
  }

}
export class GetCommentRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetCommentRequest();
    return result;
  }

}
export class GetCommentResponse {
  comment: Comment;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetCommentResponse();
    result.comment = source["comment"] ? Comment.createFrom(source["comment"]) : null;
    return result;
  }

}
export class PostCommentRequest {
  content: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostCommentRequest();
    result.content = source["content"];
    return result;
  }

}
export class PostCommentResponse {
  commentId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostCommentResponse();
    result.commentId = source["commentId"];
    return result;
  }

}
export class DeleteCommentRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteCommentRequest();
    return result;
  }

}
export class DeleteCommentResponse {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteCommentResponse();
    return result;
  }

}
export class Tag {
  type: string;
  id: string;
  name: string;
  color: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Tag();
    result.type = source["type"];
    result.id = source["id"];
    result.name = source["name"];
    result.color = source["color"];
    return result;
  }

}
export class ListTagsRequest {
  key: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListTagsRequest();
    result.key = source["key"];
    return result;
  }

}
export class ListTagsResponse {
  tags: Tag[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListTagsResponse();
    result.tags = source["tags"] ? source["tags"].map(function(element: any) { return Tag.createFrom(element); }) : null;
    return result;
  }

}
export class GetTagRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetTagRequest();
    return result;
  }

}
export class GetTagResponse {
  tag: Tag;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetTagResponse();
    result.tag = source["tag"] ? Tag.createFrom(source["tag"]) : null;
    return result;
  }

}
export class PutTagRequest {
  name: string;
  color: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutTagRequest();
    result.name = source["name"];
    result.color = source["color"];
    return result;
  }

}
export class PutTagResponse {
  tagId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutTagResponse();
    result.tagId = source["tagId"];
    return result;
  }

}
export class ListTagsOnClipRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListTagsOnClipRequest();
    return result;
  }

}
export class ListTagsOnClipResponse {
  clipId: string;
  tags: Tag[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListTagsOnClipResponse();
    result.clipId = source["clipId"];
    result.tags = source["tags"] ? source["tags"].map(function(element: any) { return Tag.createFrom(element); }) : null;
    return result;
  }

}
export class PutTagOnClipRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutTagOnClipRequest();
    return result;
  }

}
export class PutTagOnClipResponse {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutTagOnClipResponse();
    return result;
  }

}
export class DeleteTagOnClipRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteTagOnClipRequest();
    return result;
  }

}
export class DeleteTagOnClipResponse {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteTagOnClipResponse();
    return result;
  }

}

export class LoginUser {
  type: string;
  id: string;
  name: string;
  iconUrl: string;
  email: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new LoginUser();
    result.type = source["type"];
    result.id = source["id"];
    result.name = source["name"];
    result.iconUrl = source["iconUrl"];
    result.email = source["email"];
    return result;
  }

}
export class ListUserRequest {
  limit: number;
  orderBy: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListUserRequest();
    result.limit = source["limit"];
    result.orderBy = source["orderBy"];
    return result;
  }

}
export class ListUserResponse {
  users: User[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListUserResponse();
    result.users = source["users"] ? source["users"].map(function(element: any) { return User.createFrom(element); }) : null;
    return result;
  }

}
export class GetUserRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetUserRequest();
    return result;
  }

}
export class GetUserResponse {
  user: User;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetUserResponse();
    result.user = source["user"] ? User.createFrom(source["user"]) : null;
    return result;
  }

}
export class GetLoginUserRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetLoginUserRequest();
    return result;
  }

}
export class GetLoginUserResponse {
  loginUser: LoginUser;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetLoginUserResponse();
    result.loginUser = source["loginUser"] ? LoginUser.createFrom(source["loginUser"]) : null;
    return result;
  }

}
export class GetUserFavoritesRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetUserFavoritesRequest();
    return result;
  }

}
export class GetUserFavoritesResponse {
  favoriteClips: Clip[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetUserFavoritesResponse();
    result.favoriteClips = source["favoriteClips"] ? source["favoriteClips"].map(function(element: any) { return Clip.createFrom(element); }) : null;
    return result;
  }

}

export class GetFavoriteRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetFavoriteRequest();
    return result;
  }

}
export class GetFavoriteResponse {
  favorite: boolean;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetFavoriteResponse();
    result.favorite = source["favorite"];
    return result;
  }

}
export class PutFavoriteRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutFavoriteRequest();
    return result;
  }

}
export class PutFavoriteResponse {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutFavoriteResponse();
    return result;
  }

}
export class DeleteFavoriteRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteFavoriteRequest();
    return result;
  }

}
export class DeleteFavoriteResponse {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteFavoriteResponse();
    return result;
  }

}
export class CliplistItem {
  type: string;
  id: string;
  title: string;
  description: string;
  beginAt: number;
  endAt: number;
  favoriteCount: number;
  video: Video;
  available: boolean;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new CliplistItem();
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
  }

}
export class Cliplist {
  type: string;
  id: string;
  title: string;
  description: string;
  length: number;
  firstItem: CliplistItem;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Cliplist();
    result.type = source["type"];
    result.id = source["id"];
    result.title = source["title"];
    result.description = source["description"];
    result.length = source["length"];
    result.firstItem = source["firstItem"] ? CliplistItem.createFrom(source["firstItem"]) : null;
    return result;
  }

}

export class ListCliplistsRequest {
  limit: number;
  orderBy: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListCliplistsRequest();
    result.limit = source["limit"];
    result.orderBy = source["orderBy"];
    return result;
  }

}
export class ListCliplistsResponse {
  cliplists: Cliplist[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListCliplistsResponse();
    result.cliplists = source["cliplists"] ? source["cliplists"].map(function(element: any) { return Cliplist.createFrom(element); }) : null;
    return result;
  }

}
export class GetCliplistRequest {
  page: number;
  itemPerPage: number;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetCliplistRequest();
    result.page = source["page"];
    result.itemPerPage = source["itemPerPage"];
    return result;
  }

}
export class GetCliplistResponse {
  cliplist: Cliplist;
  pageInfo: PageInfo;
  cliplistItems: CliplistItem[];

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetCliplistResponse();
    result.cliplist = source["cliplist"] ? Cliplist.createFrom(source["cliplist"]) : null;
    result.pageInfo = source["pageInfo"] ? PageInfo.createFrom(source["pageInfo"]) : null;
    result.cliplistItems = source["cliplistItems"] ? source["cliplistItems"].map(function(element: any) { return CliplistItem.createFrom(element); }) : null;
    return result;
  }

}
export class PostCliplistRequest {
  title: string;
  description: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostCliplistRequest();
    result.title = source["title"];
    result.description = source["description"];
    return result;
  }

}
export class PostCliplistResponse {
  cliplistId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostCliplistResponse();
    result.cliplistId = source["cliplistId"];
    return result;
  }

}
export class PutCliplistRequest {
  title: string;
  description: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutCliplistRequest();
    result.title = source["title"];
    result.description = source["description"];
    return result;
  }

}
export class PutCliplistResponse {
  cliplistId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PutCliplistResponse();
    result.cliplistId = source["cliplistId"];
    return result;
  }

}
export class DeleteCliplistRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteCliplistRequest();
    return result;
  }

}
export class DeleteCliplistResponse {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteCliplistResponse();
    return result;
  }

}
export class GetCliplistItemRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetCliplistItemRequest();
    return result;
  }

}
export class GetCliplistItemResponse {
  cliplistItem: CliplistItem;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new GetCliplistItemResponse();
    result.cliplistItem = source["cliplistItem"] ? CliplistItem.createFrom(source["cliplistItem"]) : null;
    return result;
  }

}
export class PostCliplistItemRequest {
  clipId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostCliplistItemRequest();
    result.clipId = source["clipId"];
    return result;
  }

}
export class PostCliplistItemResponse {
  cliplistId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new PostCliplistItemResponse();
    result.cliplistId = source["cliplistId"];
    return result;
  }

}
export class DeleteCliplistItemRequest {

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteCliplistItemRequest();
    return result;
  }

}
export class DeleteCliplistItemResponse {
  cliplistId: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new DeleteCliplistItemResponse();
    result.cliplistId = source["cliplistId"];
    return result;
  }

}