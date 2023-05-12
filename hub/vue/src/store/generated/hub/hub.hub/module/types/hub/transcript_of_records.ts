/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface TranscriptOfRecords {
  examsData: string;
  totalExams: number;
  examsPassed: number;
  totalCredits: number;
  achievedCredits: number;
}

const baseTranscriptOfRecords: object = {
  examsData: "",
  totalExams: 0,
  examsPassed: 0,
  totalCredits: 0,
  achievedCredits: 0,
};

export const TranscriptOfRecords = {
  encode(
    message: TranscriptOfRecords,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.examsData !== "") {
      writer.uint32(10).string(message.examsData);
    }
    if (message.totalExams !== 0) {
      writer.uint32(16).uint64(message.totalExams);
    }
    if (message.examsPassed !== 0) {
      writer.uint32(24).uint64(message.examsPassed);
    }
    if (message.totalCredits !== 0) {
      writer.uint32(32).uint64(message.totalCredits);
    }
    if (message.achievedCredits !== 0) {
      writer.uint32(40).uint64(message.achievedCredits);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TranscriptOfRecords {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTranscriptOfRecords } as TranscriptOfRecords;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.examsData = reader.string();
          break;
        case 2:
          message.totalExams = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.examsPassed = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.totalCredits = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.achievedCredits = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TranscriptOfRecords {
    const message = { ...baseTranscriptOfRecords } as TranscriptOfRecords;
    if (object.examsData !== undefined && object.examsData !== null) {
      message.examsData = String(object.examsData);
    } else {
      message.examsData = "";
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
    return message;
  },

  toJSON(message: TranscriptOfRecords): unknown {
    const obj: any = {};
    message.examsData !== undefined && (obj.examsData = message.examsData);
    message.totalExams !== undefined && (obj.totalExams = message.totalExams);
    message.examsPassed !== undefined &&
      (obj.examsPassed = message.examsPassed);
    message.totalCredits !== undefined &&
      (obj.totalCredits = message.totalCredits);
    message.achievedCredits !== undefined &&
      (obj.achievedCredits = message.achievedCredits);
    return obj;
  },

  fromPartial(object: DeepPartial<TranscriptOfRecords>): TranscriptOfRecords {
    const message = { ...baseTranscriptOfRecords } as TranscriptOfRecords;
    if (object.examsData !== undefined && object.examsData !== null) {
      message.examsData = object.examsData;
    } else {
      message.examsData = "";
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
