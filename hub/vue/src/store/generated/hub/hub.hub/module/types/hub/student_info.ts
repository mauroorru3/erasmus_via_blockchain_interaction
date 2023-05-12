/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface StudentInfo {
  name: string;
  surname: string;
  courseType: number;
  courseOfStudy: string;
  status: number;
  currentYearOfStudy: number;
  outOfCourse: boolean;
  numberOfYearsOutOfCourse: number;
  studentKey: string;
}

const baseStudentInfo: object = {
  name: "",
  surname: "",
  courseType: 0,
  courseOfStudy: "",
  status: 0,
  currentYearOfStudy: 0,
  outOfCourse: false,
  numberOfYearsOutOfCourse: 0,
  studentKey: "",
};

export const StudentInfo = {
  encode(message: StudentInfo, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.surname !== "") {
      writer.uint32(18).string(message.surname);
    }
    if (message.courseType !== 0) {
      writer.uint32(24).uint64(message.courseType);
    }
    if (message.courseOfStudy !== "") {
      writer.uint32(34).string(message.courseOfStudy);
    }
    if (message.status !== 0) {
      writer.uint32(40).uint64(message.status);
    }
    if (message.currentYearOfStudy !== 0) {
      writer.uint32(48).uint64(message.currentYearOfStudy);
    }
    if (message.outOfCourse === true) {
      writer.uint32(56).bool(message.outOfCourse);
    }
    if (message.numberOfYearsOutOfCourse !== 0) {
      writer.uint32(64).uint64(message.numberOfYearsOutOfCourse);
    }
    if (message.studentKey !== "") {
      writer.uint32(74).string(message.studentKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StudentInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStudentInfo } as StudentInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.surname = reader.string();
          break;
        case 3:
          message.courseType = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.courseOfStudy = reader.string();
          break;
        case 5:
          message.status = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.currentYearOfStudy = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.outOfCourse = reader.bool();
          break;
        case 8:
          message.numberOfYearsOutOfCourse = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 9:
          message.studentKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StudentInfo {
    const message = { ...baseStudentInfo } as StudentInfo;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.surname !== undefined && object.surname !== null) {
      message.surname = String(object.surname);
    } else {
      message.surname = "";
    }
    if (object.courseType !== undefined && object.courseType !== null) {
      message.courseType = Number(object.courseType);
    } else {
      message.courseType = 0;
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = String(object.courseOfStudy);
    } else {
      message.courseOfStudy = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    if (
      object.currentYearOfStudy !== undefined &&
      object.currentYearOfStudy !== null
    ) {
      message.currentYearOfStudy = Number(object.currentYearOfStudy);
    } else {
      message.currentYearOfStudy = 0;
    }
    if (object.outOfCourse !== undefined && object.outOfCourse !== null) {
      message.outOfCourse = Boolean(object.outOfCourse);
    } else {
      message.outOfCourse = false;
    }
    if (
      object.numberOfYearsOutOfCourse !== undefined &&
      object.numberOfYearsOutOfCourse !== null
    ) {
      message.numberOfYearsOutOfCourse = Number(
        object.numberOfYearsOutOfCourse
      );
    } else {
      message.numberOfYearsOutOfCourse = 0;
    }
    if (object.studentKey !== undefined && object.studentKey !== null) {
      message.studentKey = String(object.studentKey);
    } else {
      message.studentKey = "";
    }
    return message;
  },

  toJSON(message: StudentInfo): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.surname !== undefined && (obj.surname = message.surname);
    message.courseType !== undefined && (obj.courseType = message.courseType);
    message.courseOfStudy !== undefined &&
      (obj.courseOfStudy = message.courseOfStudy);
    message.status !== undefined && (obj.status = message.status);
    message.currentYearOfStudy !== undefined &&
      (obj.currentYearOfStudy = message.currentYearOfStudy);
    message.outOfCourse !== undefined &&
      (obj.outOfCourse = message.outOfCourse);
    message.numberOfYearsOutOfCourse !== undefined &&
      (obj.numberOfYearsOutOfCourse = message.numberOfYearsOutOfCourse);
    message.studentKey !== undefined && (obj.studentKey = message.studentKey);
    return obj;
  },

  fromPartial(object: DeepPartial<StudentInfo>): StudentInfo {
    const message = { ...baseStudentInfo } as StudentInfo;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.surname !== undefined && object.surname !== null) {
      message.surname = object.surname;
    } else {
      message.surname = "";
    }
    if (object.courseType !== undefined && object.courseType !== null) {
      message.courseType = object.courseType;
    } else {
      message.courseType = 0;
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = object.courseOfStudy;
    } else {
      message.courseOfStudy = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (
      object.currentYearOfStudy !== undefined &&
      object.currentYearOfStudy !== null
    ) {
      message.currentYearOfStudy = object.currentYearOfStudy;
    } else {
      message.currentYearOfStudy = 0;
    }
    if (object.outOfCourse !== undefined && object.outOfCourse !== null) {
      message.outOfCourse = object.outOfCourse;
    } else {
      message.outOfCourse = false;
    }
    if (
      object.numberOfYearsOutOfCourse !== undefined &&
      object.numberOfYearsOutOfCourse !== null
    ) {
      message.numberOfYearsOutOfCourse = object.numberOfYearsOutOfCourse;
    } else {
      message.numberOfYearsOutOfCourse = 0;
    }
    if (object.studentKey !== undefined && object.studentKey !== null) {
      message.studentKey = object.studentKey;
    } else {
      message.studentKey = "";
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
