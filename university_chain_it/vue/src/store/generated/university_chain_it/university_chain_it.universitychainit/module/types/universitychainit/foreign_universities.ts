/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface ForeignUniversities {
  universityName: string;
  foreignUniversitiesKey: string;
  foreignUniversitiesCountry: string;
}

const baseForeignUniversities: object = {
  universityName: "",
  foreignUniversitiesKey: "",
  foreignUniversitiesCountry: "",
};

export const ForeignUniversities = {
  encode(
    message: ForeignUniversities,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.universityName !== "") {
      writer.uint32(10).string(message.universityName);
    }
    if (message.foreignUniversitiesKey !== "") {
      writer.uint32(18).string(message.foreignUniversitiesKey);
    }
    if (message.foreignUniversitiesCountry !== "") {
      writer.uint32(26).string(message.foreignUniversitiesCountry);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ForeignUniversities {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseForeignUniversities } as ForeignUniversities;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityName = reader.string();
          break;
        case 2:
          message.foreignUniversitiesKey = reader.string();
          break;
        case 3:
          message.foreignUniversitiesCountry = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ForeignUniversities {
    const message = { ...baseForeignUniversities } as ForeignUniversities;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = String(object.universityName);
    } else {
      message.universityName = "";
    }
    if (
      object.foreignUniversitiesKey !== undefined &&
      object.foreignUniversitiesKey !== null
    ) {
      message.foreignUniversitiesKey = String(object.foreignUniversitiesKey);
    } else {
      message.foreignUniversitiesKey = "";
    }
    if (
      object.foreignUniversitiesCountry !== undefined &&
      object.foreignUniversitiesCountry !== null
    ) {
      message.foreignUniversitiesCountry = String(
        object.foreignUniversitiesCountry
      );
    } else {
      message.foreignUniversitiesCountry = "";
    }
    return message;
  },

  toJSON(message: ForeignUniversities): unknown {
    const obj: any = {};
    message.universityName !== undefined &&
      (obj.universityName = message.universityName);
    message.foreignUniversitiesKey !== undefined &&
      (obj.foreignUniversitiesKey = message.foreignUniversitiesKey);
    message.foreignUniversitiesCountry !== undefined &&
      (obj.foreignUniversitiesCountry = message.foreignUniversitiesCountry);
    return obj;
  },

  fromPartial(object: DeepPartial<ForeignUniversities>): ForeignUniversities {
    const message = { ...baseForeignUniversities } as ForeignUniversities;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = object.universityName;
    } else {
      message.universityName = "";
    }
    if (
      object.foreignUniversitiesKey !== undefined &&
      object.foreignUniversitiesKey !== null
    ) {
      message.foreignUniversitiesKey = object.foreignUniversitiesKey;
    } else {
      message.foreignUniversitiesKey = "";
    }
    if (
      object.foreignUniversitiesCountry !== undefined &&
      object.foreignUniversitiesCountry !== null
    ) {
      message.foreignUniversitiesCountry = object.foreignUniversitiesCountry;
    } else {
      message.foreignUniversitiesCountry = "";
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
