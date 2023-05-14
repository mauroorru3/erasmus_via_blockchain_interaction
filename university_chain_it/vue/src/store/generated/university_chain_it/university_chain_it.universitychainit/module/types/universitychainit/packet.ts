/* eslint-disable */
import { StoredStudent } from "../universitychainit/stored_student";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface UniversitychainitPacketData {
  noData: NoData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  erasmusStudentPacket: ErasmusStudentPacketData | undefined;
}

export interface NoData {}

/** ErasmusStudentPacketData defines a struct for the packet payload */
export interface ErasmusStudentPacketData {
  student: StoredStudent | undefined;
}

/** ErasmusStudentPacketAck defines a struct for the packet acknowledgment */
export interface ErasmusStudentPacketAck {
  index: string;
  foreign_index: string;
}

const baseUniversitychainitPacketData: object = {};

export const UniversitychainitPacketData = {
  encode(
    message: UniversitychainitPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.erasmusStudentPacket !== undefined) {
      ErasmusStudentPacketData.encode(
        message.erasmusStudentPacket,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): UniversitychainitPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseUniversitychainitPacketData,
    } as UniversitychainitPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 2:
          message.erasmusStudentPacket = ErasmusStudentPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UniversitychainitPacketData {
    const message = {
      ...baseUniversitychainitPacketData,
    } as UniversitychainitPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.erasmusStudentPacket !== undefined &&
      object.erasmusStudentPacket !== null
    ) {
      message.erasmusStudentPacket = ErasmusStudentPacketData.fromJSON(
        object.erasmusStudentPacket
      );
    } else {
      message.erasmusStudentPacket = undefined;
    }
    return message;
  },

  toJSON(message: UniversitychainitPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined &&
      (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.erasmusStudentPacket !== undefined &&
      (obj.erasmusStudentPacket = message.erasmusStudentPacket
        ? ErasmusStudentPacketData.toJSON(message.erasmusStudentPacket)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<UniversitychainitPacketData>
  ): UniversitychainitPacketData {
    const message = {
      ...baseUniversitychainitPacketData,
    } as UniversitychainitPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.erasmusStudentPacket !== undefined &&
      object.erasmusStudentPacket !== null
    ) {
      message.erasmusStudentPacket = ErasmusStudentPacketData.fromPartial(
        object.erasmusStudentPacket
      );
    } else {
      message.erasmusStudentPacket = undefined;
    }
    return message;
  },
};

const baseNoData: object = {};

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNoData } as NoData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },
};

const baseErasmusStudentPacketData: object = {};

export const ErasmusStudentPacketData = {
  encode(
    message: ErasmusStudentPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.student !== undefined) {
      StoredStudent.encode(message.student, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ErasmusStudentPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseErasmusStudentPacketData,
    } as ErasmusStudentPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.student = StoredStudent.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusStudentPacketData {
    const message = {
      ...baseErasmusStudentPacketData,
    } as ErasmusStudentPacketData;
    if (object.student !== undefined && object.student !== null) {
      message.student = StoredStudent.fromJSON(object.student);
    } else {
      message.student = undefined;
    }
    return message;
  },

  toJSON(message: ErasmusStudentPacketData): unknown {
    const obj: any = {};
    message.student !== undefined &&
      (obj.student = message.student
        ? StoredStudent.toJSON(message.student)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ErasmusStudentPacketData>
  ): ErasmusStudentPacketData {
    const message = {
      ...baseErasmusStudentPacketData,
    } as ErasmusStudentPacketData;
    if (object.student !== undefined && object.student !== null) {
      message.student = StoredStudent.fromPartial(object.student);
    } else {
      message.student = undefined;
    }
    return message;
  },
};

const baseErasmusStudentPacketAck: object = { index: "", foreign_index: "" };

export const ErasmusStudentPacketAck = {
  encode(
    message: ErasmusStudentPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.foreign_index !== "") {
      writer.uint32(18).string(message.foreign_index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ErasmusStudentPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseErasmusStudentPacketAck,
    } as ErasmusStudentPacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.foreign_index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusStudentPacketAck {
    const message = {
      ...baseErasmusStudentPacketAck,
    } as ErasmusStudentPacketAck;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.foreign_index !== undefined && object.foreign_index !== null) {
      message.foreign_index = String(object.foreign_index);
    } else {
      message.foreign_index = "";
    }
    return message;
  },

  toJSON(message: ErasmusStudentPacketAck): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.foreign_index !== undefined &&
      (obj.foreign_index = message.foreign_index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ErasmusStudentPacketAck>
  ): ErasmusStudentPacketAck {
    const message = {
      ...baseErasmusStudentPacketAck,
    } as ErasmusStudentPacketAck;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.foreign_index !== undefined && object.foreign_index !== null) {
      message.foreign_index = object.foreign_index;
    } else {
      message.foreign_index = "";
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
