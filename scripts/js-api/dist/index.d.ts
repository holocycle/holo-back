export declare class ModelBase {
    type: string;
    id: string;
    static createFrom(source: any): ModelBase;
}
export declare class Liver {
    type: string;
    id: string;
    name: string;
    channelId: string;
    mainColor: string;
    subColor: string;
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
    videoId: string;
    beginAt: number;
    endAt: number;
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
export declare class User {
    type: string;
    id: string;
    name: string;
    email: string;
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
