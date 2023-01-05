/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos.tenants
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.bos = {};
proto.smartcore.bos.tenants = require('./tenants_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.tenants.TenantApiClient =
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
proto.smartcore.bos.tenants.TenantApiPromiseClient =
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
 *   !proto.smartcore.bos.tenants.ListTenantsRequest,
 *   !proto.smartcore.bos.tenants.ListTenantsResponse>}
 */
const methodDescriptor_TenantApi_ListTenants = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/ListTenants',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.ListTenantsRequest,
  proto.smartcore.bos.tenants.ListTenantsResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.ListTenantsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.ListTenantsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.ListTenantsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.ListTenantsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.ListTenantsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.listTenants =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/ListTenants',
      request,
      metadata || {},
      methodDescriptor_TenantApi_ListTenants,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.ListTenantsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.ListTenantsResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.listTenants =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/ListTenants',
      request,
      metadata || {},
      methodDescriptor_TenantApi_ListTenants);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.PullTenantsRequest,
 *   !proto.smartcore.bos.tenants.PullTenantsResponse>}
 */
const methodDescriptor_TenantApi_PullTenants = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/PullTenants',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.tenants.PullTenantsRequest,
  proto.smartcore.bos.tenants.PullTenantsResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.PullTenantsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.PullTenantsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.PullTenantsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullTenantsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.pullTenants =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullTenants',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullTenants);
};


/**
 * @param {!proto.smartcore.bos.tenants.PullTenantsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullTenantsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.pullTenants =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullTenants',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullTenants);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.CreateTenantRequest,
 *   !proto.smartcore.bos.tenants.Tenant>}
 */
const methodDescriptor_TenantApi_CreateTenant = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/CreateTenant',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.CreateTenantRequest,
  proto.smartcore.bos.tenants.Tenant,
  /**
   * @param {!proto.smartcore.bos.tenants.CreateTenantRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Tenant.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.CreateTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Tenant)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Tenant>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.createTenant =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/CreateTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_CreateTenant,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.CreateTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Tenant>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.createTenant =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/CreateTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_CreateTenant);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.GetTenantRequest,
 *   !proto.smartcore.bos.tenants.Tenant>}
 */
const methodDescriptor_TenantApi_GetTenant = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/GetTenant',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.GetTenantRequest,
  proto.smartcore.bos.tenants.Tenant,
  /**
   * @param {!proto.smartcore.bos.tenants.GetTenantRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Tenant.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.GetTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Tenant)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Tenant>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.getTenant =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/GetTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_GetTenant,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.GetTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Tenant>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.getTenant =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/GetTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_GetTenant);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.UpdateTenantRequest,
 *   !proto.smartcore.bos.tenants.Tenant>}
 */
const methodDescriptor_TenantApi_UpdateTenant = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/UpdateTenant',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.UpdateTenantRequest,
  proto.smartcore.bos.tenants.Tenant,
  /**
   * @param {!proto.smartcore.bos.tenants.UpdateTenantRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Tenant.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.UpdateTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Tenant)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Tenant>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.updateTenant =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/UpdateTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_UpdateTenant,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.UpdateTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Tenant>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.updateTenant =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/UpdateTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_UpdateTenant);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.DeleteTenantRequest,
 *   !proto.smartcore.bos.tenants.DeleteTenantResponse>}
 */
const methodDescriptor_TenantApi_DeleteTenant = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/DeleteTenant',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.DeleteTenantRequest,
  proto.smartcore.bos.tenants.DeleteTenantResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.DeleteTenantRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.DeleteTenantResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.DeleteTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.DeleteTenantResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.DeleteTenantResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.deleteTenant =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/DeleteTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_DeleteTenant,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.DeleteTenantRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.DeleteTenantResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.deleteTenant =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/DeleteTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_DeleteTenant);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.PullTenantRequest,
 *   !proto.smartcore.bos.tenants.PullTenantResponse>}
 */
const methodDescriptor_TenantApi_PullTenant = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/PullTenant',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.tenants.PullTenantRequest,
  proto.smartcore.bos.tenants.PullTenantResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.PullTenantRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.PullTenantResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.PullTenantRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullTenantResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.pullTenant =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullTenant);
};


/**
 * @param {!proto.smartcore.bos.tenants.PullTenantRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullTenantResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.pullTenant =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullTenant',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullTenant);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.AddTenantZonesRequest,
 *   !proto.smartcore.bos.tenants.Tenant>}
 */
const methodDescriptor_TenantApi_AddTenantZones = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/AddTenantZones',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.AddTenantZonesRequest,
  proto.smartcore.bos.tenants.Tenant,
  /**
   * @param {!proto.smartcore.bos.tenants.AddTenantZonesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Tenant.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.AddTenantZonesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Tenant)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Tenant>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.addTenantZones =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/AddTenantZones',
      request,
      metadata || {},
      methodDescriptor_TenantApi_AddTenantZones,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.AddTenantZonesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Tenant>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.addTenantZones =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/AddTenantZones',
      request,
      metadata || {},
      methodDescriptor_TenantApi_AddTenantZones);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.RemoveTenantZonesRequest,
 *   !proto.smartcore.bos.tenants.Tenant>}
 */
const methodDescriptor_TenantApi_RemoveTenantZones = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/RemoveTenantZones',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.RemoveTenantZonesRequest,
  proto.smartcore.bos.tenants.Tenant,
  /**
   * @param {!proto.smartcore.bos.tenants.RemoveTenantZonesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Tenant.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.RemoveTenantZonesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Tenant)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Tenant>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.removeTenantZones =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/RemoveTenantZones',
      request,
      metadata || {},
      methodDescriptor_TenantApi_RemoveTenantZones,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.RemoveTenantZonesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Tenant>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.removeTenantZones =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/RemoveTenantZones',
      request,
      metadata || {},
      methodDescriptor_TenantApi_RemoveTenantZones);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.ListSecretsRequest,
 *   !proto.smartcore.bos.tenants.ListSecretsResponse>}
 */
const methodDescriptor_TenantApi_ListSecrets = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/ListSecrets',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.ListSecretsRequest,
  proto.smartcore.bos.tenants.ListSecretsResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.ListSecretsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.ListSecretsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.ListSecretsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.ListSecretsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.ListSecretsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.listSecrets =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/ListSecrets',
      request,
      metadata || {},
      methodDescriptor_TenantApi_ListSecrets,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.ListSecretsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.ListSecretsResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.listSecrets =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/ListSecrets',
      request,
      metadata || {},
      methodDescriptor_TenantApi_ListSecrets);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.PullSecretsRequest,
 *   !proto.smartcore.bos.tenants.PullSecretsResponse>}
 */
const methodDescriptor_TenantApi_PullSecrets = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/PullSecrets',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.tenants.PullSecretsRequest,
  proto.smartcore.bos.tenants.PullSecretsResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.PullSecretsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.PullSecretsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.PullSecretsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullSecretsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.pullSecrets =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullSecrets',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullSecrets);
};


/**
 * @param {!proto.smartcore.bos.tenants.PullSecretsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullSecretsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.pullSecrets =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullSecrets',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullSecrets);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.CreateSecretRequest,
 *   !proto.smartcore.bos.tenants.Secret>}
 */
const methodDescriptor_TenantApi_CreateSecret = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/CreateSecret',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.CreateSecretRequest,
  proto.smartcore.bos.tenants.Secret,
  /**
   * @param {!proto.smartcore.bos.tenants.CreateSecretRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Secret.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.CreateSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Secret)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Secret>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.createSecret =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/CreateSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_CreateSecret,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.CreateSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Secret>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.createSecret =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/CreateSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_CreateSecret);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.VerifySecretRequest,
 *   !proto.smartcore.bos.tenants.Secret>}
 */
const methodDescriptor_TenantApi_VerifySecret = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/VerifySecret',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.VerifySecretRequest,
  proto.smartcore.bos.tenants.Secret,
  /**
   * @param {!proto.smartcore.bos.tenants.VerifySecretRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Secret.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.VerifySecretRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Secret)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Secret>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.verifySecret =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/VerifySecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_VerifySecret,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.VerifySecretRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Secret>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.verifySecret =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/VerifySecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_VerifySecret);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.GetSecretRequest,
 *   !proto.smartcore.bos.tenants.Secret>}
 */
const methodDescriptor_TenantApi_GetSecret = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/GetSecret',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.GetSecretRequest,
  proto.smartcore.bos.tenants.Secret,
  /**
   * @param {!proto.smartcore.bos.tenants.GetSecretRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Secret.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.GetSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Secret)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Secret>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.getSecret =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/GetSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_GetSecret,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.GetSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Secret>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.getSecret =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/GetSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_GetSecret);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.UpdateSecretRequest,
 *   !proto.smartcore.bos.tenants.Secret>}
 */
const methodDescriptor_TenantApi_UpdateSecret = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/UpdateSecret',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.UpdateSecretRequest,
  proto.smartcore.bos.tenants.Secret,
  /**
   * @param {!proto.smartcore.bos.tenants.UpdateSecretRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Secret.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.UpdateSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Secret)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Secret>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.updateSecret =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/UpdateSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_UpdateSecret,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.UpdateSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Secret>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.updateSecret =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/UpdateSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_UpdateSecret);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.DeleteSecretRequest,
 *   !proto.smartcore.bos.tenants.DeleteSecretResponse>}
 */
const methodDescriptor_TenantApi_DeleteSecret = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/DeleteSecret',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.DeleteSecretRequest,
  proto.smartcore.bos.tenants.DeleteSecretResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.DeleteSecretRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.DeleteSecretResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.DeleteSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.DeleteSecretResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.DeleteSecretResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.deleteSecret =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/DeleteSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_DeleteSecret,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.DeleteSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.DeleteSecretResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.deleteSecret =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/DeleteSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_DeleteSecret);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.PullSecretRequest,
 *   !proto.smartcore.bos.tenants.PullSecretResponse>}
 */
const methodDescriptor_TenantApi_PullSecret = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/PullSecret',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.tenants.PullSecretRequest,
  proto.smartcore.bos.tenants.PullSecretResponse,
  /**
   * @param {!proto.smartcore.bos.tenants.PullSecretRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.PullSecretResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.PullSecretRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullSecretResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.pullSecret =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullSecret);
};


/**
 * @param {!proto.smartcore.bos.tenants.PullSecretRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.PullSecretResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.pullSecret =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/PullSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_PullSecret);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.tenants.RegenerateSecretRequest,
 *   !proto.smartcore.bos.tenants.Secret>}
 */
const methodDescriptor_TenantApi_RegenerateSecret = new grpc.web.MethodDescriptor(
  '/smartcore.bos.tenants.TenantApi/RegenerateSecret',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.tenants.RegenerateSecretRequest,
  proto.smartcore.bos.tenants.Secret,
  /**
   * @param {!proto.smartcore.bos.tenants.RegenerateSecretRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.tenants.Secret.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.tenants.RegenerateSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.tenants.Secret)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.tenants.Secret>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.tenants.TenantApiClient.prototype.regenerateSecret =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/RegenerateSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_RegenerateSecret,
      callback);
};


/**
 * @param {!proto.smartcore.bos.tenants.RegenerateSecretRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.tenants.Secret>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.tenants.TenantApiPromiseClient.prototype.regenerateSecret =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.tenants.TenantApi/RegenerateSecret',
      request,
      metadata || {},
      methodDescriptor_TenantApi_RegenerateSecret);
};


module.exports = proto.smartcore.bos.tenants;

