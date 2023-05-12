/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface TaxesInfo {
  status: boolean;
  totalAmount: number;
  taxesHistory: string;
}

const baseTaxesInfo: object = {
  status: false,
  totalAmount: 0,
  taxesHistory: "",
};

export const TaxesInfo = {
  encode(message: TaxesInfo, writer: Writer = Writer.create()): Writer {
    if (message.status === true) {
      writer.uint32(8).bool(message.status);
    }
    if (message.totalAmount !== 0) {
      writer.uint32(16).uint32(message.totalAmount);
    }
    if (message.taxesHistory !== "") {
      writer.uint32(26).string(message.taxesHistory);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TaxesInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTaxesInfo } as TaxesInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.bool();
          break;
        case 2:
          message.totalAmount = reader.uint32();
          break;
        case 3:
          message.taxesHistory = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TaxesInfo {
    const message = { ...baseTaxesInfo } as TaxesInfo;
    if (object.status !== undefined && object.status !== null) {
      message.status = Boolean(object.status);
    } else {
      message.status = false;
    }
    if (object.totalAmount !== undefined && object.totalAmount !== null) {
      message.totalAmount = Number(object.totalAmount);
    } else {
      message.totalAmount = 0;
    }
    if (object.taxesHistory !== undefined && object.taxesHistory !== null) {
      message.taxesHistory = String(object.taxesHistory);
    } else {
      message.taxesHistory = "";
    }
    return message;
  },

  toJSON(message: TaxesInfo): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    message.totalAmount !== undefined &&
      (obj.totalAmount = message.totalAmount);
    message.taxesHistory !== undefined &&
      (obj.taxesHistory = message.taxesHistory);
    return obj;
  },

  fromPartial(object: DeepPartial<TaxesInfo>): TaxesInfo {
    const message = { ...baseTaxesInfo } as TaxesInfo;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = false;
    }
    if (object.totalAmount !== undefined && object.totalAmount !== null) {
      message.totalAmount = object.totalAmount;
    } else {
      message.totalAmount = 0;
    }
    if (object.taxesHistory !== undefined && object.taxesHistory !== null) {
      message.taxesHistory = object.taxesHistory;
    } else {
      message.taxesHistory = "";
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
