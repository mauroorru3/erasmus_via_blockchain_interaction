/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface ErasmusExams {
  examName: string;
  examLabel: string;
  examDate: string;
  credits: number;
  marks: number;
  courseYear: number;
  status: boolean;
  attendanceYear: number;
  examType: number;
  courseOfStudy: string;
  homeUniversityId: string;
}

const baseErasmusExams: object = {
  examName: "",
  examLabel: "",
  examDate: "",
  credits: 0,
  marks: 0,
  courseYear: 0,
  status: false,
  attendanceYear: 0,
  examType: 0,
  courseOfStudy: "",
  homeUniversityId: "",
};

export const ErasmusExams = {
  encode(message: ErasmusExams, writer: Writer = Writer.create()): Writer {
    if (message.examName !== "") {
      writer.uint32(10).string(message.examName);
    }
    if (message.examLabel !== "") {
      writer.uint32(18).string(message.examLabel);
    }
    if (message.examDate !== "") {
      writer.uint32(26).string(message.examDate);
    }
    if (message.credits !== 0) {
      writer.uint32(32).uint64(message.credits);
    }
    if (message.marks !== 0) {
      writer.uint32(40).uint64(message.marks);
    }
    if (message.courseYear !== 0) {
      writer.uint32(48).uint64(message.courseYear);
    }
    if (message.status === true) {
      writer.uint32(56).bool(message.status);
    }
    if (message.attendanceYear !== 0) {
      writer.uint32(64).uint64(message.attendanceYear);
    }
    if (message.examType !== 0) {
      writer.uint32(72).uint64(message.examType);
    }
    if (message.courseOfStudy !== "") {
      writer.uint32(82).string(message.courseOfStudy);
    }
    if (message.homeUniversityId !== "") {
      writer.uint32(90).string(message.homeUniversityId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ErasmusExams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseErasmusExams } as ErasmusExams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.examName = reader.string();
          break;
        case 2:
          message.examLabel = reader.string();
          break;
        case 3:
          message.examDate = reader.string();
          break;
        case 4:
          message.credits = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.marks = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.courseYear = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.status = reader.bool();
          break;
        case 8:
          message.attendanceYear = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.examType = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.courseOfStudy = reader.string();
          break;
        case 11:
          message.homeUniversityId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusExams {
    const message = { ...baseErasmusExams } as ErasmusExams;
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = String(object.examName);
    } else {
      message.examName = "";
    }
    if (object.examLabel !== undefined && object.examLabel !== null) {
      message.examLabel = String(object.examLabel);
    } else {
      message.examLabel = "";
    }
    if (object.examDate !== undefined && object.examDate !== null) {
      message.examDate = String(object.examDate);
    } else {
      message.examDate = "";
    }
    if (object.credits !== undefined && object.credits !== null) {
      message.credits = Number(object.credits);
    } else {
      message.credits = 0;
    }
    if (object.marks !== undefined && object.marks !== null) {
      message.marks = Number(object.marks);
    } else {
      message.marks = 0;
    }
    if (object.courseYear !== undefined && object.courseYear !== null) {
      message.courseYear = Number(object.courseYear);
    } else {
      message.courseYear = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = Boolean(object.status);
    } else {
      message.status = false;
    }
    if (object.attendanceYear !== undefined && object.attendanceYear !== null) {
      message.attendanceYear = Number(object.attendanceYear);
    } else {
      message.attendanceYear = 0;
    }
    if (object.examType !== undefined && object.examType !== null) {
      message.examType = Number(object.examType);
    } else {
      message.examType = 0;
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = String(object.courseOfStudy);
    } else {
      message.courseOfStudy = "";
    }
    if (
      object.homeUniversityId !== undefined &&
      object.homeUniversityId !== null
    ) {
      message.homeUniversityId = String(object.homeUniversityId);
    } else {
      message.homeUniversityId = "";
    }
    return message;
  },

  toJSON(message: ErasmusExams): unknown {
    const obj: any = {};
    message.examName !== undefined && (obj.examName = message.examName);
    message.examLabel !== undefined && (obj.examLabel = message.examLabel);
    message.examDate !== undefined && (obj.examDate = message.examDate);
    message.credits !== undefined && (obj.credits = message.credits);
    message.marks !== undefined && (obj.marks = message.marks);
    message.courseYear !== undefined && (obj.courseYear = message.courseYear);
    message.status !== undefined && (obj.status = message.status);
    message.attendanceYear !== undefined &&
      (obj.attendanceYear = message.attendanceYear);
    message.examType !== undefined && (obj.examType = message.examType);
    message.courseOfStudy !== undefined &&
      (obj.courseOfStudy = message.courseOfStudy);
    message.homeUniversityId !== undefined &&
      (obj.homeUniversityId = message.homeUniversityId);
    return obj;
  },

  fromPartial(object: DeepPartial<ErasmusExams>): ErasmusExams {
    const message = { ...baseErasmusExams } as ErasmusExams;
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = object.examName;
    } else {
      message.examName = "";
    }
    if (object.examLabel !== undefined && object.examLabel !== null) {
      message.examLabel = object.examLabel;
    } else {
      message.examLabel = "";
    }
    if (object.examDate !== undefined && object.examDate !== null) {
      message.examDate = object.examDate;
    } else {
      message.examDate = "";
    }
    if (object.credits !== undefined && object.credits !== null) {
      message.credits = object.credits;
    } else {
      message.credits = 0;
    }
    if (object.marks !== undefined && object.marks !== null) {
      message.marks = object.marks;
    } else {
      message.marks = 0;
    }
    if (object.courseYear !== undefined && object.courseYear !== null) {
      message.courseYear = object.courseYear;
    } else {
      message.courseYear = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = false;
    }
    if (object.attendanceYear !== undefined && object.attendanceYear !== null) {
      message.attendanceYear = object.attendanceYear;
    } else {
      message.attendanceYear = 0;
    }
    if (object.examType !== undefined && object.examType !== null) {
      message.examType = object.examType;
    } else {
      message.examType = 0;
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = object.courseOfStudy;
    } else {
      message.courseOfStudy = "";
    }
    if (
      object.homeUniversityId !== undefined &&
      object.homeUniversityId !== null
    ) {
      message.homeUniversityId = object.homeUniversityId;
    } else {
      message.homeUniversityId = "";
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
