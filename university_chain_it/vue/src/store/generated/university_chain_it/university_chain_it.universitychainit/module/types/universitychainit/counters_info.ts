/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface CountersInfo {
  PacketsRetriesStartErasmus: number[];
  AcksReceivedStartErasmus: number;
  retryNumberOperations: number;
  maximumNumberRetries: number;
}

const baseCountersInfo: object = {
  PacketsRetriesStartErasmus: 0,
  AcksReceivedStartErasmus: 0,
  retryNumberOperations: 0,
  maximumNumberRetries: 0,
};

export const CountersInfo = {
  encode(message: CountersInfo, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork();
    for (const v of message.PacketsRetriesStartErasmus) {
      writer.int32(v);
    }
    writer.ldelim();
    if (message.AcksReceivedStartErasmus !== 0) {
      writer.uint32(16).int32(message.AcksReceivedStartErasmus);
    }
    if (message.retryNumberOperations !== 0) {
      writer.uint32(24).int32(message.retryNumberOperations);
    }
    if (message.maximumNumberRetries !== 0) {
      writer.uint32(32).int32(message.maximumNumberRetries);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CountersInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCountersInfo } as CountersInfo;
    message.PacketsRetriesStartErasmus = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.PacketsRetriesStartErasmus.push(reader.int32());
            }
          } else {
            message.PacketsRetriesStartErasmus.push(reader.int32());
          }
          break;
        case 2:
          message.AcksReceivedStartErasmus = reader.int32();
          break;
        case 3:
          message.retryNumberOperations = reader.int32();
          break;
        case 4:
          message.maximumNumberRetries = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CountersInfo {
    const message = { ...baseCountersInfo } as CountersInfo;
    message.PacketsRetriesStartErasmus = [];
    if (
      object.PacketsRetriesStartErasmus !== undefined &&
      object.PacketsRetriesStartErasmus !== null
    ) {
      for (const e of object.PacketsRetriesStartErasmus) {
        message.PacketsRetriesStartErasmus.push(Number(e));
      }
    }
    if (
      object.AcksReceivedStartErasmus !== undefined &&
      object.AcksReceivedStartErasmus !== null
    ) {
      message.AcksReceivedStartErasmus = Number(
        object.AcksReceivedStartErasmus
      );
    } else {
      message.AcksReceivedStartErasmus = 0;
    }
    if (
      object.retryNumberOperations !== undefined &&
      object.retryNumberOperations !== null
    ) {
      message.retryNumberOperations = Number(object.retryNumberOperations);
    } else {
      message.retryNumberOperations = 0;
    }
    if (
      object.maximumNumberRetries !== undefined &&
      object.maximumNumberRetries !== null
    ) {
      message.maximumNumberRetries = Number(object.maximumNumberRetries);
    } else {
      message.maximumNumberRetries = 0;
    }
    return message;
  },

  toJSON(message: CountersInfo): unknown {
    const obj: any = {};
    if (message.PacketsRetriesStartErasmus) {
      obj.PacketsRetriesStartErasmus = message.PacketsRetriesStartErasmus.map(
        (e) => e
      );
    } else {
      obj.PacketsRetriesStartErasmus = [];
    }
    message.AcksReceivedStartErasmus !== undefined &&
      (obj.AcksReceivedStartErasmus = message.AcksReceivedStartErasmus);
    message.retryNumberOperations !== undefined &&
      (obj.retryNumberOperations = message.retryNumberOperations);
    message.maximumNumberRetries !== undefined &&
      (obj.maximumNumberRetries = message.maximumNumberRetries);
    return obj;
  },

  fromPartial(object: DeepPartial<CountersInfo>): CountersInfo {
    const message = { ...baseCountersInfo } as CountersInfo;
    message.PacketsRetriesStartErasmus = [];
    if (
      object.PacketsRetriesStartErasmus !== undefined &&
      object.PacketsRetriesStartErasmus !== null
    ) {
      for (const e of object.PacketsRetriesStartErasmus) {
        message.PacketsRetriesStartErasmus.push(e);
      }
    }
    if (
      object.AcksReceivedStartErasmus !== undefined &&
      object.AcksReceivedStartErasmus !== null
    ) {
      message.AcksReceivedStartErasmus = object.AcksReceivedStartErasmus;
    } else {
      message.AcksReceivedStartErasmus = 0;
    }
    if (
      object.retryNumberOperations !== undefined &&
      object.retryNumberOperations !== null
    ) {
      message.retryNumberOperations = object.retryNumberOperations;
    } else {
      message.retryNumberOperations = 0;
    }
    if (
      object.maximumNumberRetries !== undefined &&
      object.maximumNumberRetries !== null
    ) {
      message.maximumNumberRetries = object.maximumNumberRetries;
    } else {
      message.maximumNumberRetries = 0;
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
