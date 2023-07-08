/* eslint-disable */
import { StoredStudent } from "../universitychainde/stored_student";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_de.universitychainde";

export interface UniversitychaindePacketData {
  noData: NoData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  erasmusRestictedDataPacket: ErasmusRestictedDataPacketData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  extendErasmusPeriodPacket: ExtendErasmusPeriodPacketData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  finalErasmusDataPacket: FinalErasmusDataPacketData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  endErasmusPeriodRequestPacket: EndErasmusPeriodRequestPacketData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  erasmusIndexPacket: ErasmusIndexPacketData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  erasmusStudentPacket: ErasmusStudentPacketData | undefined;
}

export interface NoData {}

/** ErasmusStudentPacketData defines a struct for the packet payload */
export interface ErasmusStudentPacketData {
  student: StoredStudent | undefined;
}

/** ErasmusStudentPacketAck defines a struct for the packet acknowledgment */
export interface ErasmusStudentPacketAck {
  foreignIndex: string;
}

/** ErasmusIndexPacketData defines a struct for the packet payload */
export interface ErasmusIndexPacketData {
  index: string;
  foreignIndex: string;
}

/** ErasmusIndexPacketAck defines a struct for the packet acknowledgment */
export interface ErasmusIndexPacketAck {}

/** EndErasmusPeriodRequestPacketData defines a struct for the packet payload */
export interface EndErasmusPeriodRequestPacketData {
  startingUniversityName: string;
  destinationUniversityName: string;
  index: string;
  foreignIndex: string;
}

/** EndErasmusPeriodRequestPacketAck defines a struct for the packet acknowledgment */
export interface EndErasmusPeriodRequestPacketAck {
  erasmusRestrictedInfo: string;
}

/** FinalErasmusDataPacketData defines a struct for the packet payload */
export interface FinalErasmusDataPacketData {
  erasmusRestrictedInfo: string;
  homeIndex: string;
}

/** FinalErasmusDataPacketAck defines a struct for the packet acknowledgment */
export interface FinalErasmusDataPacketAck {}

/** ExtendErasmusPeriodPacketData defines a struct for the packet payload */
export interface ExtendErasmusPeriodPacketData {
  durationInMonths: number;
  destinationUniversityName: string;
  foreignIndex: string;
  finalDate: string;
}

/** ExtendErasmusPeriodPacketAck defines a struct for the packet acknowledgment */
export interface ExtendErasmusPeriodPacketAck {}

/** ErasmusRestictedDataPacketData defines a struct for the packet payload */
export interface ErasmusRestictedDataPacketData {
  erasmusRestrictedInfo: string;
}

/** ErasmusRestictedDataPacketAck defines a struct for the packet acknowledgment */
export interface ErasmusRestictedDataPacketAck {
  erasmusRestrictedInfo: string;
}

const baseUniversitychaindePacketData: object = {};

export const UniversitychaindePacketData = {
  encode(
    message: UniversitychaindePacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.erasmusRestictedDataPacket !== undefined) {
      ErasmusRestictedDataPacketData.encode(
        message.erasmusRestictedDataPacket,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.extendErasmusPeriodPacket !== undefined) {
      ExtendErasmusPeriodPacketData.encode(
        message.extendErasmusPeriodPacket,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.finalErasmusDataPacket !== undefined) {
      FinalErasmusDataPacketData.encode(
        message.finalErasmusDataPacket,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.endErasmusPeriodRequestPacket !== undefined) {
      EndErasmusPeriodRequestPacketData.encode(
        message.endErasmusPeriodRequestPacket,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.erasmusIndexPacket !== undefined) {
      ErasmusIndexPacketData.encode(
        message.erasmusIndexPacket,
        writer.uint32(26).fork()
      ).ldelim();
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
  ): UniversitychaindePacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseUniversitychaindePacketData,
    } as UniversitychaindePacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 7:
          message.erasmusRestictedDataPacket = ErasmusRestictedDataPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.extendErasmusPeriodPacket = ExtendErasmusPeriodPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.finalErasmusDataPacket = FinalErasmusDataPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        case 4:
          message.endErasmusPeriodRequestPacket = EndErasmusPeriodRequestPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        case 3:
          message.erasmusIndexPacket = ErasmusIndexPacketData.decode(
            reader,
            reader.uint32()
          );
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

  fromJSON(object: any): UniversitychaindePacketData {
    const message = {
      ...baseUniversitychaindePacketData,
    } as UniversitychaindePacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.erasmusRestictedDataPacket !== undefined &&
      object.erasmusRestictedDataPacket !== null
    ) {
      message.erasmusRestictedDataPacket = ErasmusRestictedDataPacketData.fromJSON(
        object.erasmusRestictedDataPacket
      );
    } else {
      message.erasmusRestictedDataPacket = undefined;
    }
    if (
      object.extendErasmusPeriodPacket !== undefined &&
      object.extendErasmusPeriodPacket !== null
    ) {
      message.extendErasmusPeriodPacket = ExtendErasmusPeriodPacketData.fromJSON(
        object.extendErasmusPeriodPacket
      );
    } else {
      message.extendErasmusPeriodPacket = undefined;
    }
    if (
      object.finalErasmusDataPacket !== undefined &&
      object.finalErasmusDataPacket !== null
    ) {
      message.finalErasmusDataPacket = FinalErasmusDataPacketData.fromJSON(
        object.finalErasmusDataPacket
      );
    } else {
      message.finalErasmusDataPacket = undefined;
    }
    if (
      object.endErasmusPeriodRequestPacket !== undefined &&
      object.endErasmusPeriodRequestPacket !== null
    ) {
      message.endErasmusPeriodRequestPacket = EndErasmusPeriodRequestPacketData.fromJSON(
        object.endErasmusPeriodRequestPacket
      );
    } else {
      message.endErasmusPeriodRequestPacket = undefined;
    }
    if (
      object.erasmusIndexPacket !== undefined &&
      object.erasmusIndexPacket !== null
    ) {
      message.erasmusIndexPacket = ErasmusIndexPacketData.fromJSON(
        object.erasmusIndexPacket
      );
    } else {
      message.erasmusIndexPacket = undefined;
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

  toJSON(message: UniversitychaindePacketData): unknown {
    const obj: any = {};
    message.noData !== undefined &&
      (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.erasmusRestictedDataPacket !== undefined &&
      (obj.erasmusRestictedDataPacket = message.erasmusRestictedDataPacket
        ? ErasmusRestictedDataPacketData.toJSON(
            message.erasmusRestictedDataPacket
          )
        : undefined);
    message.extendErasmusPeriodPacket !== undefined &&
      (obj.extendErasmusPeriodPacket = message.extendErasmusPeriodPacket
        ? ExtendErasmusPeriodPacketData.toJSON(
            message.extendErasmusPeriodPacket
          )
        : undefined);
    message.finalErasmusDataPacket !== undefined &&
      (obj.finalErasmusDataPacket = message.finalErasmusDataPacket
        ? FinalErasmusDataPacketData.toJSON(message.finalErasmusDataPacket)
        : undefined);
    message.endErasmusPeriodRequestPacket !== undefined &&
      (obj.endErasmusPeriodRequestPacket = message.endErasmusPeriodRequestPacket
        ? EndErasmusPeriodRequestPacketData.toJSON(
            message.endErasmusPeriodRequestPacket
          )
        : undefined);
    message.erasmusIndexPacket !== undefined &&
      (obj.erasmusIndexPacket = message.erasmusIndexPacket
        ? ErasmusIndexPacketData.toJSON(message.erasmusIndexPacket)
        : undefined);
    message.erasmusStudentPacket !== undefined &&
      (obj.erasmusStudentPacket = message.erasmusStudentPacket
        ? ErasmusStudentPacketData.toJSON(message.erasmusStudentPacket)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<UniversitychaindePacketData>
  ): UniversitychaindePacketData {
    const message = {
      ...baseUniversitychaindePacketData,
    } as UniversitychaindePacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.erasmusRestictedDataPacket !== undefined &&
      object.erasmusRestictedDataPacket !== null
    ) {
      message.erasmusRestictedDataPacket = ErasmusRestictedDataPacketData.fromPartial(
        object.erasmusRestictedDataPacket
      );
    } else {
      message.erasmusRestictedDataPacket = undefined;
    }
    if (
      object.extendErasmusPeriodPacket !== undefined &&
      object.extendErasmusPeriodPacket !== null
    ) {
      message.extendErasmusPeriodPacket = ExtendErasmusPeriodPacketData.fromPartial(
        object.extendErasmusPeriodPacket
      );
    } else {
      message.extendErasmusPeriodPacket = undefined;
    }
    if (
      object.finalErasmusDataPacket !== undefined &&
      object.finalErasmusDataPacket !== null
    ) {
      message.finalErasmusDataPacket = FinalErasmusDataPacketData.fromPartial(
        object.finalErasmusDataPacket
      );
    } else {
      message.finalErasmusDataPacket = undefined;
    }
    if (
      object.endErasmusPeriodRequestPacket !== undefined &&
      object.endErasmusPeriodRequestPacket !== null
    ) {
      message.endErasmusPeriodRequestPacket = EndErasmusPeriodRequestPacketData.fromPartial(
        object.endErasmusPeriodRequestPacket
      );
    } else {
      message.endErasmusPeriodRequestPacket = undefined;
    }
    if (
      object.erasmusIndexPacket !== undefined &&
      object.erasmusIndexPacket !== null
    ) {
      message.erasmusIndexPacket = ErasmusIndexPacketData.fromPartial(
        object.erasmusIndexPacket
      );
    } else {
      message.erasmusIndexPacket = undefined;
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

const baseErasmusStudentPacketAck: object = { foreignIndex: "" };

export const ErasmusStudentPacketAck = {
  encode(
    message: ErasmusStudentPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.foreignIndex !== "") {
      writer.uint32(10).string(message.foreignIndex);
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
          message.foreignIndex = reader.string();
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
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = String(object.foreignIndex);
    } else {
      message.foreignIndex = "";
    }
    return message;
  },

  toJSON(message: ErasmusStudentPacketAck): unknown {
    const obj: any = {};
    message.foreignIndex !== undefined &&
      (obj.foreignIndex = message.foreignIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ErasmusStudentPacketAck>
  ): ErasmusStudentPacketAck {
    const message = {
      ...baseErasmusStudentPacketAck,
    } as ErasmusStudentPacketAck;
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = object.foreignIndex;
    } else {
      message.foreignIndex = "";
    }
    return message;
  },
};

const baseErasmusIndexPacketData: object = { index: "", foreignIndex: "" };

export const ErasmusIndexPacketData = {
  encode(
    message: ErasmusIndexPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.foreignIndex !== "") {
      writer.uint32(18).string(message.foreignIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ErasmusIndexPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseErasmusIndexPacketData } as ErasmusIndexPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.foreignIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusIndexPacketData {
    const message = { ...baseErasmusIndexPacketData } as ErasmusIndexPacketData;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = String(object.foreignIndex);
    } else {
      message.foreignIndex = "";
    }
    return message;
  },

  toJSON(message: ErasmusIndexPacketData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.foreignIndex !== undefined &&
      (obj.foreignIndex = message.foreignIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ErasmusIndexPacketData>
  ): ErasmusIndexPacketData {
    const message = { ...baseErasmusIndexPacketData } as ErasmusIndexPacketData;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = object.foreignIndex;
    } else {
      message.foreignIndex = "";
    }
    return message;
  },
};

const baseErasmusIndexPacketAck: object = {};

export const ErasmusIndexPacketAck = {
  encode(_: ErasmusIndexPacketAck, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ErasmusIndexPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseErasmusIndexPacketAck } as ErasmusIndexPacketAck;
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

  fromJSON(_: any): ErasmusIndexPacketAck {
    const message = { ...baseErasmusIndexPacketAck } as ErasmusIndexPacketAck;
    return message;
  },

  toJSON(_: ErasmusIndexPacketAck): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<ErasmusIndexPacketAck>): ErasmusIndexPacketAck {
    const message = { ...baseErasmusIndexPacketAck } as ErasmusIndexPacketAck;
    return message;
  },
};

const baseEndErasmusPeriodRequestPacketData: object = {
  startingUniversityName: "",
  destinationUniversityName: "",
  index: "",
  foreignIndex: "",
};

export const EndErasmusPeriodRequestPacketData = {
  encode(
    message: EndErasmusPeriodRequestPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.startingUniversityName !== "") {
      writer.uint32(10).string(message.startingUniversityName);
    }
    if (message.destinationUniversityName !== "") {
      writer.uint32(18).string(message.destinationUniversityName);
    }
    if (message.index !== "") {
      writer.uint32(26).string(message.index);
    }
    if (message.foreignIndex !== "") {
      writer.uint32(34).string(message.foreignIndex);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EndErasmusPeriodRequestPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEndErasmusPeriodRequestPacketData,
    } as EndErasmusPeriodRequestPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.startingUniversityName = reader.string();
          break;
        case 2:
          message.destinationUniversityName = reader.string();
          break;
        case 3:
          message.index = reader.string();
          break;
        case 4:
          message.foreignIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EndErasmusPeriodRequestPacketData {
    const message = {
      ...baseEndErasmusPeriodRequestPacketData,
    } as EndErasmusPeriodRequestPacketData;
    if (
      object.startingUniversityName !== undefined &&
      object.startingUniversityName !== null
    ) {
      message.startingUniversityName = String(object.startingUniversityName);
    } else {
      message.startingUniversityName = "";
    }
    if (
      object.destinationUniversityName !== undefined &&
      object.destinationUniversityName !== null
    ) {
      message.destinationUniversityName = String(
        object.destinationUniversityName
      );
    } else {
      message.destinationUniversityName = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = String(object.foreignIndex);
    } else {
      message.foreignIndex = "";
    }
    return message;
  },

  toJSON(message: EndErasmusPeriodRequestPacketData): unknown {
    const obj: any = {};
    message.startingUniversityName !== undefined &&
      (obj.startingUniversityName = message.startingUniversityName);
    message.destinationUniversityName !== undefined &&
      (obj.destinationUniversityName = message.destinationUniversityName);
    message.index !== undefined && (obj.index = message.index);
    message.foreignIndex !== undefined &&
      (obj.foreignIndex = message.foreignIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EndErasmusPeriodRequestPacketData>
  ): EndErasmusPeriodRequestPacketData {
    const message = {
      ...baseEndErasmusPeriodRequestPacketData,
    } as EndErasmusPeriodRequestPacketData;
    if (
      object.startingUniversityName !== undefined &&
      object.startingUniversityName !== null
    ) {
      message.startingUniversityName = object.startingUniversityName;
    } else {
      message.startingUniversityName = "";
    }
    if (
      object.destinationUniversityName !== undefined &&
      object.destinationUniversityName !== null
    ) {
      message.destinationUniversityName = object.destinationUniversityName;
    } else {
      message.destinationUniversityName = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = object.foreignIndex;
    } else {
      message.foreignIndex = "";
    }
    return message;
  },
};

const baseEndErasmusPeriodRequestPacketAck: object = {
  erasmusRestrictedInfo: "",
};

export const EndErasmusPeriodRequestPacketAck = {
  encode(
    message: EndErasmusPeriodRequestPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.erasmusRestrictedInfo !== "") {
      writer.uint32(10).string(message.erasmusRestrictedInfo);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EndErasmusPeriodRequestPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEndErasmusPeriodRequestPacketAck,
    } as EndErasmusPeriodRequestPacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erasmusRestrictedInfo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EndErasmusPeriodRequestPacketAck {
    const message = {
      ...baseEndErasmusPeriodRequestPacketAck,
    } as EndErasmusPeriodRequestPacketAck;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = String(object.erasmusRestrictedInfo);
    } else {
      message.erasmusRestrictedInfo = "";
    }
    return message;
  },

  toJSON(message: EndErasmusPeriodRequestPacketAck): unknown {
    const obj: any = {};
    message.erasmusRestrictedInfo !== undefined &&
      (obj.erasmusRestrictedInfo = message.erasmusRestrictedInfo);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EndErasmusPeriodRequestPacketAck>
  ): EndErasmusPeriodRequestPacketAck {
    const message = {
      ...baseEndErasmusPeriodRequestPacketAck,
    } as EndErasmusPeriodRequestPacketAck;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = object.erasmusRestrictedInfo;
    } else {
      message.erasmusRestrictedInfo = "";
    }
    return message;
  },
};

const baseFinalErasmusDataPacketData: object = {
  erasmusRestrictedInfo: "",
  homeIndex: "",
};

export const FinalErasmusDataPacketData = {
  encode(
    message: FinalErasmusDataPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.erasmusRestrictedInfo !== "") {
      writer.uint32(10).string(message.erasmusRestrictedInfo);
    }
    if (message.homeIndex !== "") {
      writer.uint32(18).string(message.homeIndex);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): FinalErasmusDataPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseFinalErasmusDataPacketData,
    } as FinalErasmusDataPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erasmusRestrictedInfo = reader.string();
          break;
        case 2:
          message.homeIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FinalErasmusDataPacketData {
    const message = {
      ...baseFinalErasmusDataPacketData,
    } as FinalErasmusDataPacketData;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = String(object.erasmusRestrictedInfo);
    } else {
      message.erasmusRestrictedInfo = "";
    }
    if (object.homeIndex !== undefined && object.homeIndex !== null) {
      message.homeIndex = String(object.homeIndex);
    } else {
      message.homeIndex = "";
    }
    return message;
  },

  toJSON(message: FinalErasmusDataPacketData): unknown {
    const obj: any = {};
    message.erasmusRestrictedInfo !== undefined &&
      (obj.erasmusRestrictedInfo = message.erasmusRestrictedInfo);
    message.homeIndex !== undefined && (obj.homeIndex = message.homeIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<FinalErasmusDataPacketData>
  ): FinalErasmusDataPacketData {
    const message = {
      ...baseFinalErasmusDataPacketData,
    } as FinalErasmusDataPacketData;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = object.erasmusRestrictedInfo;
    } else {
      message.erasmusRestrictedInfo = "";
    }
    if (object.homeIndex !== undefined && object.homeIndex !== null) {
      message.homeIndex = object.homeIndex;
    } else {
      message.homeIndex = "";
    }
    return message;
  },
};

const baseFinalErasmusDataPacketAck: object = {};

export const FinalErasmusDataPacketAck = {
  encode(
    _: FinalErasmusDataPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): FinalErasmusDataPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseFinalErasmusDataPacketAck,
    } as FinalErasmusDataPacketAck;
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

  fromJSON(_: any): FinalErasmusDataPacketAck {
    const message = {
      ...baseFinalErasmusDataPacketAck,
    } as FinalErasmusDataPacketAck;
    return message;
  },

  toJSON(_: FinalErasmusDataPacketAck): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<FinalErasmusDataPacketAck>
  ): FinalErasmusDataPacketAck {
    const message = {
      ...baseFinalErasmusDataPacketAck,
    } as FinalErasmusDataPacketAck;
    return message;
  },
};

const baseExtendErasmusPeriodPacketData: object = {
  durationInMonths: 0,
  destinationUniversityName: "",
  foreignIndex: "",
  finalDate: "",
};

export const ExtendErasmusPeriodPacketData = {
  encode(
    message: ExtendErasmusPeriodPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.durationInMonths !== 0) {
      writer.uint32(8).uint32(message.durationInMonths);
    }
    if (message.destinationUniversityName !== "") {
      writer.uint32(18).string(message.destinationUniversityName);
    }
    if (message.foreignIndex !== "") {
      writer.uint32(26).string(message.foreignIndex);
    }
    if (message.finalDate !== "") {
      writer.uint32(34).string(message.finalDate);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ExtendErasmusPeriodPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseExtendErasmusPeriodPacketData,
    } as ExtendErasmusPeriodPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.durationInMonths = reader.uint32();
          break;
        case 2:
          message.destinationUniversityName = reader.string();
          break;
        case 3:
          message.foreignIndex = reader.string();
          break;
        case 4:
          message.finalDate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ExtendErasmusPeriodPacketData {
    const message = {
      ...baseExtendErasmusPeriodPacketData,
    } as ExtendErasmusPeriodPacketData;
    if (
      object.durationInMonths !== undefined &&
      object.durationInMonths !== null
    ) {
      message.durationInMonths = Number(object.durationInMonths);
    } else {
      message.durationInMonths = 0;
    }
    if (
      object.destinationUniversityName !== undefined &&
      object.destinationUniversityName !== null
    ) {
      message.destinationUniversityName = String(
        object.destinationUniversityName
      );
    } else {
      message.destinationUniversityName = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = String(object.foreignIndex);
    } else {
      message.foreignIndex = "";
    }
    if (object.finalDate !== undefined && object.finalDate !== null) {
      message.finalDate = String(object.finalDate);
    } else {
      message.finalDate = "";
    }
    return message;
  },

  toJSON(message: ExtendErasmusPeriodPacketData): unknown {
    const obj: any = {};
    message.durationInMonths !== undefined &&
      (obj.durationInMonths = message.durationInMonths);
    message.destinationUniversityName !== undefined &&
      (obj.destinationUniversityName = message.destinationUniversityName);
    message.foreignIndex !== undefined &&
      (obj.foreignIndex = message.foreignIndex);
    message.finalDate !== undefined && (obj.finalDate = message.finalDate);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ExtendErasmusPeriodPacketData>
  ): ExtendErasmusPeriodPacketData {
    const message = {
      ...baseExtendErasmusPeriodPacketData,
    } as ExtendErasmusPeriodPacketData;
    if (
      object.durationInMonths !== undefined &&
      object.durationInMonths !== null
    ) {
      message.durationInMonths = object.durationInMonths;
    } else {
      message.durationInMonths = 0;
    }
    if (
      object.destinationUniversityName !== undefined &&
      object.destinationUniversityName !== null
    ) {
      message.destinationUniversityName = object.destinationUniversityName;
    } else {
      message.destinationUniversityName = "";
    }
    if (object.foreignIndex !== undefined && object.foreignIndex !== null) {
      message.foreignIndex = object.foreignIndex;
    } else {
      message.foreignIndex = "";
    }
    if (object.finalDate !== undefined && object.finalDate !== null) {
      message.finalDate = object.finalDate;
    } else {
      message.finalDate = "";
    }
    return message;
  },
};

const baseExtendErasmusPeriodPacketAck: object = {};

export const ExtendErasmusPeriodPacketAck = {
  encode(
    _: ExtendErasmusPeriodPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ExtendErasmusPeriodPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseExtendErasmusPeriodPacketAck,
    } as ExtendErasmusPeriodPacketAck;
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

  fromJSON(_: any): ExtendErasmusPeriodPacketAck {
    const message = {
      ...baseExtendErasmusPeriodPacketAck,
    } as ExtendErasmusPeriodPacketAck;
    return message;
  },

  toJSON(_: ExtendErasmusPeriodPacketAck): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<ExtendErasmusPeriodPacketAck>
  ): ExtendErasmusPeriodPacketAck {
    const message = {
      ...baseExtendErasmusPeriodPacketAck,
    } as ExtendErasmusPeriodPacketAck;
    return message;
  },
};

const baseErasmusRestictedDataPacketData: object = {
  erasmusRestrictedInfo: "",
};

export const ErasmusRestictedDataPacketData = {
  encode(
    message: ErasmusRestictedDataPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.erasmusRestrictedInfo !== "") {
      writer.uint32(10).string(message.erasmusRestrictedInfo);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ErasmusRestictedDataPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseErasmusRestictedDataPacketData,
    } as ErasmusRestictedDataPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erasmusRestrictedInfo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusRestictedDataPacketData {
    const message = {
      ...baseErasmusRestictedDataPacketData,
    } as ErasmusRestictedDataPacketData;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = String(object.erasmusRestrictedInfo);
    } else {
      message.erasmusRestrictedInfo = "";
    }
    return message;
  },

  toJSON(message: ErasmusRestictedDataPacketData): unknown {
    const obj: any = {};
    message.erasmusRestrictedInfo !== undefined &&
      (obj.erasmusRestrictedInfo = message.erasmusRestrictedInfo);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ErasmusRestictedDataPacketData>
  ): ErasmusRestictedDataPacketData {
    const message = {
      ...baseErasmusRestictedDataPacketData,
    } as ErasmusRestictedDataPacketData;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = object.erasmusRestrictedInfo;
    } else {
      message.erasmusRestrictedInfo = "";
    }
    return message;
  },
};

const baseErasmusRestictedDataPacketAck: object = { erasmusRestrictedInfo: "" };

export const ErasmusRestictedDataPacketAck = {
  encode(
    message: ErasmusRestictedDataPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.erasmusRestrictedInfo !== "") {
      writer.uint32(10).string(message.erasmusRestrictedInfo);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ErasmusRestictedDataPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseErasmusRestictedDataPacketAck,
    } as ErasmusRestictedDataPacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.erasmusRestrictedInfo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusRestictedDataPacketAck {
    const message = {
      ...baseErasmusRestictedDataPacketAck,
    } as ErasmusRestictedDataPacketAck;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = String(object.erasmusRestrictedInfo);
    } else {
      message.erasmusRestrictedInfo = "";
    }
    return message;
  },

  toJSON(message: ErasmusRestictedDataPacketAck): unknown {
    const obj: any = {};
    message.erasmusRestrictedInfo !== undefined &&
      (obj.erasmusRestrictedInfo = message.erasmusRestrictedInfo);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ErasmusRestictedDataPacketAck>
  ): ErasmusRestictedDataPacketAck {
    const message = {
      ...baseErasmusRestictedDataPacketAck,
    } as ErasmusRestictedDataPacketAck;
    if (
      object.erasmusRestrictedInfo !== undefined &&
      object.erasmusRestrictedInfo !== null
    ) {
      message.erasmusRestrictedInfo = object.erasmusRestrictedInfo;
    } else {
      message.erasmusRestrictedInfo = "";
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
