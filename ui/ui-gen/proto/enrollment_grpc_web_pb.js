/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
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
proto.smartcore.bos = require('./enrollment_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.EnrollmentApiClient =
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
proto.smartcore.bos.EnrollmentApiPromiseClient =
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
 *   !proto.smartcore.bos.GetEnrollmentRequest,
 *   !proto.smartcore.bos.Enrollment>}
 */
const methodDescriptor_EnrollmentApi_GetEnrollment = new grpc.web.MethodDescriptor(
  '/smartcore.bos.EnrollmentApi/GetEnrollment',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetEnrollmentRequest,
  proto.smartcore.bos.Enrollment,
  /**
   * @param {!proto.smartcore.bos.GetEnrollmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Enrollment.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetEnrollmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Enrollment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Enrollment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.EnrollmentApiClient.prototype.getEnrollment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.EnrollmentApi/GetEnrollment',
      request,
      metadata || {},
      methodDescriptor_EnrollmentApi_GetEnrollment,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetEnrollmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Enrollment>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.EnrollmentApiPromiseClient.prototype.getEnrollment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.EnrollmentApi/GetEnrollment',
      request,
      metadata || {},
      methodDescriptor_EnrollmentApi_GetEnrollment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.CreateEnrollmentRequest,
 *   !proto.smartcore.bos.Enrollment>}
 */
const methodDescriptor_EnrollmentApi_CreateEnrollment = new grpc.web.MethodDescriptor(
  '/smartcore.bos.EnrollmentApi/CreateEnrollment',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.CreateEnrollmentRequest,
  proto.smartcore.bos.Enrollment,
  /**
   * @param {!proto.smartcore.bos.CreateEnrollmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Enrollment.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.CreateEnrollmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Enrollment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Enrollment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.EnrollmentApiClient.prototype.createEnrollment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.EnrollmentApi/CreateEnrollment',
      request,
      metadata || {},
      methodDescriptor_EnrollmentApi_CreateEnrollment,
      callback);
};


/**
 * @param {!proto.smartcore.bos.CreateEnrollmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Enrollment>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.EnrollmentApiPromiseClient.prototype.createEnrollment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.EnrollmentApi/CreateEnrollment',
      request,
      metadata || {},
      methodDescriptor_EnrollmentApi_CreateEnrollment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.DeleteEnrollmentRequest,
 *   !proto.smartcore.bos.Enrollment>}
 */
const methodDescriptor_EnrollmentApi_DeleteEnrollment = new grpc.web.MethodDescriptor(
  '/smartcore.bos.EnrollmentApi/DeleteEnrollment',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.DeleteEnrollmentRequest,
  proto.smartcore.bos.Enrollment,
  /**
   * @param {!proto.smartcore.bos.DeleteEnrollmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Enrollment.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.DeleteEnrollmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Enrollment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Enrollment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.EnrollmentApiClient.prototype.deleteEnrollment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.EnrollmentApi/DeleteEnrollment',
      request,
      metadata || {},
      methodDescriptor_EnrollmentApi_DeleteEnrollment,
      callback);
};


/**
 * @param {!proto.smartcore.bos.DeleteEnrollmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Enrollment>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.EnrollmentApiPromiseClient.prototype.deleteEnrollment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.EnrollmentApi/DeleteEnrollment',
      request,
      metadata || {},
      methodDescriptor_EnrollmentApi_DeleteEnrollment);
};


module.exports = proto.smartcore.bos;

