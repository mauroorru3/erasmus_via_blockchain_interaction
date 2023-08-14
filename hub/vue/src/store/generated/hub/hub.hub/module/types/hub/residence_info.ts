/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface ResidenceInfo {
  country: string;
  province: string;
  town: string;
  postCode: string;
  address: string;
  houseNumber: string;
  homePhone: string;
}

const baseResidenceInfo: object = {
  country: "",
  province: "",
  town: "",
  postCode: "",
  address: "",
  houseNumber: "",
  homePhone: "",
};

export const ResidenceInfo = {
  encode(message: ResidenceInfo, writer: Writer = Writer.create()): Writer {
    if (message.country !== "") {
      writer.uint32(10).string(message.country);
    }
    if (message.province !== "") {
      writer.uint32(18).string(message.province);
    }
    if (message.town !== "") {
      writer.uint32(26).string(message.town);
    }
    if (message.postCode !== "") {
      writer.uint32(34).string(message.postCode);
    }
    if (message.address !== "") {
      writer.uint32(42).string(message.address);
    }
    if (message.houseNumber !== "") {
      writer.uint32(50).string(message.houseNumber);
    }
    if (message.homePhone !== "") {
      writer.uint32(58).string(message.homePhone);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ResidenceInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseResidenceInfo } as ResidenceInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.country = reader.string();
          break;
        case 2:
          message.province = reader.string();
          break;
        case 3:
          message.town = reader.string();
          break;
        case 4:
          message.postCode = reader.string();
          break;
        case 5:
          message.address = reader.string();
          break;
        case 6:
          message.houseNumber = reader.string();
          break;
        case 7:
          message.homePhone = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ResidenceInfo {
    const message = { ...baseResidenceInfo } as ResidenceInfo;
    if (object.country !== undefined && object.country !== null) {
      message.country = String(object.country);
    } else {
      message.country = "";
    }
    if (object.province !== undefined && object.province !== null) {
      message.province = String(object.province);
    } else {
      message.province = "";
    }
    if (object.town !== undefined && object.town !== null) {
      message.town = String(object.town);
    } else {
      message.town = "";
    }
    if (object.postCode !== undefined && object.postCode !== null) {
      message.postCode = String(object.postCode);
    } else {
      message.postCode = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.houseNumber !== undefined && object.houseNumber !== null) {
      message.houseNumber = String(object.houseNumber);
    } else {
      message.houseNumber = "";
    }
    if (object.homePhone !== undefined && object.homePhone !== null) {
      message.homePhone = String(object.homePhone);
    } else {
      message.homePhone = "";
    }
    return message;
  },

  toJSON(message: ResidenceInfo): unknown {
    const obj: any = {};
    message.country !== undefined && (obj.country = message.country);
    message.province !== undefined && (obj.province = message.province);
    message.town !== undefined && (obj.town = message.town);
    message.postCode !== undefined && (obj.postCode = message.postCode);
    message.address !== undefined && (obj.address = message.address);
    message.houseNumber !== undefined &&
      (obj.houseNumber = message.houseNumber);
    message.homePhone !== undefined && (obj.homePhone = message.homePhone);
    return obj;
  },

  fromPartial(object: DeepPartial<ResidenceInfo>): ResidenceInfo {
    const message = { ...baseResidenceInfo } as ResidenceInfo;
    if (object.country !== undefined && object.country !== null) {
      message.country = object.country;
    } else {
      message.country = "";
    }
    if (object.province !== undefined && object.province !== null) {
      message.province = object.province;
    } else {
      message.province = "";
    }
    if (object.town !== undefined && object.town !== null) {
      message.town = object.town;
    } else {
      message.town = "";
    }
    if (object.postCode !== undefined && object.postCode !== null) {
      message.postCode = object.postCode;
    } else {
      message.postCode = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.houseNumber !== undefined && object.houseNumber !== null) {
      message.houseNumber = object.houseNumber;
    } else {
      message.houseNumber = "";
    }
    if (object.homePhone !== undefined && object.homePhone !== null) {
      message.homePhone = object.homePhone;
    } else {
      message.homePhone = "";
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
