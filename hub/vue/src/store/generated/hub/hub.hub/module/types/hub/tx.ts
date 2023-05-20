/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { ErasmusInfo } from "../hub/erasmus_info";

export const protobufPackage = "hub.hub";

export interface MsgSendErasmusStudent {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  index: string;
}

export interface MsgSendErasmusStudentResponse {
  status: number;
}

export interface MsgConfigureChain {
  creator: string;
}

export interface MsgConfigureChainResponse {
  status: number;
}

export interface MsgSendErasmusIndex {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  chain: string;
  index: string;
  foreignIndex: string;
}

export interface MsgSendErasmusIndexResponse {
  status: number;
}

export interface MsgSendEndErasmusPeriodRequest {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  startingUniversityName: string;
  destinationUniversityName: string;
  index: string;
  foreignIndex: string;
}

export interface MsgSendEndErasmusPeriodRequestResponse {
  status: number;
}

export interface MsgSendFinalErasmusData {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  erasmusData: ErasmusInfo | undefined;
  homeIndex: string;
}

export interface MsgSendFinalErasmusDataResponse {
  status: number;
}

const baseMsgSendErasmusStudent: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  index: "",
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
    if (message.index !== "") {
      writer.uint32(42).string(message.index);
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
          message.index = reader.string();
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
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
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
    message.index !== undefined && (obj.index = message.index);
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
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgSendErasmusStudentResponse: object = { status: 0 };

export const MsgSendErasmusStudentResponse = {
  encode(
    message: MsgSendErasmusStudentResponse,
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
  ): MsgSendErasmusStudentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
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

  fromJSON(object: any): MsgSendErasmusStudentResponse {
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgSendErasmusStudentResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendErasmusStudentResponse>
  ): MsgSendErasmusStudentResponse {
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
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

const baseMsgSendErasmusIndex: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  chain: "",
  index: "",
  foreignIndex: "",
};

export const MsgSendErasmusIndex = {
  encode(
    message: MsgSendErasmusIndex,
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
    if (message.chain !== "") {
      writer.uint32(42).string(message.chain);
    }
    if (message.index !== "") {
      writer.uint32(50).string(message.index);
    }
    if (message.foreignIndex !== "") {
      writer.uint32(58).string(message.foreignIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendErasmusIndex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendErasmusIndex } as MsgSendErasmusIndex;
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
          message.chain = reader.string();
          break;
        case 6:
          message.index = reader.string();
          break;
        case 7:
          message.foreignIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendErasmusIndex {
    const message = { ...baseMsgSendErasmusIndex } as MsgSendErasmusIndex;
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
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = String(object.foreignIndex);
    } else {
      message.foreignIndex = "";
    }
    return message;
  },

  toJSON(message: MsgSendErasmusIndex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.chain !== undefined && (obj.chain = message.chain);
    message.index !== undefined && (obj.index = message.index);
    message.foreignIndex !== undefined &&
      (obj.foreignIndex = message.foreignIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendErasmusIndex>): MsgSendErasmusIndex {
    const message = { ...baseMsgSendErasmusIndex } as MsgSendErasmusIndex;
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
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = object.foreignIndex;
    } else {
      message.foreignIndex = "";
    }
    return message;
  },
};

const baseMsgSendErasmusIndexResponse: object = { status: 0 };

export const MsgSendErasmusIndexResponse = {
  encode(
    message: MsgSendErasmusIndexResponse,
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
  ): MsgSendErasmusIndexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendErasmusIndexResponse,
    } as MsgSendErasmusIndexResponse;
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

  fromJSON(object: any): MsgSendErasmusIndexResponse {
    const message = {
      ...baseMsgSendErasmusIndexResponse,
    } as MsgSendErasmusIndexResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgSendErasmusIndexResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendErasmusIndexResponse>
  ): MsgSendErasmusIndexResponse {
    const message = {
      ...baseMsgSendErasmusIndexResponse,
    } as MsgSendErasmusIndexResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgSendEndErasmusPeriodRequest: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  startingUniversityName: "",
  destinationUniversityName: "",
  index: "",
  foreignIndex: "",
};

export const MsgSendEndErasmusPeriodRequest = {
  encode(
    message: MsgSendEndErasmusPeriodRequest,
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
    if (message.startingUniversityName !== "") {
      writer.uint32(42).string(message.startingUniversityName);
    }
    if (message.destinationUniversityName !== "") {
      writer.uint32(50).string(message.destinationUniversityName);
    }
    if (message.index !== "") {
      writer.uint32(58).string(message.index);
    }
    if (message.foreignIndex !== "") {
      writer.uint32(66).string(message.foreignIndex);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendEndErasmusPeriodRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendEndErasmusPeriodRequest,
    } as MsgSendEndErasmusPeriodRequest;
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
          message.startingUniversityName = reader.string();
          break;
        case 6:
          message.destinationUniversityName = reader.string();
          break;
        case 7:
          message.index = reader.string();
          break;
        case 8:
          message.foreignIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendEndErasmusPeriodRequest {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequest,
    } as MsgSendEndErasmusPeriodRequest;
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
    if (
      object.startingUniversityName !== undefined &&
      object.startingUniversityName !== null
    ) {
      message.startingUniversityName = String(object.startingUniversityName);
    } else {
      message.startingUniversityName = "";
    }
    if (
      object.destinationUniversityName !== undefined &&
      object.destinationUniversityName !== null
    ) {
      message.destinationUniversityName = String(
        object.destinationUniversityName
      );
    } else {
      message.destinationUniversityName = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = String(object.foreignIndex);
    } else {
      message.foreignIndex = "";
    }
    return message;
  },

  toJSON(message: MsgSendEndErasmusPeriodRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.startingUniversityName !== undefined &&
      (obj.startingUniversityName = message.startingUniversityName);
    message.destinationUniversityName !== undefined &&
      (obj.destinationUniversityName = message.destinationUniversityName);
    message.index !== undefined && (obj.index = message.index);
    message.foreignIndex !== undefined &&
      (obj.foreignIndex = message.foreignIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendEndErasmusPeriodRequest>
  ): MsgSendEndErasmusPeriodRequest {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequest,
    } as MsgSendEndErasmusPeriodRequest;
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
    if (
      object.startingUniversityName !== undefined &&
      object.startingUniversityName !== null
    ) {
      message.startingUniversityName = object.startingUniversityName;
    } else {
      message.startingUniversityName = "";
    }
    if (
      object.destinationUniversityName !== undefined &&
      object.destinationUniversityName !== null
    ) {
      message.destinationUniversityName = object.destinationUniversityName;
    } else {
      message.destinationUniversityName = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = object.foreignIndex;
    } else {
      message.foreignIndex = "";
    }
    return message;
  },
};

const baseMsgSendEndErasmusPeriodRequestResponse: object = { status: 0 };

export const MsgSendEndErasmusPeriodRequestResponse = {
  encode(
    message: MsgSendEndErasmusPeriodRequestResponse,
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
  ): MsgSendEndErasmusPeriodRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendEndErasmusPeriodRequestResponse,
    } as MsgSendEndErasmusPeriodRequestResponse;
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

  fromJSON(object: any): MsgSendEndErasmusPeriodRequestResponse {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequestResponse,
    } as MsgSendEndErasmusPeriodRequestResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgSendEndErasmusPeriodRequestResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendEndErasmusPeriodRequestResponse>
  ): MsgSendEndErasmusPeriodRequestResponse {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequestResponse,
    } as MsgSendEndErasmusPeriodRequestResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgSendFinalErasmusData: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  homeIndex: "",
};

export const MsgSendFinalErasmusData = {
  encode(
    message: MsgSendFinalErasmusData,
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
    if (message.erasmusData !== undefined) {
      ErasmusInfo.encode(
        message.erasmusData,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.homeIndex !== "") {
      writer.uint32(50).string(message.homeIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendFinalErasmusData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendFinalErasmusData,
    } as MsgSendFinalErasmusData;
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
          message.erasmusData = ErasmusInfo.decode(reader, reader.uint32());
          break;
        case 6:
          message.homeIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendFinalErasmusData {
    const message = {
      ...baseMsgSendFinalErasmusData,
    } as MsgSendFinalErasmusData;
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
    if (object.erasmusData !== undefined && object.erasmusData !== null) {
      message.erasmusData = ErasmusInfo.fromJSON(object.erasmusData);
    } else {
      message.erasmusData = undefined;
    }
    if (object.homeIndex !== undefined && object.homeIndex !== null) {
      message.homeIndex = String(object.homeIndex);
    } else {
      message.homeIndex = "";
    }
    return message;
  },

  toJSON(message: MsgSendFinalErasmusData): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.erasmusData !== undefined &&
      (obj.erasmusData = message.erasmusData
        ? ErasmusInfo.toJSON(message.erasmusData)
        : undefined);
    message.homeIndex !== undefined && (obj.homeIndex = message.homeIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendFinalErasmusData>
  ): MsgSendFinalErasmusData {
    const message = {
      ...baseMsgSendFinalErasmusData,
    } as MsgSendFinalErasmusData;
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
    if (object.erasmusData !== undefined && object.erasmusData !== null) {
      message.erasmusData = ErasmusInfo.fromPartial(object.erasmusData);
    } else {
      message.erasmusData = undefined;
    }
    if (object.homeIndex !== undefined && object.homeIndex !== null) {
      message.homeIndex = object.homeIndex;
    } else {
      message.homeIndex = "";
    }
    return message;
  },
};

const baseMsgSendFinalErasmusDataResponse: object = { status: 0 };

export const MsgSendFinalErasmusDataResponse = {
  encode(
    message: MsgSendFinalErasmusDataResponse,
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
  ): MsgSendFinalErasmusDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendFinalErasmusDataResponse,
    } as MsgSendFinalErasmusDataResponse;
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

  fromJSON(object: any): MsgSendFinalErasmusDataResponse {
    const message = {
      ...baseMsgSendFinalErasmusDataResponse,
    } as MsgSendFinalErasmusDataResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgSendFinalErasmusDataResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendFinalErasmusDataResponse>
  ): MsgSendFinalErasmusDataResponse {
    const message = {
      ...baseMsgSendFinalErasmusDataResponse,
    } as MsgSendFinalErasmusDataResponse;
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
  ConfigureChain(
    request: MsgConfigureChain
  ): Promise<MsgConfigureChainResponse>;
  SendErasmusIndex(
    request: MsgSendErasmusIndex
  ): Promise<MsgSendErasmusIndexResponse>;
  SendEndErasmusPeriodRequest(
    request: MsgSendEndErasmusPeriodRequest
  ): Promise<MsgSendEndErasmusPeriodRequestResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendFinalErasmusData(
    request: MsgSendFinalErasmusData
  ): Promise<MsgSendFinalErasmusDataResponse>;
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

  SendErasmusIndex(
    request: MsgSendErasmusIndex
  ): Promise<MsgSendErasmusIndexResponse> {
    const data = MsgSendErasmusIndex.encode(request).finish();
    const promise = this.rpc.request("hub.hub.Msg", "SendErasmusIndex", data);
    return promise.then((data) =>
      MsgSendErasmusIndexResponse.decode(new Reader(data))
    );
  }

  SendEndErasmusPeriodRequest(
    request: MsgSendEndErasmusPeriodRequest
  ): Promise<MsgSendEndErasmusPeriodRequestResponse> {
    const data = MsgSendEndErasmusPeriodRequest.encode(request).finish();
    const promise = this.rpc.request(
      "hub.hub.Msg",
      "SendEndErasmusPeriodRequest",
      data
    );
    return promise.then((data) =>
      MsgSendEndErasmusPeriodRequestResponse.decode(new Reader(data))
    );
  }

  SendFinalErasmusData(
    request: MsgSendFinalErasmusData
  ): Promise<MsgSendFinalErasmusDataResponse> {
    const data = MsgSendFinalErasmusData.encode(request).finish();
    const promise = this.rpc.request(
      "hub.hub.Msg",
      "SendFinalErasmusData",
      data
    );
    return promise.then((data) =>
      MsgSendFinalErasmusDataResponse.decode(new Reader(data))
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
