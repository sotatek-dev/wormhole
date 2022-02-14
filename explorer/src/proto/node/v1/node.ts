/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import { BrowserHeaders } from "browser-headers";

export const protobufPackage = "node.v1";

export interface InjectGovernanceVAARequest {
  /** Index of the current guardian set. */
  currentSetIndex: number;
  /** List of governance VAA messages to inject. */
  messages: GovernanceMessage[];
}

export interface GovernanceMessage {
  /**
   * Sequence number. This is critical for replay protection - make sure the sequence number
   * is unique for every new manually injected governance VAA. Sequences are tracked
   * by emitter, and manually injected VAAs all use a single hardcoded emitter.
   *
   * We use random sequence numbers for the manual emitter.
   */
  sequence: string;
  /** Random nonce for disambiguation. Must be identical across all nodes. */
  nonce: number;
  guardianSet: GuardianSetUpdate | undefined;
  contractUpgrade: ContractUpgrade | undefined;
  bridgeRegisterChain: BridgeRegisterChain | undefined;
  bridgeContractUpgrade: BridgeUpgradeContract | undefined;
}

export interface InjectGovernanceVAAResponse {
  /** Canonical digests of the submitted VAAs. */
  digests: Uint8Array[];
}

/**
 * GuardianSet represents a new guardian set to be submitted to and signed by the node.
 * During the genesis procedure, this data structure will be assembled using off-chain collaborative tooling
 * like GitHub using a human-readable encoding, so readability is a concern.
 */
export interface GuardianSetUpdate {
  guardians: GuardianSetUpdate_Guardian[];
}

/** List of guardian set members. */
export interface GuardianSetUpdate_Guardian {
  /**
   * Guardian key pubkey. Stored as hex string with 0x prefix for human readability -
   * this is the canonical Ethereum representation.
   */
  pubkey: string;
  /** Optional descriptive name. Not stored on any chain, purely informational. */
  name: string;
}

/** GuardianKey specifies the on-disk format for a node's guardian key. */
export interface GuardianKey {
  /** data is the binary representation of the secp256k1 private key. */
  data: Uint8Array;
  /** Whether this key is deterministically generated and unsuitable for production mode. */
  unsafeDeterministicKey: boolean;
}

export interface BridgeRegisterChain {
  /** Module identifier of the token or NFT bridge (typically "TokenBridge" or "NFTBridge") */
  module: string;
  /** ID of the chain to be registered. */
  chainId: number;
  /** Hex-encoded emitter address to be registered (without leading 0x). */
  emitterAddress: string;
}

/** ContractUpgrade represents a Wormhole contract update to be submitted to and signed by the node. */
export interface ContractUpgrade {
  /** ID of the chain where the Wormhole contract should be updated (uint8). */
  chainId: number;
  /** Hex-encoded address (without leading 0x) address of the new program/contract. */
  newContract: string;
}

export interface BridgeUpgradeContract {
  /** Module identifier of the token or NFT bridge (typically "TokenBridge" or "NFTBridge"). */
  module: string;
  /** ID of the chain where the bridge contract should be updated (uint16). */
  targetChainId: number;
  /** Hex-encoded address (without leading 0x) of the new program/contract. */
  newContract: string;
}

export interface FindMissingMessagesRequest {
  /** Emitter chain ID to iterate. */
  emitterChain: number;
  /** Hex-encoded (without leading 0x) emitter address to iterate. */
  emitterAddress: string;
}

export interface FindMissingMessagesResponse {
  /** List of missing sequence numbers. */
  missingMessages: string[];
  /** Range processed */
  firstSequence: string;
  lastSequence: string;
}

const baseInjectGovernanceVAARequest: object = { currentSetIndex: 0 };

export const InjectGovernanceVAARequest = {
  encode(
    message: InjectGovernanceVAARequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.currentSetIndex !== 0) {
      writer.uint32(8).uint32(message.currentSetIndex);
    }
    for (const v of message.messages) {
      GovernanceMessage.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InjectGovernanceVAARequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseInjectGovernanceVAARequest,
    } as InjectGovernanceVAARequest;
    message.messages = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentSetIndex = reader.uint32();
          break;
        case 2:
          message.messages.push(
            GovernanceMessage.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InjectGovernanceVAARequest {
    const message = {
      ...baseInjectGovernanceVAARequest,
    } as InjectGovernanceVAARequest;
    message.messages = [];
    if (
      object.currentSetIndex !== undefined &&
      object.currentSetIndex !== null
    ) {
      message.currentSetIndex = Number(object.currentSetIndex);
    } else {
      message.currentSetIndex = 0;
    }
    if (object.messages !== undefined && object.messages !== null) {
      for (const e of object.messages) {
        message.messages.push(GovernanceMessage.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: InjectGovernanceVAARequest): unknown {
    const obj: any = {};
    message.currentSetIndex !== undefined &&
      (obj.currentSetIndex = message.currentSetIndex);
    if (message.messages) {
      obj.messages = message.messages.map((e) =>
        e ? GovernanceMessage.toJSON(e) : undefined
      );
    } else {
      obj.messages = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<InjectGovernanceVAARequest>
  ): InjectGovernanceVAARequest {
    const message = {
      ...baseInjectGovernanceVAARequest,
    } as InjectGovernanceVAARequest;
    message.messages = [];
    if (
      object.currentSetIndex !== undefined &&
      object.currentSetIndex !== null
    ) {
      message.currentSetIndex = object.currentSetIndex;
    } else {
      message.currentSetIndex = 0;
    }
    if (object.messages !== undefined && object.messages !== null) {
      for (const e of object.messages) {
        message.messages.push(GovernanceMessage.fromPartial(e));
      }
    }
    return message;
  },
};

const baseGovernanceMessage: object = { sequence: "0", nonce: 0 };

export const GovernanceMessage = {
  encode(
    message: GovernanceMessage,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.sequence !== "0") {
      writer.uint32(16).uint64(message.sequence);
    }
    if (message.nonce !== 0) {
      writer.uint32(24).uint32(message.nonce);
    }
    if (message.guardianSet !== undefined) {
      GuardianSetUpdate.encode(
        message.guardianSet,
        writer.uint32(82).fork()
      ).ldelim();
    }
    if (message.contractUpgrade !== undefined) {
      ContractUpgrade.encode(
        message.contractUpgrade,
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.bridgeRegisterChain !== undefined) {
      BridgeRegisterChain.encode(
        message.bridgeRegisterChain,
        writer.uint32(98).fork()
      ).ldelim();
    }
    if (message.bridgeContractUpgrade !== undefined) {
      BridgeUpgradeContract.encode(
        message.bridgeContractUpgrade,
        writer.uint32(106).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GovernanceMessage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGovernanceMessage } as GovernanceMessage;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.sequence = longToString(reader.uint64() as Long);
          break;
        case 3:
          message.nonce = reader.uint32();
          break;
        case 10:
          message.guardianSet = GuardianSetUpdate.decode(
            reader,
            reader.uint32()
          );
          break;
        case 11:
          message.contractUpgrade = ContractUpgrade.decode(
            reader,
            reader.uint32()
          );
          break;
        case 12:
          message.bridgeRegisterChain = BridgeRegisterChain.decode(
            reader,
            reader.uint32()
          );
          break;
        case 13:
          message.bridgeContractUpgrade = BridgeUpgradeContract.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GovernanceMessage {
    const message = { ...baseGovernanceMessage } as GovernanceMessage;
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = String(object.sequence);
    } else {
      message.sequence = "0";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = Number(object.nonce);
    } else {
      message.nonce = 0;
    }
    if (object.guardianSet !== undefined && object.guardianSet !== null) {
      message.guardianSet = GuardianSetUpdate.fromJSON(object.guardianSet);
    } else {
      message.guardianSet = undefined;
    }
    if (
      object.contractUpgrade !== undefined &&
      object.contractUpgrade !== null
    ) {
      message.contractUpgrade = ContractUpgrade.fromJSON(
        object.contractUpgrade
      );
    } else {
      message.contractUpgrade = undefined;
    }
    if (
      object.bridgeRegisterChain !== undefined &&
      object.bridgeRegisterChain !== null
    ) {
      message.bridgeRegisterChain = BridgeRegisterChain.fromJSON(
        object.bridgeRegisterChain
      );
    } else {
      message.bridgeRegisterChain = undefined;
    }
    if (
      object.bridgeContractUpgrade !== undefined &&
      object.bridgeContractUpgrade !== null
    ) {
      message.bridgeContractUpgrade = BridgeUpgradeContract.fromJSON(
        object.bridgeContractUpgrade
      );
    } else {
      message.bridgeContractUpgrade = undefined;
    }
    return message;
  },

  toJSON(message: GovernanceMessage): unknown {
    const obj: any = {};
    message.sequence !== undefined && (obj.sequence = message.sequence);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    message.guardianSet !== undefined &&
      (obj.guardianSet = message.guardianSet
        ? GuardianSetUpdate.toJSON(message.guardianSet)
        : undefined);
    message.contractUpgrade !== undefined &&
      (obj.contractUpgrade = message.contractUpgrade
        ? ContractUpgrade.toJSON(message.contractUpgrade)
        : undefined);
    message.bridgeRegisterChain !== undefined &&
      (obj.bridgeRegisterChain = message.bridgeRegisterChain
        ? BridgeRegisterChain.toJSON(message.bridgeRegisterChain)
        : undefined);
    message.bridgeContractUpgrade !== undefined &&
      (obj.bridgeContractUpgrade = message.bridgeContractUpgrade
        ? BridgeUpgradeContract.toJSON(message.bridgeContractUpgrade)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GovernanceMessage>): GovernanceMessage {
    const message = { ...baseGovernanceMessage } as GovernanceMessage;
    if (object.sequence !== undefined && object.sequence !== null) {
      message.sequence = object.sequence;
    } else {
      message.sequence = "0";
    }
    if (object.nonce !== undefined && object.nonce !== null) {
      message.nonce = object.nonce;
    } else {
      message.nonce = 0;
    }
    if (object.guardianSet !== undefined && object.guardianSet !== null) {
      message.guardianSet = GuardianSetUpdate.fromPartial(object.guardianSet);
    } else {
      message.guardianSet = undefined;
    }
    if (
      object.contractUpgrade !== undefined &&
      object.contractUpgrade !== null
    ) {
      message.contractUpgrade = ContractUpgrade.fromPartial(
        object.contractUpgrade
      );
    } else {
      message.contractUpgrade = undefined;
    }
    if (
      object.bridgeRegisterChain !== undefined &&
      object.bridgeRegisterChain !== null
    ) {
      message.bridgeRegisterChain = BridgeRegisterChain.fromPartial(
        object.bridgeRegisterChain
      );
    } else {
      message.bridgeRegisterChain = undefined;
    }
    if (
      object.bridgeContractUpgrade !== undefined &&
      object.bridgeContractUpgrade !== null
    ) {
      message.bridgeContractUpgrade = BridgeUpgradeContract.fromPartial(
        object.bridgeContractUpgrade
      );
    } else {
      message.bridgeContractUpgrade = undefined;
    }
    return message;
  },
};

const baseInjectGovernanceVAAResponse: object = {};

export const InjectGovernanceVAAResponse = {
  encode(
    message: InjectGovernanceVAAResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.digests) {
      writer.uint32(10).bytes(v!);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InjectGovernanceVAAResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseInjectGovernanceVAAResponse,
    } as InjectGovernanceVAAResponse;
    message.digests = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.digests.push(reader.bytes());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InjectGovernanceVAAResponse {
    const message = {
      ...baseInjectGovernanceVAAResponse,
    } as InjectGovernanceVAAResponse;
    message.digests = [];
    if (object.digests !== undefined && object.digests !== null) {
      for (const e of object.digests) {
        message.digests.push(bytesFromBase64(e));
      }
    }
    return message;
  },

  toJSON(message: InjectGovernanceVAAResponse): unknown {
    const obj: any = {};
    if (message.digests) {
      obj.digests = message.digests.map((e) =>
        base64FromBytes(e !== undefined ? e : new Uint8Array())
      );
    } else {
      obj.digests = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<InjectGovernanceVAAResponse>
  ): InjectGovernanceVAAResponse {
    const message = {
      ...baseInjectGovernanceVAAResponse,
    } as InjectGovernanceVAAResponse;
    message.digests = [];
    if (object.digests !== undefined && object.digests !== null) {
      for (const e of object.digests) {
        message.digests.push(e);
      }
    }
    return message;
  },
};

const baseGuardianSetUpdate: object = {};

export const GuardianSetUpdate = {
  encode(
    message: GuardianSetUpdate,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.guardians) {
      GuardianSetUpdate_Guardian.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GuardianSetUpdate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGuardianSetUpdate } as GuardianSetUpdate;
    message.guardians = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 3:
          message.guardians.push(
            GuardianSetUpdate_Guardian.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GuardianSetUpdate {
    const message = { ...baseGuardianSetUpdate } as GuardianSetUpdate;
    message.guardians = [];
    if (object.guardians !== undefined && object.guardians !== null) {
      for (const e of object.guardians) {
        message.guardians.push(GuardianSetUpdate_Guardian.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GuardianSetUpdate): unknown {
    const obj: any = {};
    if (message.guardians) {
      obj.guardians = message.guardians.map((e) =>
        e ? GuardianSetUpdate_Guardian.toJSON(e) : undefined
      );
    } else {
      obj.guardians = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GuardianSetUpdate>): GuardianSetUpdate {
    const message = { ...baseGuardianSetUpdate } as GuardianSetUpdate;
    message.guardians = [];
    if (object.guardians !== undefined && object.guardians !== null) {
      for (const e of object.guardians) {
        message.guardians.push(GuardianSetUpdate_Guardian.fromPartial(e));
      }
    }
    return message;
  },
};

const baseGuardianSetUpdate_Guardian: object = { pubkey: "", name: "" };

export const GuardianSetUpdate_Guardian = {
  encode(
    message: GuardianSetUpdate_Guardian,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.pubkey !== "") {
      writer.uint32(10).string(message.pubkey);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): GuardianSetUpdate_Guardian {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseGuardianSetUpdate_Guardian,
    } as GuardianSetUpdate_Guardian;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pubkey = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GuardianSetUpdate_Guardian {
    const message = {
      ...baseGuardianSetUpdate_Guardian,
    } as GuardianSetUpdate_Guardian;
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = String(object.pubkey);
    } else {
      message.pubkey = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: GuardianSetUpdate_Guardian): unknown {
    const obj: any = {};
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(
    object: DeepPartial<GuardianSetUpdate_Guardian>
  ): GuardianSetUpdate_Guardian {
    const message = {
      ...baseGuardianSetUpdate_Guardian,
    } as GuardianSetUpdate_Guardian;
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = object.pubkey;
    } else {
      message.pubkey = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseGuardianKey: object = { unsafeDeterministicKey: false };

export const GuardianKey = {
  encode(
    message: GuardianKey,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data.length !== 0) {
      writer.uint32(10).bytes(message.data);
    }
    if (message.unsafeDeterministicKey === true) {
      writer.uint32(16).bool(message.unsafeDeterministicKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GuardianKey {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGuardianKey } as GuardianKey;
    message.data = new Uint8Array();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = reader.bytes();
          break;
        case 2:
          message.unsafeDeterministicKey = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GuardianKey {
    const message = { ...baseGuardianKey } as GuardianKey;
    message.data = new Uint8Array();
    if (object.data !== undefined && object.data !== null) {
      message.data = bytesFromBase64(object.data);
    }
    if (
      object.unsafeDeterministicKey !== undefined &&
      object.unsafeDeterministicKey !== null
    ) {
      message.unsafeDeterministicKey = Boolean(object.unsafeDeterministicKey);
    } else {
      message.unsafeDeterministicKey = false;
    }
    return message;
  },

  toJSON(message: GuardianKey): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : new Uint8Array()
      ));
    message.unsafeDeterministicKey !== undefined &&
      (obj.unsafeDeterministicKey = message.unsafeDeterministicKey);
    return obj;
  },

  fromPartial(object: DeepPartial<GuardianKey>): GuardianKey {
    const message = { ...baseGuardianKey } as GuardianKey;
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = new Uint8Array();
    }
    if (
      object.unsafeDeterministicKey !== undefined &&
      object.unsafeDeterministicKey !== null
    ) {
      message.unsafeDeterministicKey = object.unsafeDeterministicKey;
    } else {
      message.unsafeDeterministicKey = false;
    }
    return message;
  },
};

const baseBridgeRegisterChain: object = {
  module: "",
  chainId: 0,
  emitterAddress: "",
};

export const BridgeRegisterChain = {
  encode(
    message: BridgeRegisterChain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.module !== "") {
      writer.uint32(10).string(message.module);
    }
    if (message.chainId !== 0) {
      writer.uint32(16).uint32(message.chainId);
    }
    if (message.emitterAddress !== "") {
      writer.uint32(26).string(message.emitterAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BridgeRegisterChain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBridgeRegisterChain } as BridgeRegisterChain;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.module = reader.string();
          break;
        case 2:
          message.chainId = reader.uint32();
          break;
        case 3:
          message.emitterAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BridgeRegisterChain {
    const message = { ...baseBridgeRegisterChain } as BridgeRegisterChain;
    if (object.module !== undefined && object.module !== null) {
      message.module = String(object.module);
    } else {
      message.module = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = Number(object.chainId);
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

  toJSON(message: BridgeRegisterChain): unknown {
    const obj: any = {};
    message.module !== undefined && (obj.module = message.module);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.emitterAddress !== undefined &&
      (obj.emitterAddress = message.emitterAddress);
    return obj;
  },

  fromPartial(object: DeepPartial<BridgeRegisterChain>): BridgeRegisterChain {
    const message = { ...baseBridgeRegisterChain } as BridgeRegisterChain;
    if (object.module !== undefined && object.module !== null) {
      message.module = object.module;
    } else {
      message.module = "";
    }
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

const baseContractUpgrade: object = { chainId: 0, newContract: "" };

export const ContractUpgrade = {
  encode(
    message: ContractUpgrade,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.chainId !== 0) {
      writer.uint32(8).uint32(message.chainId);
    }
    if (message.newContract !== "") {
      writer.uint32(18).string(message.newContract);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ContractUpgrade {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseContractUpgrade } as ContractUpgrade;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chainId = reader.uint32();
          break;
        case 2:
          message.newContract = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ContractUpgrade {
    const message = { ...baseContractUpgrade } as ContractUpgrade;
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = Number(object.chainId);
    } else {
      message.chainId = 0;
    }
    if (object.newContract !== undefined && object.newContract !== null) {
      message.newContract = String(object.newContract);
    } else {
      message.newContract = "";
    }
    return message;
  },

  toJSON(message: ContractUpgrade): unknown {
    const obj: any = {};
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.newContract !== undefined &&
      (obj.newContract = message.newContract);
    return obj;
  },

  fromPartial(object: DeepPartial<ContractUpgrade>): ContractUpgrade {
    const message = { ...baseContractUpgrade } as ContractUpgrade;
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = 0;
    }
    if (object.newContract !== undefined && object.newContract !== null) {
      message.newContract = object.newContract;
    } else {
      message.newContract = "";
    }
    return message;
  },
};

const baseBridgeUpgradeContract: object = {
  module: "",
  targetChainId: 0,
  newContract: "",
};

export const BridgeUpgradeContract = {
  encode(
    message: BridgeUpgradeContract,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.module !== "") {
      writer.uint32(10).string(message.module);
    }
    if (message.targetChainId !== 0) {
      writer.uint32(16).uint32(message.targetChainId);
    }
    if (message.newContract !== "") {
      writer.uint32(26).string(message.newContract);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): BridgeUpgradeContract {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBridgeUpgradeContract } as BridgeUpgradeContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.module = reader.string();
          break;
        case 2:
          message.targetChainId = reader.uint32();
          break;
        case 3:
          message.newContract = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BridgeUpgradeContract {
    const message = { ...baseBridgeUpgradeContract } as BridgeUpgradeContract;
    if (object.module !== undefined && object.module !== null) {
      message.module = String(object.module);
    } else {
      message.module = "";
    }
    if (object.targetChainId !== undefined && object.targetChainId !== null) {
      message.targetChainId = Number(object.targetChainId);
    } else {
      message.targetChainId = 0;
    }
    if (object.newContract !== undefined && object.newContract !== null) {
      message.newContract = String(object.newContract);
    } else {
      message.newContract = "";
    }
    return message;
  },

  toJSON(message: BridgeUpgradeContract): unknown {
    const obj: any = {};
    message.module !== undefined && (obj.module = message.module);
    message.targetChainId !== undefined &&
      (obj.targetChainId = message.targetChainId);
    message.newContract !== undefined &&
      (obj.newContract = message.newContract);
    return obj;
  },

  fromPartial(
    object: DeepPartial<BridgeUpgradeContract>
  ): BridgeUpgradeContract {
    const message = { ...baseBridgeUpgradeContract } as BridgeUpgradeContract;
    if (object.module !== undefined && object.module !== null) {
      message.module = object.module;
    } else {
      message.module = "";
    }
    if (object.targetChainId !== undefined && object.targetChainId !== null) {
      message.targetChainId = object.targetChainId;
    } else {
      message.targetChainId = 0;
    }
    if (object.newContract !== undefined && object.newContract !== null) {
      message.newContract = object.newContract;
    } else {
      message.newContract = "";
    }
    return message;
  },
};

const baseFindMissingMessagesRequest: object = {
  emitterChain: 0,
  emitterAddress: "",
};

export const FindMissingMessagesRequest = {
  encode(
    message: FindMissingMessagesRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.emitterChain !== 0) {
      writer.uint32(8).uint32(message.emitterChain);
    }
    if (message.emitterAddress !== "") {
      writer.uint32(18).string(message.emitterAddress);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): FindMissingMessagesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseFindMissingMessagesRequest,
    } as FindMissingMessagesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.emitterChain = reader.uint32();
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

  fromJSON(object: any): FindMissingMessagesRequest {
    const message = {
      ...baseFindMissingMessagesRequest,
    } as FindMissingMessagesRequest;
    if (object.emitterChain !== undefined && object.emitterChain !== null) {
      message.emitterChain = Number(object.emitterChain);
    } else {
      message.emitterChain = 0;
    }
    if (object.emitterAddress !== undefined && object.emitterAddress !== null) {
      message.emitterAddress = String(object.emitterAddress);
    } else {
      message.emitterAddress = "";
    }
    return message;
  },

  toJSON(message: FindMissingMessagesRequest): unknown {
    const obj: any = {};
    message.emitterChain !== undefined &&
      (obj.emitterChain = message.emitterChain);
    message.emitterAddress !== undefined &&
      (obj.emitterAddress = message.emitterAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<FindMissingMessagesRequest>
  ): FindMissingMessagesRequest {
    const message = {
      ...baseFindMissingMessagesRequest,
    } as FindMissingMessagesRequest;
    if (object.emitterChain !== undefined && object.emitterChain !== null) {
      message.emitterChain = object.emitterChain;
    } else {
      message.emitterChain = 0;
    }
    if (object.emitterAddress !== undefined && object.emitterAddress !== null) {
      message.emitterAddress = object.emitterAddress;
    } else {
      message.emitterAddress = "";
    }
    return message;
  },
};

const baseFindMissingMessagesResponse: object = {
  missingMessages: "",
  firstSequence: "0",
  lastSequence: "0",
};

export const FindMissingMessagesResponse = {
  encode(
    message: FindMissingMessagesResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.missingMessages) {
      writer.uint32(10).string(v!);
    }
    if (message.firstSequence !== "0") {
      writer.uint32(16).uint64(message.firstSequence);
    }
    if (message.lastSequence !== "0") {
      writer.uint32(24).uint64(message.lastSequence);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): FindMissingMessagesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseFindMissingMessagesResponse,
    } as FindMissingMessagesResponse;
    message.missingMessages = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.missingMessages.push(reader.string());
          break;
        case 2:
          message.firstSequence = longToString(reader.uint64() as Long);
          break;
        case 3:
          message.lastSequence = longToString(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FindMissingMessagesResponse {
    const message = {
      ...baseFindMissingMessagesResponse,
    } as FindMissingMessagesResponse;
    message.missingMessages = [];
    if (
      object.missingMessages !== undefined &&
      object.missingMessages !== null
    ) {
      for (const e of object.missingMessages) {
        message.missingMessages.push(String(e));
      }
    }
    if (object.firstSequence !== undefined && object.firstSequence !== null) {
      message.firstSequence = String(object.firstSequence);
    } else {
      message.firstSequence = "0";
    }
    if (object.lastSequence !== undefined && object.lastSequence !== null) {
      message.lastSequence = String(object.lastSequence);
    } else {
      message.lastSequence = "0";
    }
    return message;
  },

  toJSON(message: FindMissingMessagesResponse): unknown {
    const obj: any = {};
    if (message.missingMessages) {
      obj.missingMessages = message.missingMessages.map((e) => e);
    } else {
      obj.missingMessages = [];
    }
    message.firstSequence !== undefined &&
      (obj.firstSequence = message.firstSequence);
    message.lastSequence !== undefined &&
      (obj.lastSequence = message.lastSequence);
    return obj;
  },

  fromPartial(
    object: DeepPartial<FindMissingMessagesResponse>
  ): FindMissingMessagesResponse {
    const message = {
      ...baseFindMissingMessagesResponse,
    } as FindMissingMessagesResponse;
    message.missingMessages = [];
    if (
      object.missingMessages !== undefined &&
      object.missingMessages !== null
    ) {
      for (const e of object.missingMessages) {
        message.missingMessages.push(e);
      }
    }
    if (object.firstSequence !== undefined && object.firstSequence !== null) {
      message.firstSequence = object.firstSequence;
    } else {
      message.firstSequence = "0";
    }
    if (object.lastSequence !== undefined && object.lastSequence !== null) {
      message.lastSequence = object.lastSequence;
    } else {
      message.lastSequence = "0";
    }
    return message;
  },
};

/**
 * NodePrivilegedService exposes an administrative API. It runs on a UNIX socket and is authenticated
 * using Linux filesystem permissions.
 */
export interface NodePrivilegedService {
  /**
   * InjectGovernanceVAA injects a governance VAA into the guardian node.
   * The node will inject the VAA into the aggregator and sign/broadcast the VAA signature.
   *
   * A consensus majority of nodes on the network will have to inject the VAA within the
   * VAA timeout window for it to reach consensus.
   */
  InjectGovernanceVAA(
    request: DeepPartial<InjectGovernanceVAARequest>,
    metadata?: grpc.Metadata
  ): Promise<InjectGovernanceVAAResponse>;
  /**
   * FindMissingMessages will detect message sequence gaps in the local VAA store for a
   * specific emitter chain and address. Start and end slots are the lowest and highest
   * sequence numbers available in the local store, respectively.
   *
   * An error is returned if more than 1000 gaps are found.
   */
  FindMissingMessages(
    request: DeepPartial<FindMissingMessagesRequest>,
    metadata?: grpc.Metadata
  ): Promise<FindMissingMessagesResponse>;
}

export class NodePrivilegedServiceClientImpl implements NodePrivilegedService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.InjectGovernanceVAA = this.InjectGovernanceVAA.bind(this);
    this.FindMissingMessages = this.FindMissingMessages.bind(this);
  }

  InjectGovernanceVAA(
    request: DeepPartial<InjectGovernanceVAARequest>,
    metadata?: grpc.Metadata
  ): Promise<InjectGovernanceVAAResponse> {
    return this.rpc.unary(
      NodePrivilegedServiceInjectGovernanceVAADesc,
      InjectGovernanceVAARequest.fromPartial(request),
      metadata
    );
  }

  FindMissingMessages(
    request: DeepPartial<FindMissingMessagesRequest>,
    metadata?: grpc.Metadata
  ): Promise<FindMissingMessagesResponse> {
    return this.rpc.unary(
      NodePrivilegedServiceFindMissingMessagesDesc,
      FindMissingMessagesRequest.fromPartial(request),
      metadata
    );
  }
}

export const NodePrivilegedServiceDesc = {
  serviceName: "node.v1.NodePrivilegedService",
};

export const NodePrivilegedServiceInjectGovernanceVAADesc: UnaryMethodDefinitionish =
  {
    methodName: "InjectGovernanceVAA",
    service: NodePrivilegedServiceDesc,
    requestStream: false,
    responseStream: false,
    requestType: {
      serializeBinary() {
        return InjectGovernanceVAARequest.encode(this).finish();
      },
    } as any,
    responseType: {
      deserializeBinary(data: Uint8Array) {
        return {
          ...InjectGovernanceVAAResponse.decode(data),
          toObject() {
            return this;
          },
        };
      },
    } as any,
  };

export const NodePrivilegedServiceFindMissingMessagesDesc: UnaryMethodDefinitionish =
  {
    methodName: "FindMissingMessages",
    service: NodePrivilegedServiceDesc,
    requestStream: false,
    responseStream: false,
    requestType: {
      serializeBinary() {
        return FindMissingMessagesRequest.encode(this).finish();
      },
    } as any,
    responseType: {
      deserializeBinary(data: Uint8Array) {
        return {
          ...FindMissingMessagesResponse.decode(data),
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
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;

    debug?: boolean;
    metadata?: grpc.Metadata;
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;

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

function longToString(long: Long) {
  return long.toString();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
