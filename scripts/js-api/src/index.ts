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

  static createFrom(source: any) {
    if ('string' === typeof source) source = JSON.parse(source);
    const result = new ListClipsRequest();
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