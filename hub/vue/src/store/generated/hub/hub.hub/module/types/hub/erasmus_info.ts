/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface ErasmusInfo {
  erasmusStudent: string;
  numberTimes: number;
  numberMonths: number;
  totalExams: number;
  examsPassed: number;
  totalCredits: number;
  achievedCredits: number;
  career: string;
  previousStudentFifo: string;
  nextStudentFifo: string;
}

const baseErasmusInfo: object = {
  erasmusStudent: "",
  numberTimes: 0,
  numberMonths: 0,
  totalExams: 0,
  examsPassed: 0,
  totalCredits: 0,
  achievedCredits: 0,
  career: "",
  previousStudentFifo: "",
  nextStudentFifo: "",
};

export const ErasmusInfo = {
  encode(message: ErasmusInfo, writer: Writer = Writer.create()): Writer {
    if (message.erasmusStudent !== "") {
      writer.uint32(10).string(message.erasmusStudent);
    }
    if (message.numberTimes !== 0) {
      writer.uint32(16).uint64(message.numberTimes);
    }
    if (message.numberMonths !== 0) {
      writer.uint32(24).uint64(message.numberMonths);
    }
    if (message.totalExams !== 0) {
      writer.uint32(32).uint64(message.totalExams);
    }
    if (message.examsPassed !== 0) {
      writer.uint32(40).uint64(message.examsPassed);
    }
    if (message.totalCredits !== 0) {
      writer.uint32(48).uint64(message.totalCredits);
    }
    if (message.achievedCredits !== 0) {
      writer.uint32(56).uint64(message.achievedCredits);
    }
    if (message.career !== "") {
      writer.uint32(66).string(message.career);
    }
    if (message.previousStudentFifo !== "") {
      writer.uint32(74).string(message.previousStudentFifo);
    }
    if (message.nextStudentFifo !== "") {
      writer.uint32(82).string(message.nextStudentFifo);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ErasmusInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseErasmusInfo } as ErasmusInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erasmusStudent = reader.string();
          break;
        case 2:
          message.numberTimes = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.numberMonths = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.totalExams = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.examsPassed = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.totalCredits = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.achievedCredits = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.career = reader.string();
          break;
        case 9:
          message.previousStudentFifo = reader.string();
          break;
        case 10:
          message.nextStudentFifo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusInfo {
    const message = { ...baseErasmusInfo } as ErasmusInfo;
    if (object.erasmusStudent !== undefined && object.erasmusStudent !== null) {
      message.erasmusStudent = String(object.erasmusStudent);
    } else {
      message.erasmusStudent = "";
    }
    if (object.numberTimes !== undefined && object.numberTimes !== null) {
      message.numberTimes = Number(object.numberTimes);
    } else {
      message.numberTimes = 0;
    }
    if (object.numberMonths !== undefined && object.numberMonths !== null) {
      message.numberMonths = Number(object.numberMonths);
    } else {
      message.numberMonths = 0;
    }
    if (object.totalExams !== undefined && object.totalExams !== null) {
      message.totalExams = Number(object.totalExams);
    } else {
      message.totalExams = 0;
    }
    if (object.examsPassed !== undefined && object.examsPassed !== null) {
      message.examsPassed = Number(object.examsPassed);
    } else {
      message.examsPassed = 0;
    }
    if (object.totalCredits !== undefined && object.totalCredits !== null) {
      message.totalCredits = Number(object.totalCredits);
    } else {
      message.totalCredits = 0;
    }
    if (
      object.achievedCredits !== undefined &&
      object.achievedCredits !== null
    ) {
      message.achievedCredits = Number(object.achievedCredits);
    } else {
      message.achievedCredits = 0;
    }
    if (object.career !== undefined && object.career !== null) {
      message.career = String(object.career);
    } else {
      message.career = "";
    }
    if (
      object.previousStudentFifo !== undefined &&
      object.previousStudentFifo !== null
    ) {
      message.previousStudentFifo = String(object.previousStudentFifo);
    } else {
      message.previousStudentFifo = "";
    }
    if (
      object.nextStudentFifo !== undefined &&
      object.nextStudentFifo !== null
    ) {
      message.nextStudentFifo = String(object.nextStudentFifo);
    } else {
      message.nextStudentFifo = "";
    }
    return message;
  },

  toJSON(message: ErasmusInfo): unknown {
    const obj: any = {};
    message.erasmusStudent !== undefined &&
      (obj.erasmusStudent = message.erasmusStudent);
    message.numberTimes !== undefined &&
      (obj.numberTimes = message.numberTimes);
    message.numberMonths !== undefined &&
      (obj.numberMonths = message.numberMonths);
    message.totalExams !== undefined && (obj.totalExams = message.totalExams);
    message.examsPassed !== undefined &&
      (obj.examsPassed = message.examsPassed);
    message.totalCredits !== undefined &&
      (obj.totalCredits = message.totalCredits);
    message.achievedCredits !== undefined &&
      (obj.achievedCredits = message.achievedCredits);
    message.career !== undefined && (obj.career = message.career);
    message.previousStudentFifo !== undefined &&
      (obj.previousStudentFifo = message.previousStudentFifo);
    message.nextStudentFifo !== undefined &&
      (obj.nextStudentFifo = message.nextStudentFifo);
    return obj;
  },

  fromPartial(object: DeepPartial<ErasmusInfo>): ErasmusInfo {
    const message = { ...baseErasmusInfo } as ErasmusInfo;
    if (object.erasmusStudent !== undefined && object.erasmusStudent !== null) {
      message.erasmusStudent = object.erasmusStudent;
    } else {
      message.erasmusStudent = "";
    }
    if (object.numberTimes !== undefined && object.numberTimes !== null) {
      message.numberTimes = object.numberTimes;
    } else {
      message.numberTimes = 0;
    }
    if (object.numberMonths !== undefined && object.numberMonths !== null) {
      message.numberMonths = object.numberMonths;
    } else {
      message.numberMonths = 0;
    }
    if (object.totalExams !== undefined && object.totalExams !== null) {
      message.totalExams = object.totalExams;
    } else {
      message.totalExams = 0;
    }
    if (object.examsPassed !== undefined && object.examsPassed !== null) {
      message.examsPassed = object.examsPassed;
    } else {
      message.examsPassed = 0;
    }
    if (object.totalCredits !== undefined && object.totalCredits !== null) {
      message.totalCredits = object.totalCredits;
    } else {
      message.totalCredits = 0;
    }
    if (
      object.achievedCredits !== undefined &&
      object.achievedCredits !== null
    ) {
      message.achievedCredits = object.achievedCredits;
    } else {
      message.achievedCredits = 0;
    }
    if (object.career !== undefined && object.career !== null) {
      message.career = object.career;
    } else {
      message.career = "";
    }
    if (
      object.previousStudentFifo !== undefined &&
      object.previousStudentFifo !== null
    ) {
      message.previousStudentFifo = object.previousStudentFifo;
    } else {
      message.previousStudentFifo = "";
    }
    if (
      object.nextStudentFifo !== undefined &&
      object.nextStudentFifo !== null
    ) {
      message.nextStudentFifo = object.nextStudentFifo;
    } else {
      message.nextStudentFifo = "";
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
