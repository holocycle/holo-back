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
