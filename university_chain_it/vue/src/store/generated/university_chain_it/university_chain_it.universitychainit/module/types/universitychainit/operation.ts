/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface OperationInfo {
  studentOperationDetails: string;
  previousStudentOperationFifo: string;
  nextStudentOperationFifo: string;
}

const baseOperationInfo: object = {
  studentOperationDetails: "",
  previousStudentOperationFifo: "",
  nextStudentOperationFifo: "",
};

export const OperationInfo = {
  encode(message: OperationInfo, writer: Writer = Writer.create()): Writer {
    if (message.studentOperationDetails !== "") {
      writer.uint32(10).string(message.studentOperationDetails);
    }
    if (message.previousStudentOperationFifo !== "") {
      writer.uint32(18).string(message.previousStudentOperationFifo);
    }
    if (message.nextStudentOperationFifo !== "") {
      writer.uint32(26).string(message.nextStudentOperationFifo);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OperationInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOperationInfo } as OperationInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.studentOperationDetails = reader.string();
          break;
        case 2:
          message.previousStudentOperationFifo = reader.string();
          break;
        case 3:
          message.nextStudentOperationFifo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OperationInfo {
    const message = { ...baseOperationInfo } as OperationInfo;
    if (
      object.studentOperationDetails !== undefined &&
      object.studentOperationDetails !== null
    ) {
      message.studentOperationDetails = String(object.studentOperationDetails);
    } else {
      message.studentOperationDetails = "";
    }
    if (
      object.previousStudentOperationFifo !== undefined &&
      object.previousStudentOperationFifo !== null
    ) {
      message.previousStudentOperationFifo = String(
        object.previousStudentOperationFifo
      );
    } else {
      message.previousStudentOperationFifo = "";
    }
    if (
      object.nextStudentOperationFifo !== undefined &&
      object.nextStudentOperationFifo !== null
    ) {
      message.nextStudentOperationFifo = String(
        object.nextStudentOperationFifo
      );
    } else {
      message.nextStudentOperationFifo = "";
    }
    return message;
  },

  toJSON(message: OperationInfo): unknown {
    const obj: any = {};
    message.studentOperationDetails !== undefined &&
      (obj.studentOperationDetails = message.studentOperationDetails);
    message.previousStudentOperationFifo !== undefined &&
      (obj.previousStudentOperationFifo = message.previousStudentOperationFifo);
    message.nextStudentOperationFifo !== undefined &&
      (obj.nextStudentOperationFifo = message.nextStudentOperationFifo);
    return obj;
  },

  fromPartial(object: DeepPartial<OperationInfo>): OperationInfo {
    const message = { ...baseOperationInfo } as OperationInfo;
    if (
      object.studentOperationDetails !== undefined &&
      object.studentOperationDetails !== null
    ) {
      message.studentOperationDetails = object.studentOperationDetails;
    } else {
      message.studentOperationDetails = "";
    }
    if (
      object.previousStudentOperationFifo !== undefined &&
      object.previousStudentOperationFifo !== null
    ) {
      message.previousStudentOperationFifo =
        object.previousStudentOperationFifo;
    } else {
      message.previousStudentOperationFifo = "";
    }
    if (
      object.nextStudentOperationFifo !== undefined &&
      object.nextStudentOperationFifo !== null
    ) {
      message.nextStudentOperationFifo = object.nextStudentOperationFifo;
    } else {
      message.nextStudentOperationFifo = "";
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
