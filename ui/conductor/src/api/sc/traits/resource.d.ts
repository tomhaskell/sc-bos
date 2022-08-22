import * as grpcWeb from "grpc-web";
import {Message} from "google-protobuf";
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";
import {ChangeType} from "@smart-core-os/sc-api-grpc-web/types/change_pb";

type Opt<T> = T | null | undefined;
type Msg<T> = Message | { toObject(includeInstance?: boolean): T }

export function closeResource(resource: RemoteResource<any> | null);

export function setValue<V>(resource: ResourceValue<V, any>, val: V);

export function setCollection<V>(resource: ResourceCollection<V, any>, change: CollectionChange<V, any>, idFunc: (T) => string)

export function setError(resource: RemoteResource<any>, err: Error);

export function pullResource<M extends Msg<V>, V>(logPrefix: string, resource: RemoteResource<M>, newStream: StreamFactory<M>);

export function trackAction<V, M extends Msg<V>>(logPrefix: string, tracker: ActionTracker<V>, action: Action<V, M>): Promise<V>

export function newActionTracker<T>(): ActionTracker<T>

export interface RemoteResource<M> {
  loading?: boolean;
  stream?: grpcWeb.ClientReadableStream<M>;
  streamError?: Error;
  updateTime?: Date;
}

export interface ResourceValue<V, M extends Msg<V>> extends RemoteResource<M> {
  value?: V;
}


export interface ResourceCollection<V, M extends Msg<V>> extends RemoteResource<M> {
  value?: { [id: string]: V };
}

export interface ResourceCallback<V> {
  data(val: V);

  error(e: Error);
}

export interface StreamFactory<M> {
  (endpoint: string): grpcWeb.ClientReadableStream<M>;
}

export interface CollectionChange<V, M extends Msg<V>> {
  getName(): string;

  getChangeTime(): Timestamp;

  getChangeType(): ChangeType;

  getOldValue(): M | undefined;

  getNewValue(): M | undefined;
}

export interface ActionTracker<V> {
  loading?: boolean;
  error?: Error;
  response?: V;
  duration?: number;
}

export interface Action<V, M extends Msg<V>> {
  (endpoint: string): M
}
