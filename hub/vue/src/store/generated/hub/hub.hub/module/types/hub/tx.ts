/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { StoredStudent } from "../hub/stored_student";

export const protobufPackage = "hub.hub";

export interface MsgSendErasmusStudent {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  student: StoredStudent | undefined;
}

export interface MsgSendErasmusStudentResponse {}

export interface MsgConfigureChain {
  creator: string;
}

export interface MsgConfigureChainResponse {
  status: number;
}

const baseMsgSendErasmusStudent: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
};

export const MsgSendErasmusStudent = {
  encode(
    message: MsgSendErasmusStudent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.student !== undefined) {
      StoredStudent.encode(message.student, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendErasmusStudent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendErasmusStudent } as MsgSendErasmusStudent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.student = StoredStudent.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendErasmusStudent {
    const message = { ...baseMsgSendErasmusStudent } as MsgSendErasmusStudent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.student !== undefined && object.student !== null) {
      message.student = StoredStudent.fromJSON(object.student);
    } else {
      message.student = undefined;
    }
    return message;
  },

  toJSON(message: MsgSendErasmusStudent): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.student !== undefined &&
      (obj.student = message.student
        ? StoredStudent.toJSON(message.student)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendErasmusStudent>
  ): MsgSendErasmusStudent {
    const message = { ...baseMsgSendErasmusStudent } as MsgSendErasmusStudent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.student !== undefined && object.student !== null) {
      message.student = StoredStudent.fromPartial(object.student);
    } else {
      message.student = undefined;
    }
    return message;
  },
};

const baseMsgSendErasmusStudentResponse: object = {};

export const MsgSendErasmusStudentResponse = {
  encode(
    _: MsgSendErasmusStudentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendErasmusStudentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSendErasmusStudentResponse {
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    return message;
  },

  toJSON(_: MsgSendErasmusStudentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendErasmusStudentResponse>
  ): MsgSendErasmusStudentResponse {
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    return message;
  },
};

const baseMsgConfigureChain: object = { creator: "" };

export const MsgConfigureChain = {
  encode(message: MsgConfigureChain, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConfigureChain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgConfigureChain } as MsgConfigureChain;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfigureChain {
    const message = { ...baseMsgConfigureChain } as MsgConfigureChain;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgConfigureChain): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgConfigureChain>): MsgConfigureChain {
    const message = { ...baseMsgConfigureChain } as MsgConfigureChain;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgConfigureChainResponse: object = { status: 0 };

export const MsgConfigureChainResponse = {
  encode(
    message: MsgConfigureChainResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgConfigureChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgConfigureChainResponse,
    } as MsgConfigureChainResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfigureChainResponse {
    const message = {
      ...baseMsgConfigureChainResponse,
    } as MsgConfigureChainResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgConfigureChainResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgConfigureChainResponse>
  ): MsgConfigureChainResponse {
    const message = {
      ...baseMsgConfigureChainResponse,
    } as MsgConfigureChainResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SendErasmusStudent(
    request: MsgSendErasmusStudent
  ): Promise<MsgSendErasmusStudentResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ConfigureChain(
    request: MsgConfigureChain
  ): Promise<MsgConfigureChainResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SendErasmusStudent(
    request: MsgSendErasmusStudent
  ): Promise<MsgSendErasmusStudentResponse> {
    const data = MsgSendErasmusStudent.encode(request).finish();
    const promise = this.rpc.request("hub.hub.Msg", "SendErasmusStudent", data);
    return promise.then((data) =>
      MsgSendErasmusStudentResponse.decode(new Reader(data))
    );
  }

  ConfigureChain(
    request: MsgConfigureChain
  ): Promise<MsgConfigureChainResponse> {
    const data = MsgConfigureChain.encode(request).finish();
    const promise = this.rpc.request("hub.hub.Msg", "ConfigureChain", data);
    return promise.then((data) =>
      MsgConfigureChainResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
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

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
