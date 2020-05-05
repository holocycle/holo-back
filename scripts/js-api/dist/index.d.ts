export declare class ModelBase {
    type: string;
    id: string;
    static createFrom(source: any): ModelBase;
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
