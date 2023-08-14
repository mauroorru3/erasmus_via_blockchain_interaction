/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface Universities {
  universityName: string;
  universitiesKey: string;
  universitiesCountry: string;
  chainName: string;
  port: string;
  channelID: string;
}

const baseUniversities: object = {
  universityName: "",
  universitiesKey: "",
  universitiesCountry: "",
  chainName: "",
  port: "",
  channelID: "",
};

export const Universities = {
  encode(message: Universities, writer: Writer = Writer.create()): Writer {
    if (message.universityName !== "") {
      writer.uint32(10).string(message.universityName);
    }
    if (message.universitiesKey !== "") {
      writer.uint32(18).string(message.universitiesKey);
    }
    if (message.universitiesCountry !== "") {
      writer.uint32(26).string(message.universitiesCountry);
    }
    if (message.chainName !== "") {
      writer.uint32(34).string(message.chainName);
    }
    if (message.port !== "") {
      writer.uint32(42).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(50).string(message.channelID);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Universities {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUniversities } as Universities;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityName = reader.string();
          break;
        case 2:
          message.universitiesKey = reader.string();
          break;
        case 3:
          message.universitiesCountry = reader.string();
          break;
        case 4:
          message.chainName = reader.string();
          break;
        case 5:
          message.port = reader.string();
          break;
        case 6:
          message.channelID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Universities {
    const message = { ...baseUniversities } as Universities;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = String(object.universityName);
    } else {
      message.universityName = "";
    }
    if (
      object.universitiesKey !== undefined &&
      object.universitiesKey !== null
    ) {
      message.universitiesKey = String(object.universitiesKey);
    } else {
      message.universitiesKey = "";
    }
    if (
      object.universitiesCountry !== undefined &&
      object.universitiesCountry !== null
    ) {
      message.universitiesCountry = String(object.universitiesCountry);
    } else {
      message.universitiesCountry = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = String(object.chainName);
    } else {
      message.chainName = "";
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
    return message;
  },

  toJSON(message: Universities): unknown {
    const obj: any = {};
    message.universityName !== undefined &&
      (obj.universityName = message.universityName);
    message.universitiesKey !== undefined &&
      (obj.universitiesKey = message.universitiesKey);
    message.universitiesCountry !== undefined &&
      (obj.universitiesCountry = message.universitiesCountry);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    return obj;
  },

  fromPartial(object: DeepPartial<Universities>): Universities {
    const message = { ...baseUniversities } as Universities;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = object.universityName;
    } else {
      message.universityName = "";
    }
    if (
      object.universitiesKey !== undefined &&
      object.universitiesKey !== null
    ) {
      message.universitiesKey = object.universitiesKey;
    } else {
      message.universitiesKey = "";
    }
    if (
      object.universitiesCountry !== undefined &&
      object.universitiesCountry !== null
    ) {
      message.universitiesCountry = object.universitiesCountry;
    } else {
      message.universitiesCountry = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = object.chainName;
    } else {
      message.chainName = "";
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
    return message;
  },
};

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
