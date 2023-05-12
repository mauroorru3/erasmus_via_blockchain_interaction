/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { ErasmusContribution } from "../universitychainit/erasmus_contribution";
import { ErasmusExams } from "../universitychainit/erasmus_exams";

export const protobufPackage = "university_chain_it.universitychainit";

export interface ErasmusCareer {
  id: number;
  durationInMonths: number;
  startDate: string;
  endDate: string;
  erasmusType: number;
  totalCredits: number;
  achievedCredits: number;
  totalExams: number;
  examsPassed: number;
  foreignUniversityName: string;
  foreignUniversityCountry: string;
  foreignUniversityStudentId: string;
  status: string;
  contribution: ErasmusContribution | undefined;
  examsData: ErasmusExams | undefined;
}

const baseErasmusCareer: object = {
  id: 0,
  durationInMonths: 0,
  startDate: "",
  endDate: "",
  erasmusType: 0,
  totalCredits: 0,
  achievedCredits: 0,
  totalExams: 0,
  examsPassed: 0,
  foreignUniversityName: "",
  foreignUniversityCountry: "",
  foreignUniversityStudentId: "",
  status: "",
};

export const ErasmusCareer = {
  encode(message: ErasmusCareer, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.durationInMonths !== 0) {
      writer.uint32(16).uint64(message.durationInMonths);
    }
    if (message.startDate !== "") {
      writer.uint32(26).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(34).string(message.endDate);
    }
    if (message.erasmusType !== 0) {
      writer.uint32(40).uint64(message.erasmusType);
    }
    if (message.totalCredits !== 0) {
      writer.uint32(48).uint64(message.totalCredits);
    }
    if (message.achievedCredits !== 0) {
      writer.uint32(56).uint64(message.achievedCredits);
    }
    if (message.totalExams !== 0) {
      writer.uint32(64).uint64(message.totalExams);
    }
    if (message.examsPassed !== 0) {
      writer.uint32(72).uint64(message.examsPassed);
    }
    if (message.foreignUniversityName !== "") {
      writer.uint32(82).string(message.foreignUniversityName);
    }
    if (message.foreignUniversityCountry !== "") {
      writer.uint32(90).string(message.foreignUniversityCountry);
    }
    if (message.foreignUniversityStudentId !== "") {
      writer.uint32(98).string(message.foreignUniversityStudentId);
    }
    if (message.status !== "") {
      writer.uint32(106).string(message.status);
    }
    if (message.contribution !== undefined) {
      ErasmusContribution.encode(
        message.contribution,
        writer.uint32(114).fork()
      ).ldelim();
    }
    if (message.examsData !== undefined) {
      ErasmusExams.encode(
        message.examsData,
        writer.uint32(122).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ErasmusCareer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseErasmusCareer } as ErasmusCareer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.durationInMonths = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.startDate = reader.string();
          break;
        case 4:
          message.endDate = reader.string();
          break;
        case 5:
          message.erasmusType = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.totalCredits = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.achievedCredits = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.totalExams = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.examsPassed = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.foreignUniversityName = reader.string();
          break;
        case 11:
          message.foreignUniversityCountry = reader.string();
          break;
        case 12:
          message.foreignUniversityStudentId = reader.string();
          break;
        case 13:
          message.status = reader.string();
          break;
        case 14:
          message.contribution = ErasmusContribution.decode(
            reader,
            reader.uint32()
          );
          break;
        case 15:
          message.examsData = ErasmusExams.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusCareer {
    const message = { ...baseErasmusCareer } as ErasmusCareer;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (
      object.durationInMonths !== undefined &&
      object.durationInMonths !== null
    ) {
      message.durationInMonths = Number(object.durationInMonths);
    } else {
      message.durationInMonths = 0;
    }
    if (object.startDate !== undefined && object.startDate !== null) {
      message.startDate = String(object.startDate);
    } else {
      message.startDate = "";
    }
    if (object.endDate !== undefined && object.endDate !== null) {
      message.endDate = String(object.endDate);
    } else {
      message.endDate = "";
    }
    if (object.erasmusType !== undefined && object.erasmusType !== null) {
      message.erasmusType = Number(object.erasmusType);
    } else {
      message.erasmusType = 0;
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
    if (
      object.foreignUniversityName !== undefined &&
      object.foreignUniversityName !== null
    ) {
      message.foreignUniversityName = String(object.foreignUniversityName);
    } else {
      message.foreignUniversityName = "";
    }
    if (
      object.foreignUniversityCountry !== undefined &&
      object.foreignUniversityCountry !== null
    ) {
      message.foreignUniversityCountry = String(
        object.foreignUniversityCountry
      );
    } else {
      message.foreignUniversityCountry = "";
    }
    if (
      object.foreignUniversityStudentId !== undefined &&
      object.foreignUniversityStudentId !== null
    ) {
      message.foreignUniversityStudentId = String(
        object.foreignUniversityStudentId
      );
    } else {
      message.foreignUniversityStudentId = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.contribution !== undefined && object.contribution !== null) {
      message.contribution = ErasmusContribution.fromJSON(object.contribution);
    } else {
      message.contribution = undefined;
    }
    if (object.examsData !== undefined && object.examsData !== null) {
      message.examsData = ErasmusExams.fromJSON(object.examsData);
    } else {
      message.examsData = undefined;
    }
    return message;
  },

  toJSON(message: ErasmusCareer): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.durationInMonths !== undefined &&
      (obj.durationInMonths = message.durationInMonths);
    message.startDate !== undefined && (obj.startDate = message.startDate);
    message.endDate !== undefined && (obj.endDate = message.endDate);
    message.erasmusType !== undefined &&
      (obj.erasmusType = message.erasmusType);
    message.totalCredits !== undefined &&
      (obj.totalCredits = message.totalCredits);
    message.achievedCredits !== undefined &&
      (obj.achievedCredits = message.achievedCredits);
    message.totalExams !== undefined && (obj.totalExams = message.totalExams);
    message.examsPassed !== undefined &&
      (obj.examsPassed = message.examsPassed);
    message.foreignUniversityName !== undefined &&
      (obj.foreignUniversityName = message.foreignUniversityName);
    message.foreignUniversityCountry !== undefined &&
      (obj.foreignUniversityCountry = message.foreignUniversityCountry);
    message.foreignUniversityStudentId !== undefined &&
      (obj.foreignUniversityStudentId = message.foreignUniversityStudentId);
    message.status !== undefined && (obj.status = message.status);
    message.contribution !== undefined &&
      (obj.contribution = message.contribution
        ? ErasmusContribution.toJSON(message.contribution)
        : undefined);
    message.examsData !== undefined &&
      (obj.examsData = message.examsData
        ? ErasmusExams.toJSON(message.examsData)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<ErasmusCareer>): ErasmusCareer {
    const message = { ...baseErasmusCareer } as ErasmusCareer;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (
      object.durationInMonths !== undefined &&
      object.durationInMonths !== null
    ) {
      message.durationInMonths = object.durationInMonths;
    } else {
      message.durationInMonths = 0;
    }
    if (object.startDate !== undefined && object.startDate !== null) {
      message.startDate = object.startDate;
    } else {
      message.startDate = "";
    }
    if (object.endDate !== undefined && object.endDate !== null) {
      message.endDate = object.endDate;
    } else {
      message.endDate = "";
    }
    if (object.erasmusType !== undefined && object.erasmusType !== null) {
      message.erasmusType = object.erasmusType;
    } else {
      message.erasmusType = 0;
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
    if (
      object.foreignUniversityName !== undefined &&
      object.foreignUniversityName !== null
    ) {
      message.foreignUniversityName = object.foreignUniversityName;
    } else {
      message.foreignUniversityName = "";
    }
    if (
      object.foreignUniversityCountry !== undefined &&
      object.foreignUniversityCountry !== null
    ) {
      message.foreignUniversityCountry = object.foreignUniversityCountry;
    } else {
      message.foreignUniversityCountry = "";
    }
    if (
      object.foreignUniversityStudentId !== undefined &&
      object.foreignUniversityStudentId !== null
    ) {
      message.foreignUniversityStudentId = object.foreignUniversityStudentId;
    } else {
      message.foreignUniversityStudentId = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.contribution !== undefined && object.contribution !== null) {
      message.contribution = ErasmusContribution.fromPartial(
        object.contribution
      );
    } else {
      message.contribution = undefined;
    }
    if (object.examsData !== undefined && object.examsData !== null) {
      message.examsData = ErasmusExams.fromPartial(object.examsData);
    } else {
      message.examsData = undefined;
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
