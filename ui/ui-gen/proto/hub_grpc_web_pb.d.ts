import * as grpcWeb from 'grpc-web';

import * as hub_pb from './hub_pb';


export class HubApiClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getHubNode(
    request: hub_pb.GetHubNodeRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hub_pb.HubNode) => void
  ): grpcWeb.ClientReadableStream<hub_pb.HubNode>;

  listHubNodes(
    request: hub_pb.ListHubNodesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hub_pb.ListHubNodesResponse) => void
  ): grpcWeb.ClientReadableStream<hub_pb.ListHubNodesResponse>;

  inspectHubNode(
    request: hub_pb.InspectHubNodeRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hub_pb.HubNodeInspection) => void
  ): grpcWeb.ClientReadableStream<hub_pb.HubNodeInspection>;

  enrollHubNode(
    request: hub_pb.EnrollHubNodeRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hub_pb.HubNode) => void
  ): grpcWeb.ClientReadableStream<hub_pb.HubNode>;

  renewHubNode(
    request: hub_pb.RenewHubNodeRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hub_pb.HubNode) => void
  ): grpcWeb.ClientReadableStream<hub_pb.HubNode>;

  testHubNode(
    request: hub_pb.TestHubNodeRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hub_pb.TestHubNodeResponse) => void
  ): grpcWeb.ClientReadableStream<hub_pb.TestHubNodeResponse>;

}

export class HubApiPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getHubNode(
    request: hub_pb.GetHubNodeRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hub_pb.HubNode>;

  listHubNodes(
    request: hub_pb.ListHubNodesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hub_pb.ListHubNodesResponse>;

  inspectHubNode(
    request: hub_pb.InspectHubNodeRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hub_pb.HubNodeInspection>;

  enrollHubNode(
    request: hub_pb.EnrollHubNodeRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hub_pb.HubNode>;

  renewHubNode(
    request: hub_pb.RenewHubNodeRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hub_pb.HubNode>;

  testHubNode(
    request: hub_pb.TestHubNodeRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hub_pb.TestHubNodeResponse>;

}
