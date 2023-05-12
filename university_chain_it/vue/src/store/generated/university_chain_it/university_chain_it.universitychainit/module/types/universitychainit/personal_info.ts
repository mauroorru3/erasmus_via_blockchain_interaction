/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface PersonalInfo {
  gender: string;
  dateOfBirth: string;
  primaryNationality: string;
  countryOfBirth: string;
  provinceOfBirth: string;
  townOfBirth: string;
  taxCode: string;
}

const basePersonalInfo: object = {
  gender: "",
  dateOfBirth: "",
  primaryNationality: "",
  countryOfBirth: "",
  provinceOfBirth: "",
  townOfBirth: "",
  taxCode: "",
};

export const PersonalInfo = {
  encode(message: PersonalInfo, writer: Writer = Writer.create()): Writer {
    if (message.gender !== "") {
      writer.uint32(10).string(message.gender);
    }
    if (message.dateOfBirth !== "") {
      writer.uint32(18).string(message.dateOfBirth);
    }
    if (message.primaryNationality !== "") {
      writer.uint32(26).string(message.primaryNationality);
    }
    if (message.countryOfBirth !== "") {
      writer.uint32(34).string(message.countryOfBirth);
    }
    if (message.provinceOfBirth !== "") {
      writer.uint32(42).string(message.provinceOfBirth);
    }
    if (message.townOfBirth !== "") {
      writer.uint32(50).string(message.townOfBirth);
    }
    if (message.taxCode !== "") {
      writer.uint32(58).string(message.taxCode);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PersonalInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePersonalInfo } as PersonalInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gender = reader.string();
          break;
        case 2:
          message.dateOfBirth = reader.string();
          break;
        case 3:
          message.primaryNationality = reader.string();
          break;
        case 4:
          message.countryOfBirth = reader.string();
          break;
        case 5:
          message.provinceOfBirth = reader.string();
          break;
        case 6:
          message.townOfBirth = reader.string();
          break;
        case 7:
          message.taxCode = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PersonalInfo {
    const message = { ...basePersonalInfo } as PersonalInfo;
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = String(object.gender);
    } else {
      message.gender = "";
    }
    if (object.dateOfBirth !== undefined && object.dateOfBirth !== null) {
      message.dateOfBirth = String(object.dateOfBirth);
    } else {
      message.dateOfBirth = "";
    }
    if (
      object.primaryNationality !== undefined &&
      object.primaryNationality !== null
    ) {
      message.primaryNationality = String(object.primaryNationality);
    } else {
      message.primaryNationality = "";
    }
    if (object.countryOfBirth !== undefined && object.countryOfBirth !== null) {
      message.countryOfBirth = String(object.countryOfBirth);
    } else {
      message.countryOfBirth = "";
    }
    if (
      object.provinceOfBirth !== undefined &&
      object.provinceOfBirth !== null
    ) {
      message.provinceOfBirth = String(object.provinceOfBirth);
    } else {
      message.provinceOfBirth = "";
    }
    if (object.townOfBirth !== undefined && object.townOfBirth !== null) {
      message.townOfBirth = String(object.townOfBirth);
    } else {
      message.townOfBirth = "";
    }
    if (object.taxCode !== undefined && object.taxCode !== null) {
      message.taxCode = String(object.taxCode);
    } else {
      message.taxCode = "";
    }
    return message;
  },

  toJSON(message: PersonalInfo): unknown {
    const obj: any = {};
    message.gender !== undefined && (obj.gender = message.gender);
    message.dateOfBirth !== undefined &&
      (obj.dateOfBirth = message.dateOfBirth);
    message.primaryNationality !== undefined &&
      (obj.primaryNationality = message.primaryNationality);
    message.countryOfBirth !== undefined &&
      (obj.countryOfBirth = message.countryOfBirth);
    message.provinceOfBirth !== undefined &&
      (obj.provinceOfBirth = message.provinceOfBirth);
    message.townOfBirth !== undefined &&
      (obj.townOfBirth = message.townOfBirth);
    message.taxCode !== undefined && (obj.taxCode = message.taxCode);
    return obj;
  },

  fromPartial(object: DeepPartial<PersonalInfo>): PersonalInfo {
    const message = { ...basePersonalInfo } as PersonalInfo;
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = object.gender;
    } else {
      message.gender = "";
    }
    if (object.dateOfBirth !== undefined && object.dateOfBirth !== null) {
      message.dateOfBirth = object.dateOfBirth;
    } else {
      message.dateOfBirth = "";
    }
    if (
      object.primaryNationality !== undefined &&
      object.primaryNationality !== null
    ) {
      message.primaryNationality = object.primaryNationality;
    } else {
      message.primaryNationality = "";
    }
    if (object.countryOfBirth !== undefined && object.countryOfBirth !== null) {
      message.countryOfBirth = object.countryOfBirth;
    } else {
      message.countryOfBirth = "";
    }
    if (
      object.provinceOfBirth !== undefined &&
      object.provinceOfBirth !== null
    ) {
      message.provinceOfBirth = object.provinceOfBirth;
    } else {
      message.provinceOfBirth = "";
    }
    if (object.townOfBirth !== undefined && object.townOfBirth !== null) {
      message.townOfBirth = object.townOfBirth;
    } else {
      message.townOfBirth = "";
    }
    if (object.taxCode !== undefined && object.taxCode !== null) {
      message.taxCode = object.taxCode;
    } else {
      message.taxCode = "";
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
