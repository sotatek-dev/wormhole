/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import {
  ChainID,
  chainIDFromJSON,
  chainIDToJSON,
} from "../../publicrpc/v1/publicrpc";
import { Observable } from "rxjs";
import { BrowserHeaders } from "browser-headers";
import { share } from "rxjs/operators";

export const protobufPackage = "spy.v1";

/** A MessageFilter represents an exact match for an emitter. */
export interface EmitterFilter {
  /** Source chain */
  chainId: ChainID;
  /** Hex-encoded (without leading 0x) emitter address. */
  emitterAddress: string;
}

export interface FilterEntry {
  emitterFilter: EmitterFilter | undefined;
}

export interface SubscribeSignedVAARequest {
  /**
   * List of filters to apply to the stream (OR).
   * If empty, all messages are streamed.
   */
  filters: FilterEntry[];
}

export interface SubscribeSignedVAAResponse {
  /** Raw VAA bytes */
  vaaBytes: Uint8Array;
}

const baseEmitterFilter: object = { chainId: 0, emitterAddress: "" };

export const EmitterFilter = {
  encode(
    message: EmitterFilter,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.chainId !== 0) {
      writer.uint32(8).int32(message.chainId);
    }
    if (message.emitterAddress !== "") {
      writer.uint32(18).string(message.emitterAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EmitterFilter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEmitterFilter } as EmitterFilter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chainId = reader.int32() as any;
          break;
        case 2:
          message.emitterAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmitterFilter {
    const message = { ...baseEmitterFilter } as EmitterFilter;
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = chainIDFromJSON(object.chainId);
    } else {
      message.chainId = 0;
    }
    if (object.emitterAddress !== undefined && object.emitterAddress !== null) {
      message.emitterAddress = String(object.emitterAddress);
    } else {
      message.emitterAddress = "";
    }
    return message;
  },

  toJSON(message: EmitterFilter): unknown {
    const obj: any = {};
    message.chainId !== undefined &&
      (obj.chainId = chainIDToJSON(message.chainId));
    message.emitterAddress !== undefined &&
      (obj.emitterAddress = message.emitterAddress);
    return obj;
  },

  fromPartial(object: DeepPartial<EmitterFilter>): EmitterFilter {
    const message = { ...baseEmitterFilter } as EmitterFilter;
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = 0;
    }
    if (object.emitterAddress !== undefined && object.emitterAddress !== null) {
      message.emitterAddress = object.emitterAddress;
    } else {
      message.emitterAddress = "";
    }
    return message;
  },
};

const baseFilterEntry: object = {};

export const FilterEntry = {
  encode(
    message: FilterEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.emitterFilter !== undefined) {
      EmitterFilter.encode(
        message.emitterFilter,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FilterEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFilterEntry } as FilterEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.emitterFilter = EmitterFilter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FilterEntry {
    const message = { ...baseFilterEntry } as FilterEntry;
    if (object.emitterFilter !== undefined && object.emitterFilter !== null) {
      message.emitterFilter = EmitterFilter.fromJSON(object.emitterFilter);
    } else {
      message.emitterFilter = undefined;
    }
    return message;
  },

  toJSON(message: FilterEntry): unknown {
    const obj: any = {};
    message.emitterFilter !== undefined &&
      (obj.emitterFilter = message.emitterFilter
        ? EmitterFilter.toJSON(message.emitterFilter)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<FilterEntry>): FilterEntry {
    const message = { ...baseFilterEntry } as FilterEntry;
    if (object.emitterFilter !== undefined && object.emitterFilter !== null) {
      message.emitterFilter = EmitterFilter.fromPartial(object.emitterFilter);
    } else {
      message.emitterFilter = undefined;
    }
    return message;
  },
};

const baseSubscribeSignedVAARequest: object = {};

export const SubscribeSignedVAARequest = {
  encode(
    message: SubscribeSignedVAARequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.filters) {
      FilterEntry.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): SubscribeSignedVAARequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseSubscribeSignedVAARequest,
    } as SubscribeSignedVAARequest;
    message.filters = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.filters.push(FilterEntry.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubscribeSignedVAARequest {
    const message = {
      ...baseSubscribeSignedVAARequest,
    } as SubscribeSignedVAARequest;
    message.filters = [];
    if (object.filters !== undefined && object.filters !== null) {
      for (const e of object.filters) {
        message.filters.push(FilterEntry.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: SubscribeSignedVAARequest): unknown {
    const obj: any = {};
    if (message.filters) {
      obj.filters = message.filters.map((e) =>
        e ? FilterEntry.toJSON(e) : undefined
      );
    } else {
      obj.filters = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<SubscribeSignedVAARequest>
  ): SubscribeSignedVAARequest {
    const message = {
      ...baseSubscribeSignedVAARequest,
    } as SubscribeSignedVAARequest;
    message.filters = [];
    if (object.filters !== undefined && object.filters !== null) {
      for (const e of object.filters) {
        message.filters.push(FilterEntry.fromPartial(e));
      }
    }
    return message;
  },
};

const baseSubscribeSignedVAAResponse: object = {};

export const SubscribeSignedVAAResponse = {
  encode(
    message: SubscribeSignedVAAResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.vaaBytes.length !== 0) {
      writer.uint32(10).bytes(message.vaaBytes);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): SubscribeSignedVAAResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseSubscribeSignedVAAResponse,
    } as SubscribeSignedVAAResponse;
    message.vaaBytes = new Uint8Array();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vaaBytes = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubscribeSignedVAAResponse {
    const message = {
      ...baseSubscribeSignedVAAResponse,
    } as SubscribeSignedVAAResponse;
    message.vaaBytes = new Uint8Array();
    if (object.vaaBytes !== undefined && object.vaaBytes !== null) {
      message.vaaBytes = bytesFromBase64(object.vaaBytes);
    }
    return message;
  },

  toJSON(message: SubscribeSignedVAAResponse): unknown {
    const obj: any = {};
    message.vaaBytes !== undefined &&
      (obj.vaaBytes = base64FromBytes(
        message.vaaBytes !== undefined ? message.vaaBytes : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(
    object: DeepPartial<SubscribeSignedVAAResponse>
  ): SubscribeSignedVAAResponse {
    const message = {
      ...baseSubscribeSignedVAAResponse,
    } as SubscribeSignedVAAResponse;
    if (object.vaaBytes !== undefined && object.vaaBytes !== null) {
      message.vaaBytes = object.vaaBytes;
    } else {
      message.vaaBytes = new Uint8Array();
    }
    return message;
  },
};

/** SpyRPCService exposes a gossip introspection service, allowing sniffing of gossip messages. */
export interface SpyRPCService {
  /** SubscribeSignedVAA returns a stream of signed VAA messages received on the network. */
  SubscribeSignedVAA(
    request: DeepPartial<SubscribeSignedVAARequest>,
    metadata?: grpc.Metadata
  ): Observable<SubscribeSignedVAAResponse>;
}

export class SpyRPCServiceClientImpl implements SpyRPCService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SubscribeSignedVAA = this.SubscribeSignedVAA.bind(this);
  }

  SubscribeSignedVAA(
    request: DeepPartial<SubscribeSignedVAARequest>,
    metadata?: grpc.Metadata
  ): Observable<SubscribeSignedVAAResponse> {
    return this.rpc.invoke(
      SpyRPCServiceSubscribeSignedVAADesc,
      SubscribeSignedVAARequest.fromPartial(request),
      metadata
    );
  }
}

export const SpyRPCServiceDesc = {
  serviceName: "spy.v1.SpyRPCService",
};

export const SpyRPCServiceSubscribeSignedVAADesc: UnaryMethodDefinitionish = {
  methodName: "SubscribeSignedVAA",
  service: SpyRPCServiceDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return SubscribeSignedVAARequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SubscribeSignedVAAResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

interface UnaryMethodDefinitionishR
  extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any;
  responseStream: any;
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR;

interface Rpc {
  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined
  ): Promise<any>;
  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined
  ): Observable<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;
    streamingTransport?: grpc.TransportFactory;
    debug?: boolean;
    metadata?: grpc.Metadata;
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;
      streamingTransport?: grpc.TransportFactory;
      debug?: boolean;
      metadata?: grpc.Metadata;
    }
  ) {
    this.host = host;
    this.options = options;
  }

  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata =
      metadata && this.options.metadata
        ? new BrowserHeaders({
            ...this.options?.metadata.headersMap,
            ...metadata?.headersMap,
          })
        : metadata || this.options.metadata;
    return new Promise((resolve, reject) => {
      grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata,
        transport: this.options.transport,
        debug: this.options.debug,
        onEnd: function (response) {
          if (response.status === grpc.Code.OK) {
            resolve(response.message);
          } else {
            const err = new Error(response.statusMessage) as any;
            err.code = response.status;
            err.metadata = response.trailers;
            reject(err);
          }
        },
      });
    });
  }

  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined
  ): Observable<any> {
    // Status Response Codes (https://developers.google.com/maps-booking/reference/grpc-api/status_codes)
    const upStreamCodes = [2, 4, 8, 9, 10, 13, 14, 15];
    const DEFAULT_TIMEOUT_TIME: number = 3_000;
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata =
      metadata && this.options.metadata
        ? new BrowserHeaders({
            ...this.options?.metadata.headersMap,
            ...metadata?.headersMap,
          })
        : metadata || this.options.metadata;
    return new Observable((observer) => {
      const upStream = () => {
        const client = grpc.invoke(methodDesc, {
          host: this.host,
          request,
          transport: this.options.streamingTransport || this.options.transport,
          metadata: maybeCombinedMetadata,
          debug: this.options.debug,
          onMessage: (next) => observer.next(next),
          onEnd: (code: grpc.Code, message: string) => {
            if (code === 0) {
              observer.complete();
            } else if (upStreamCodes.includes(code)) {
              setTimeout(upStream, DEFAULT_TIMEOUT_TIME);
            } else {
              observer.error(new Error(`Error ${code} ${message}`));
            }
          },
        });
        observer.add(() => client.close());
      };
      upStream();
    }).pipe(share());
  }
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (const byte of arr) {
    bin.push(String.fromCharCode(byte));
  }
  return btoa(bin.join(""));
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
