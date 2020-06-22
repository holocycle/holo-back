export declare class ModelBase {
    type: string;
    id: string;
    static createFrom(source: any): ModelBase;
}
export declare class Time {
    static createFrom(source: any): Time;
}
export declare class Channel {
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
    static createFrom(source: any): Channel;
}
export declare class Liver {
    type: string;
    id: string;
    name: string;
    mainColor: string;
    subColor: string;
    channel: Channel;
    static createFrom(source: any): Liver;
}
export declare class ListLiversRequest {
    static createFrom(source: any): ListLiversRequest;
}
export declare class ListLiversResponse {
    livers: Liver[];
    static createFrom(source: any): ListLiversResponse;
}
export declare class GetLiverRequest {
    static createFrom(source: any): GetLiverRequest;
}
export declare class GetLiverResponse {
    liver: Liver;
    static createFrom(source: any): GetLiverResponse;
}
export declare class Video {
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
    static createFrom(source: any): Video;
}
export declare class Clip {
    type: string;
    id: string;
    title: string;
    description: string;
    beginAt: number;
    endAt: number;
    favoriteCount: number;
    video: Video;
    static createFrom(source: any): Clip;
}
export declare class ListClipsRequest {
    limit: number;
    orderBy: string;
    static createFrom(source: any): ListClipsRequest;
}
export declare class ListClipsResponse {
    clips: Clip[];
    static createFrom(source: any): ListClipsResponse;
}
export declare class PostClipRequest {
    videoId: string;
    title: string;
    description: string;
    beginAt: number;
    endAt: number;
    static createFrom(source: any): PostClipRequest;
}
export declare class PostClipResponse {
    clipId: string;
    static createFrom(source: any): PostClipResponse;
}
export declare class GetClipRequest {
    static createFrom(source: any): GetClipRequest;
}
export declare class GetClipResponse {
    clip: Clip;
    static createFrom(source: any): GetClipResponse;
}
export declare class PutClipRequest {
    title: string;
    description: string;
    beginAt: number;
    endAt: number;
    static createFrom(source: any): PutClipRequest;
}
export declare class PutClipResponse {
    clipId: string;
    static createFrom(source: any): PutClipResponse;
}
export declare class DeleteClipRequest {
    static createFrom(source: any): DeleteClipRequest;
}
export declare class DeleteClipResponse {
    static createFrom(source: any): DeleteClipResponse;
}
export declare class User {
    type: string;
    id: string;
    name: string;
    iconUrl: string;
    static createFrom(source: any): User;
}
export declare class Comment {
    type: string;
    id: string;
    userId: string;
    clipId: string;
    content: string;
    user: User;
    static createFrom(source: any): Comment;
}
export declare class ListCommentsRequest {
    limit: number;
    orderBy: string;
    static createFrom(source: any): ListCommentsRequest;
}
export declare class ListCommentsResponse {
    comments: Comment[];
    static createFrom(source: any): ListCommentsResponse;
}
export declare class GetCommentRequest {
    static createFrom(source: any): GetCommentRequest;
}
export declare class GetCommentResponse {
    comment: Comment;
    static createFrom(source: any): GetCommentResponse;
}
export declare class PostCommentRequest {
    content: string;
    static createFrom(source: any): PostCommentRequest;
}
export declare class PostCommentResponse {
    commentId: string;
    static createFrom(source: any): PostCommentResponse;
}
export declare class DeleteCommentRequest {
    static createFrom(source: any): DeleteCommentRequest;
}
export declare class DeleteCommentResponse {
    static createFrom(source: any): DeleteCommentResponse;
}
export declare class Tag {
    type: string;
    id: string;
    name: string;
    color: string;
    static createFrom(source: any): Tag;
}
export declare class ListTagsRequest {
    static createFrom(source: any): ListTagsRequest;
}
export declare class ListTagsResponse {
    tags: Tag[];
    static createFrom(source: any): ListTagsResponse;
}
export declare class GetTagRequest {
    static createFrom(source: any): GetTagRequest;
}
export declare class GetTagResponse {
    tag: Tag;
    static createFrom(source: any): GetTagResponse;
}
export declare class PutTagRequest {
    name: string;
    color: string;
    static createFrom(source: any): PutTagRequest;
}
export declare class PutTagResponse {
    tagId: string;
    static createFrom(source: any): PutTagResponse;
}
export declare class ListTagsOnClipRequest {
    static createFrom(source: any): ListTagsOnClipRequest;
}
export declare class ListTagsOnClipResponse {
    clipId: string;
    tags: Tag[];
    static createFrom(source: any): ListTagsOnClipResponse;
}
export declare class PutTagOnClipRequest {
    static createFrom(source: any): PutTagOnClipRequest;
}
export declare class PutTagOnClipResponse {
    static createFrom(source: any): PutTagOnClipResponse;
}
export declare class DeleteTagOnClipRequest {
    static createFrom(source: any): DeleteTagOnClipRequest;
}
export declare class DeleteTagOnClipResponse {
    static createFrom(source: any): DeleteTagOnClipResponse;
}
export declare class LoginUser {
    type: string;
    id: string;
    name: string;
    iconUrl: string;
    email: string;
    static createFrom(source: any): LoginUser;
}
export declare class ListUserRequest {
    limit: number;
    orderBy: string;
    static createFrom(source: any): ListUserRequest;
}
export declare class ListUserResponse {
    users: User[];
    static createFrom(source: any): ListUserResponse;
}
export declare class GetUserRequest {
    static createFrom(source: any): GetUserRequest;
}
export declare class GetUserResponse {
    user: User;
    static createFrom(source: any): GetUserResponse;
}
export declare class GetLoginUserRequest {
    static createFrom(source: any): GetLoginUserRequest;
}
export declare class GetLoginUserResponse {
    loginUser: LoginUser;
    static createFrom(source: any): GetLoginUserResponse;
}
export declare class GetUserFavoritesRequest {
    static createFrom(source: any): GetUserFavoritesRequest;
}
export declare class GetUserFavoritesResponse {
    favoriteClips: Clip[];
    static createFrom(source: any): GetUserFavoritesResponse;
}
export declare class PutFavoriteRequest {
    static createFrom(source: any): PutFavoriteRequest;
}
export declare class PutFavoriteResponse {
    static createFrom(source: any): PutFavoriteResponse;
}
export declare class DeleteFavoriteRequest {
    static createFrom(source: any): DeleteFavoriteRequest;
}
export declare class DeleteFavoriteResponse {
    static createFrom(source: any): DeleteFavoriteResponse;
}
