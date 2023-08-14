/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface ProfessorsExams {
  examName: string;
  professorName: string;
  professorId: string;
  professorKey: string;
}

const baseProfessorsExams: object = {
  examName: "",
  professorName: "",
  professorId: "",
  professorKey: "",
};

export const ProfessorsExams = {
  encode(message: ProfessorsExams, writer: Writer = Writer.create()): Writer {
    if (message.examName !== "") {
      writer.uint32(10).string(message.examName);
    }
    if (message.professorName !== "") {
      writer.uint32(18).string(message.professorName);
    }
    if (message.professorId !== "") {
      writer.uint32(26).string(message.professorId);
    }
    if (message.professorKey !== "") {
      writer.uint32(34).string(message.professorKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ProfessorsExams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProfessorsExams } as ProfessorsExams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.examName = reader.string();
          break;
        case 2:
          message.professorName = reader.string();
          break;
        case 3:
          message.professorId = reader.string();
          break;
        case 4:
          message.professorKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ProfessorsExams {
    const message = { ...baseProfessorsExams } as ProfessorsExams;
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = String(object.examName);
    } else {
      message.examName = "";
    }
    if (object.professorName !== undefined && object.professorName !== null) {
      message.professorName = String(object.professorName);
    } else {
      message.professorName = "";
    }
    if (object.professorId !== undefined && object.professorId !== null) {
      message.professorId = String(object.professorId);
    } else {
      message.professorId = "";
    }
    if (object.professorKey !== undefined && object.professorKey !== null) {
      message.professorKey = String(object.professorKey);
    } else {
      message.professorKey = "";
    }
    return message;
  },

  toJSON(message: ProfessorsExams): unknown {
    const obj: any = {};
    message.examName !== undefined && (obj.examName = message.examName);
    message.professorName !== undefined &&
      (obj.professorName = message.professorName);
    message.professorId !== undefined &&
      (obj.professorId = message.professorId);
    message.professorKey !== undefined &&
      (obj.professorKey = message.professorKey);
    return obj;
  },

  fromPartial(object: DeepPartial<ProfessorsExams>): ProfessorsExams {
    const message = { ...baseProfessorsExams } as ProfessorsExams;
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = object.examName;
    } else {
      message.examName = "";
    }
    if (object.professorName !== undefined && object.professorName !== null) {
      message.professorName = object.professorName;
    } else {
      message.professorName = "";
    }
    if (object.professorId !== undefined && object.professorId !== null) {
      message.professorId = object.professorId;
    } else {
      message.professorId = "";
    }
    if (object.professorKey !== undefined && object.professorKey !== null) {
      message.professorKey = object.professorKey;
    } else {
      message.professorKey = "";
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
