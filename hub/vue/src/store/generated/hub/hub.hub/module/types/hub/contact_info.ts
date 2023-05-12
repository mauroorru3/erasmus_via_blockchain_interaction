/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface ContactInfo {
  contactAddress: string;
  email: string;
  mobilePhone: string;
}

const baseContactInfo: object = {
  contactAddress: "",
  email: "",
  mobilePhone: "",
};

export const ContactInfo = {
  encode(message: ContactInfo, writer: Writer = Writer.create()): Writer {
    if (message.contactAddress !== "") {
      writer.uint32(10).string(message.contactAddress);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    if (message.mobilePhone !== "") {
      writer.uint32(26).string(message.mobilePhone);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ContactInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseContactInfo } as ContactInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contactAddress = reader.string();
          break;
        case 2:
          message.email = reader.string();
          break;
        case 3:
          message.mobilePhone = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ContactInfo {
    const message = { ...baseContactInfo } as ContactInfo;
    if (object.contactAddress !== undefined && object.contactAddress !== null) {
      message.contactAddress = String(object.contactAddress);
    } else {
      message.contactAddress = "";
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = String(object.email);
    } else {
      message.email = "";
    }
    if (object.mobilePhone !== undefined && object.mobilePhone !== null) {
      message.mobilePhone = String(object.mobilePhone);
    } else {
      message.mobilePhone = "";
    }
    return message;
  },

  toJSON(message: ContactInfo): unknown {
    const obj: any = {};
    message.contactAddress !== undefined &&
      (obj.contactAddress = message.contactAddress);
    message.email !== undefined && (obj.email = message.email);
    message.mobilePhone !== undefined &&
      (obj.mobilePhone = message.mobilePhone);
    return obj;
  },

  fromPartial(object: DeepPartial<ContactInfo>): ContactInfo {
    const message = { ...baseContactInfo } as ContactInfo;
    if (object.contactAddress !== undefined && object.contactAddress !== null) {
      message.contactAddress = object.contactAddress;
    } else {
      message.contactAddress = "";
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = object.email;
    } else {
      message.email = "";
    }
    if (object.mobilePhone !== undefined && object.mobilePhone !== null) {
      message.mobilePhone = object.mobilePhone;
    } else {
      message.mobilePhone = "";
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
