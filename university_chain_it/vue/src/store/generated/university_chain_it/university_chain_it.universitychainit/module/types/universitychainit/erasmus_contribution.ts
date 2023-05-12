/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface ErasmusContribution {
  amount: number;
  incomeBracket: number;
  paymentMade: boolean;
  dateOfPayment: string;
}

const baseErasmusContribution: object = {
  amount: 0,
  incomeBracket: 0,
  paymentMade: false,
  dateOfPayment: "",
};

export const ErasmusContribution = {
  encode(
    message: ErasmusContribution,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.amount !== 0) {
      writer.uint32(8).uint64(message.amount);
    }
    if (message.incomeBracket !== 0) {
      writer.uint32(16).uint64(message.incomeBracket);
    }
    if (message.paymentMade === true) {
      writer.uint32(24).bool(message.paymentMade);
    }
    if (message.dateOfPayment !== "") {
      writer.uint32(34).string(message.dateOfPayment);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ErasmusContribution {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseErasmusContribution } as ErasmusContribution;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.incomeBracket = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.paymentMade = reader.bool();
          break;
        case 4:
          message.dateOfPayment = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ErasmusContribution {
    const message = { ...baseErasmusContribution } as ErasmusContribution;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.incomeBracket !== undefined && object.incomeBracket !== null) {
      message.incomeBracket = Number(object.incomeBracket);
    } else {
      message.incomeBracket = 0;
    }
    if (object.paymentMade !== undefined && object.paymentMade !== null) {
      message.paymentMade = Boolean(object.paymentMade);
    } else {
      message.paymentMade = false;
    }
    if (object.dateOfPayment !== undefined && object.dateOfPayment !== null) {
      message.dateOfPayment = String(object.dateOfPayment);
    } else {
      message.dateOfPayment = "";
    }
    return message;
  },

  toJSON(message: ErasmusContribution): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    message.incomeBracket !== undefined &&
      (obj.incomeBracket = message.incomeBracket);
    message.paymentMade !== undefined &&
      (obj.paymentMade = message.paymentMade);
    message.dateOfPayment !== undefined &&
      (obj.dateOfPayment = message.dateOfPayment);
    return obj;
  },

  fromPartial(object: DeepPartial<ErasmusContribution>): ErasmusContribution {
    const message = { ...baseErasmusContribution } as ErasmusContribution;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.incomeBracket !== undefined && object.incomeBracket !== null) {
      message.incomeBracket = object.incomeBracket;
    } else {
      message.incomeBracket = 0;
    }
    if (object.paymentMade !== undefined && object.paymentMade !== null) {
      message.paymentMade = object.paymentMade;
    } else {
      message.paymentMade = false;
    }
    if (object.dateOfPayment !== undefined && object.dateOfPayment !== null) {
      message.dateOfPayment = object.dateOfPayment;
    } else {
      message.dateOfPayment = "";
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
