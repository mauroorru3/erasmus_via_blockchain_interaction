/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface UniversityInfo {
  universityName: string;
  nextStudentId: number;
  secretariatKey: string;
  universityKey: string;
  caiKey: string;
  fifoHeadErasmus: string;
  fifoTailErasmus: string;
  deadlineTaxes: string;
  deadlineErasmus: string;
  taxesBrackets: string;
  maxErasmusExams: number;
}

const baseUniversityInfo: object = {
  universityName: "",
  nextStudentId: 0,
  secretariatKey: "",
  universityKey: "",
  caiKey: "",
  fifoHeadErasmus: "",
  fifoTailErasmus: "",
  deadlineTaxes: "",
  deadlineErasmus: "",
  taxesBrackets: "",
  maxErasmusExams: 0,
};

export const UniversityInfo = {
  encode(message: UniversityInfo, writer: Writer = Writer.create()): Writer {
    if (message.universityName !== "") {
      writer.uint32(10).string(message.universityName);
    }
    if (message.nextStudentId !== 0) {
      writer.uint32(16).uint32(message.nextStudentId);
    }
    if (message.secretariatKey !== "") {
      writer.uint32(26).string(message.secretariatKey);
    }
    if (message.universityKey !== "") {
      writer.uint32(34).string(message.universityKey);
    }
    if (message.caiKey !== "") {
      writer.uint32(42).string(message.caiKey);
    }
    if (message.fifoHeadErasmus !== "") {
      writer.uint32(50).string(message.fifoHeadErasmus);
    }
    if (message.fifoTailErasmus !== "") {
      writer.uint32(58).string(message.fifoTailErasmus);
    }
    if (message.deadlineTaxes !== "") {
      writer.uint32(66).string(message.deadlineTaxes);
    }
    if (message.deadlineErasmus !== "") {
      writer.uint32(74).string(message.deadlineErasmus);
    }
    if (message.taxesBrackets !== "") {
      writer.uint32(82).string(message.taxesBrackets);
    }
    if (message.maxErasmusExams !== 0) {
      writer.uint32(88).int32(message.maxErasmusExams);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UniversityInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUniversityInfo } as UniversityInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityName = reader.string();
          break;
        case 2:
          message.nextStudentId = reader.uint32();
          break;
        case 3:
          message.secretariatKey = reader.string();
          break;
        case 4:
          message.universityKey = reader.string();
          break;
        case 5:
          message.caiKey = reader.string();
          break;
        case 6:
          message.fifoHeadErasmus = reader.string();
          break;
        case 7:
          message.fifoTailErasmus = reader.string();
          break;
        case 8:
          message.deadlineTaxes = reader.string();
          break;
        case 9:
          message.deadlineErasmus = reader.string();
          break;
        case 10:
          message.taxesBrackets = reader.string();
          break;
        case 11:
          message.maxErasmusExams = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UniversityInfo {
    const message = { ...baseUniversityInfo } as UniversityInfo;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = String(object.universityName);
    } else {
      message.universityName = "";
    }
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
    if (object.taxesBrackets !== undefined && object.taxesBrackets !== null) {
      message.taxesBrackets = String(object.taxesBrackets);
    } else {
      message.taxesBrackets = "";
    }
    if (
      object.maxErasmusExams !== undefined &&
      object.maxErasmusExams !== null
    ) {
      message.maxErasmusExams = Number(object.maxErasmusExams);
    } else {
      message.maxErasmusExams = 0;
    }
    return message;
  },

  toJSON(message: UniversityInfo): unknown {
    const obj: any = {};
    message.universityName !== undefined &&
      (obj.universityName = message.universityName);
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
    message.taxesBrackets !== undefined &&
      (obj.taxesBrackets = message.taxesBrackets);
    message.maxErasmusExams !== undefined &&
      (obj.maxErasmusExams = message.maxErasmusExams);
    return obj;
  },

  fromPartial(object: DeepPartial<UniversityInfo>): UniversityInfo {
    const message = { ...baseUniversityInfo } as UniversityInfo;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = object.universityName;
    } else {
      message.universityName = "";
    }
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
    if (object.taxesBrackets !== undefined && object.taxesBrackets !== null) {
      message.taxesBrackets = object.taxesBrackets;
    } else {
      message.taxesBrackets = "";
    }
    if (
      object.maxErasmusExams !== undefined &&
      object.maxErasmusExams !== null
    ) {
      message.maxErasmusExams = object.maxErasmusExams;
    } else {
      message.maxErasmusExams = 0;
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
