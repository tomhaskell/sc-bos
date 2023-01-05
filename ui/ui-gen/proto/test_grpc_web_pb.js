/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos.ew
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.smartcore = {};
proto.smartcore.bos = {};
proto.smartcore.bos.ew = require('./test_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.ew.TestApiClient =
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
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.ew.TestApiPromiseClient =
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
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.ew.GetTestRequest,
 *   !proto.smartcore.bos.ew.Test>}
 */
const methodDescriptor_TestApi_GetTest = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ew.TestApi/GetTest',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.ew.GetTestRequest,
  proto.smartcore.bos.ew.Test,
  /**
   * @param {!proto.smartcore.bos.ew.GetTestRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ew.Test.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.ew.GetTestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.ew.Test)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ew.Test>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ew.TestApiClient.prototype.getTest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ew.TestApi/GetTest',
      request,
      metadata || {},
      methodDescriptor_TestApi_GetTest,
      callback);
};


/**
 * @param {!proto.smartcore.bos.ew.GetTestRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ew.Test>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ew.TestApiPromiseClient.prototype.getTest =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ew.TestApi/GetTest',
      request,
      metadata || {},
      methodDescriptor_TestApi_GetTest);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.ew.UpdateTestRequest,
 *   !proto.smartcore.bos.ew.Test>}
 */
const methodDescriptor_TestApi_UpdateTest = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ew.TestApi/UpdateTest',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.ew.UpdateTestRequest,
  proto.smartcore.bos.ew.Test,
  /**
   * @param {!proto.smartcore.bos.ew.UpdateTestRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ew.Test.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.ew.UpdateTestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.ew.Test)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ew.Test>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ew.TestApiClient.prototype.updateTest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ew.TestApi/UpdateTest',
      request,
      metadata || {},
      methodDescriptor_TestApi_UpdateTest,
      callback);
};


/**
 * @param {!proto.smartcore.bos.ew.UpdateTestRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ew.Test>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ew.TestApiPromiseClient.prototype.updateTest =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ew.TestApi/UpdateTest',
      request,
      metadata || {},
      methodDescriptor_TestApi_UpdateTest);
};


module.exports = proto.smartcore.bos.ew;

