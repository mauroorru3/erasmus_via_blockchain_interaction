/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface UniversityInfoUniroma1 {
  nextStudentId: number;
  secretariatKey: string;
  universityKey: string;
  caiKey: string;
  fifoHeadErasmus: string;
  fifoTailErasmus: string;
  deadlineTaxes: string;
  deadlineErasmus: string;
}

const baseUniversityInfoUniroma1: object = {
  nextStudentId: 0,
  secretariatKey: "",
  universityKey: "",
  caiKey: "",
  fifoHeadErasmus: "",
  fifoTailErasmus: "",
  deadlineTaxes: "",
  deadlineErasmus: "",
};

export const UniversityInfoUniroma1 = {
  encode(
    message: UniversityInfoUniroma1,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nextStudentId !== 0) {
      writer.uint32(8).uint64(message.nextStudentId);
    }
    if (message.secretariatKey !== "") {
      writer.uint32(18).string(message.secretariatKey);
    }
    if (message.universityKey !== "") {
      writer.uint32(26).string(message.universityKey);
    }
    if (message.caiKey !== "") {
      writer.uint32(34).string(message.caiKey);
    }
    if (message.fifoHeadErasmus !== "") {
      writer.uint32(42).string(message.fifoHeadErasmus);
    }
    if (message.fifoTailErasmus !== "") {
      writer.uint32(50).string(message.fifoTailErasmus);
    }
    if (message.deadlineTaxes !== "") {
      writer.uint32(58).string(message.deadlineTaxes);
    }
    if (message.deadlineErasmus !== "") {
      writer.uint32(66).string(message.deadlineErasmus);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UniversityInfoUniroma1 {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUniversityInfoUniroma1 } as UniversityInfoUniroma1;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nextStudentId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.secretariatKey = reader.string();
          break;
        case 3:
          message.universityKey = reader.string();
          break;
        case 4:
          message.caiKey = reader.string();
          break;
        case 5:
          message.fifoHeadErasmus = reader.string();
          break;
        case 6:
          message.fifoTailErasmus = reader.string();
          break;
        case 7:
          message.deadlineTaxes = reader.string();
          break;
        case 8:
          message.deadlineErasmus = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UniversityInfoUniroma1 {
    const message = { ...baseUniversityInfoUniroma1 } as UniversityInfoUniroma1;
    if (object.nextStudentId !== undefined && object.nextStudentId !== null) {
      message.nextStudentId = Number(object.nextStudentId);
    } else {
      message.nextStudentId = 0;
    }
    if (object.secretariatKey !== undefined && object.secretariatKey !== null) {
      message.secretariatKey = String(object.secretariatKey);
    } else {
      message.secretariatKey = "";
    }
    if (object.universityKey !== undefined && object.universityKey !== null) {
      message.universityKey = String(object.universityKey);
    } else {
      message.universityKey = "";
    }
    if (object.caiKey !== undefined && object.caiKey !== null) {
      message.caiKey = String(object.caiKey);
    } else {
      message.caiKey = "";
    }
    if (
      object.fifoHeadErasmus !== undefined &&
      object.fifoHeadErasmus !== null
    ) {
      message.fifoHeadErasmus = String(object.fifoHeadErasmus);
    } else {
      message.fifoHeadErasmus = "";
    }
    if (
      object.fifoTailErasmus !== undefined &&
      object.fifoTailErasmus !== null
    ) {
      message.fifoTailErasmus = String(object.fifoTailErasmus);
    } else {
      message.fifoTailErasmus = "";
    }
    if (object.deadlineTaxes !== undefined && object.deadlineTaxes !== null) {
      message.deadlineTaxes = String(object.deadlineTaxes);
    } else {
      message.deadlineTaxes = "";
    }
    if (
      object.deadlineErasmus !== undefined &&
      object.deadlineErasmus !== null
    ) {
      message.deadlineErasmus = String(object.deadlineErasmus);
    } else {
      message.deadlineErasmus = "";
    }
    return message;
  },

  toJSON(message: UniversityInfoUniroma1): unknown {
    const obj: any = {};
    message.nextStudentId !== undefined &&
      (obj.nextStudentId = message.nextStudentId);
    message.secretariatKey !== undefined &&
      (obj.secretariatKey = message.secretariatKey);
    message.universityKey !== undefined &&
      (obj.universityKey = message.universityKey);
    message.caiKey !== undefined && (obj.caiKey = message.caiKey);
    message.fifoHeadErasmus !== undefined &&
      (obj.fifoHeadErasmus = message.fifoHeadErasmus);
    message.fifoTailErasmus !== undefined &&
      (obj.fifoTailErasmus = message.fifoTailErasmus);
    message.deadlineTaxes !== undefined &&
      (obj.deadlineTaxes = message.deadlineTaxes);
    message.deadlineErasmus !== undefined &&
      (obj.deadlineErasmus = message.deadlineErasmus);
    return obj;
  },

  fromPartial(
    object: DeepPartial<UniversityInfoUniroma1>
  ): UniversityInfoUniroma1 {
    const message = { ...baseUniversityInfoUniroma1 } as UniversityInfoUniroma1;
    if (object.nextStudentId !== undefined && object.nextStudentId !== null) {
      message.nextStudentId = object.nextStudentId;
    } else {
      message.nextStudentId = 0;
    }
    if (object.secretariatKey !== undefined && object.secretariatKey !== null) {
      message.secretariatKey = object.secretariatKey;
    } else {
      message.secretariatKey = "";
    }
    if (object.universityKey !== undefined && object.universityKey !== null) {
      message.universityKey = object.universityKey;
    } else {
      message.universityKey = "";
    }
    if (object.caiKey !== undefined && object.caiKey !== null) {
      message.caiKey = object.caiKey;
    } else {
      message.caiKey = "";
    }
    if (
      object.fifoHeadErasmus !== undefined &&
      object.fifoHeadErasmus !== null
    ) {
      message.fifoHeadErasmus = object.fifoHeadErasmus;
    } else {
      message.fifoHeadErasmus = "";
    }
    if (
      object.fifoTailErasmus !== undefined &&
      object.fifoTailErasmus !== null
    ) {
      message.fifoTailErasmus = object.fifoTailErasmus;
    } else {
      message.fifoTailErasmus = "";
    }
    if (object.deadlineTaxes !== undefined && object.deadlineTaxes !== null) {
      message.deadlineTaxes = object.deadlineTaxes;
    } else {
      message.deadlineTaxes = "";
    }
    if (
      object.deadlineErasmus !== undefined &&
      object.deadlineErasmus !== null
    ) {
      message.deadlineErasmus = object.deadlineErasmus;
    } else {
      message.deadlineErasmus = "";
    }
    return message;
  },
};

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
