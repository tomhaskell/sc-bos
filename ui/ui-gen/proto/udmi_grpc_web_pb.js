/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: udmi.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.smartcore = {};
proto.smartcore.bos = require('./udmi_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.UdmiServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.UdmiServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.PullControlTopicsRequest,
 *   !proto.smartcore.bos.PullControlTopicsResponse>}
 */
const methodDescriptor_UdmiService_PullControlTopics = new grpc.web.MethodDescriptor(
  '/smartcore.bos.UdmiService/PullControlTopics',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullControlTopicsRequest,
  proto.smartcore.bos.PullControlTopicsResponse,
  /**
   * @param {!proto.smartcore.bos.PullControlTopicsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullControlTopicsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullControlTopicsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullControlTopicsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.UdmiServiceClient.prototype.pullControlTopics =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.UdmiService/PullControlTopics',
      request,
      metadata || {},
      methodDescriptor_UdmiService_PullControlTopics);
};


/**
 * @param {!proto.smartcore.bos.PullControlTopicsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullControlTopicsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.UdmiServicePromiseClient.prototype.pullControlTopics =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.UdmiService/PullControlTopics',
      request,
      metadata || {},
      methodDescriptor_UdmiService_PullControlTopics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.OnMessageRequest,
 *   !proto.smartcore.bos.OnMessageResponse>}
 */
const methodDescriptor_UdmiService_OnMessage = new grpc.web.MethodDescriptor(
  '/smartcore.bos.UdmiService/OnMessage',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.OnMessageRequest,
  proto.smartcore.bos.OnMessageResponse,
  /**
   * @param {!proto.smartcore.bos.OnMessageRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.OnMessageResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.OnMessageRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.OnMessageResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.OnMessageResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.UdmiServiceClient.prototype.onMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.UdmiService/OnMessage',
      request,
      metadata || {},
      methodDescriptor_UdmiService_OnMessage,
      callback);
};


/**
 * @param {!proto.smartcore.bos.OnMessageRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.OnMessageResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.UdmiServicePromiseClient.prototype.onMessage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.UdmiService/OnMessage',
      request,
      metadata || {},
      methodDescriptor_UdmiService_OnMessage);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.PullExportMessagesRequest,
 *   !proto.smartcore.bos.PullExportMessagesResponse>}
 */
const methodDescriptor_UdmiService_PullExportMessages = new grpc.web.MethodDescriptor(
  '/smartcore.bos.UdmiService/PullExportMessages',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullExportMessagesRequest,
  proto.smartcore.bos.PullExportMessagesResponse,
  /**
   * @param {!proto.smartcore.bos.PullExportMessagesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullExportMessagesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullExportMessagesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullExportMessagesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.UdmiServiceClient.prototype.pullExportMessages =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.UdmiService/PullExportMessages',
      request,
      metadata || {},
      methodDescriptor_UdmiService_PullExportMessages);
};


/**
 * @param {!proto.smartcore.bos.PullExportMessagesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullExportMessagesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.UdmiServicePromiseClient.prototype.pullExportMessages =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.UdmiService/PullExportMessages',
      request,
      metadata || {},
      methodDescriptor_UdmiService_PullExportMessages);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.GetExportMessageRequest,
 *   !proto.smartcore.bos.MqttMessage>}
 */
const methodDescriptor_UdmiService_GetExportMessage = new grpc.web.MethodDescriptor(
  '/smartcore.bos.UdmiService/GetExportMessage',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetExportMessageRequest,
  proto.smartcore.bos.MqttMessage,
  /**
   * @param {!proto.smartcore.bos.GetExportMessageRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.MqttMessage.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetExportMessageRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.MqttMessage)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.MqttMessage>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.UdmiServiceClient.prototype.getExportMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.UdmiService/GetExportMessage',
      request,
      metadata || {},
      methodDescriptor_UdmiService_GetExportMessage,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetExportMessageRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.MqttMessage>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.UdmiServicePromiseClient.prototype.getExportMessage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.UdmiService/GetExportMessage',
      request,
      metadata || {},
      methodDescriptor_UdmiService_GetExportMessage);
};


module.exports = proto.smartcore.bos;

