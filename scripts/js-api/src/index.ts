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
  videoId: string;
  beginAt: number;
  endAt: number;
  video: Video;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new Clip();
    result.type = source["type"];
    result.id = source["id"];
    result.title = source["title"];
    result.description = source["description"];
    result.videoId = source["videoId"];
    result.beginAt = source["beginAt"];
    result.endAt = source["endAt"];
    result.video = source["video"] ? Video.createFrom(source["video"]) : null;
    return result;
  }

}
export class ListClipsRequest {
  limit: number;
  orderBy: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListClipsRequest();
    result.limit = source["limit"];
    result.orderBy = source["orderBy"];
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
export class User {
  type: string;
  id: string;
  name: string;
  email: string;
  iconUrl: string;

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new User();
    result.type = source["type"];
    result.id = source["id"];
    result.name = source["name"];
    result.email = source["email"];
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

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListTagsRequest();
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

