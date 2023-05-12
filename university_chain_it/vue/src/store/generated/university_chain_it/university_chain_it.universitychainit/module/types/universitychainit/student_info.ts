/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface StudentInfo {
  name: string;
  surname: string;
  courseType: string;
  courseOfStudy: string;
  status: string;
  currentYearOfStudy: number;
  outOfCourse: boolean;
  numberOfYearsOutOfCourse: number;
  studentKey: string;
  completeInformation: number[];
}

const baseStudentInfo: object = {
  name: "",
  surname: "",
  courseType: "",
  courseOfStudy: "",
  status: "",
  currentYearOfStudy: 0,
  outOfCourse: false,
  numberOfYearsOutOfCourse: 0,
  studentKey: "",
  completeInformation: 0,
};

export const StudentInfo = {
  encode(message: StudentInfo, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.surname !== "") {
      writer.uint32(18).string(message.surname);
    }
    if (message.courseType !== "") {
      writer.uint32(26).string(message.courseType);
    }
    if (message.courseOfStudy !== "") {
      writer.uint32(34).string(message.courseOfStudy);
    }
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    if (message.currentYearOfStudy !== 0) {
      writer.uint32(48).uint32(message.currentYearOfStudy);
    }
    if (message.outOfCourse === true) {
      writer.uint32(56).bool(message.outOfCourse);
    }
    if (message.numberOfYearsOutOfCourse !== 0) {
      writer.uint32(64).uint32(message.numberOfYearsOutOfCourse);
    }
    if (message.studentKey !== "") {
      writer.uint32(74).string(message.studentKey);
    }
    writer.uint32(82).fork();
    for (const v of message.completeInformation) {
      writer.int32(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StudentInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStudentInfo } as StudentInfo;
    message.completeInformation = [];
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
          message.courseType = reader.string();
          break;
        case 4:
          message.courseOfStudy = reader.string();
          break;
        case 5:
          message.status = reader.string();
          break;
        case 6:
          message.currentYearOfStudy = reader.uint32();
          break;
        case 7:
          message.outOfCourse = reader.bool();
          break;
        case 8:
          message.numberOfYearsOutOfCourse = reader.uint32();
          break;
        case 9:
          message.studentKey = reader.string();
          break;
        case 10:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.completeInformation.push(reader.int32());
            }
          } else {
            message.completeInformation.push(reader.int32());
          }
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
    message.completeInformation = [];
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
      message.courseType = String(object.courseType);
    } else {
      message.courseType = "";
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = String(object.courseOfStudy);
    } else {
      message.courseOfStudy = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
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
    if (
      object.completeInformation !== undefined &&
      object.completeInformation !== null
    ) {
      for (const e of object.completeInformation) {
        message.completeInformation.push(Number(e));
      }
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
    if (message.completeInformation) {
      obj.completeInformation = message.completeInformation.map((e) => e);
    } else {
      obj.completeInformation = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<StudentInfo>): StudentInfo {
    const message = { ...baseStudentInfo } as StudentInfo;
    message.completeInformation = [];
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
      message.courseType = "";
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = object.courseOfStudy;
    } else {
      message.courseOfStudy = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
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
    if (
      object.completeInformation !== undefined &&
      object.completeInformation !== null
    ) {
      for (const e of object.completeInformation) {
        message.completeInformation.push(e);
      }
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
