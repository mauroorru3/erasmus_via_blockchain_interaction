/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface CountersInfo {
  packetsRetriesStartErasmus: number[];
  acksReceivedStartErasmus: number;
  retryNumberOperations: number;
  maximumNumberRetries: number;
  errorAcksOrTimeoutsReceived: number;
  revertErasmusCareerCompleted: boolean;
}

const baseCountersInfo: object = {
  packetsRetriesStartErasmus: 0,
  acksReceivedStartErasmus: 0,
  retryNumberOperations: 0,
  maximumNumberRetries: 0,
  errorAcksOrTimeoutsReceived: 0,
  revertErasmusCareerCompleted: false,
};

export const CountersInfo = {
  encode(message: CountersInfo, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork();
    for (const v of message.packetsRetriesStartErasmus) {
      writer.int32(v);
    }
    writer.ldelim();
    if (message.acksReceivedStartErasmus !== 0) {
      writer.uint32(16).int32(message.acksReceivedStartErasmus);
    }
    if (message.retryNumberOperations !== 0) {
      writer.uint32(24).int32(message.retryNumberOperations);
    }
    if (message.maximumNumberRetries !== 0) {
      writer.uint32(32).int32(message.maximumNumberRetries);
    }
    if (message.errorAcksOrTimeoutsReceived !== 0) {
      writer.uint32(40).int32(message.errorAcksOrTimeoutsReceived);
    }
    if (message.revertErasmusCareerCompleted === true) {
      writer.uint32(48).bool(message.revertErasmusCareerCompleted);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CountersInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCountersInfo } as CountersInfo;
    message.packetsRetriesStartErasmus = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.packetsRetriesStartErasmus.push(reader.int32());
            }
          } else {
            message.packetsRetriesStartErasmus.push(reader.int32());
          }
          break;
        case 2:
          message.acksReceivedStartErasmus = reader.int32();
          break;
        case 3:
          message.retryNumberOperations = reader.int32();
          break;
        case 4:
          message.maximumNumberRetries = reader.int32();
          break;
        case 5:
          message.errorAcksOrTimeoutsReceived = reader.int32();
          break;
        case 6:
          message.revertErasmusCareerCompleted = reader.bool();
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
    message.packetsRetriesStartErasmus = [];
    if (
      object.packetsRetriesStartErasmus !== undefined &&
      object.packetsRetriesStartErasmus !== null
    ) {
      for (const e of object.packetsRetriesStartErasmus) {
        message.packetsRetriesStartErasmus.push(Number(e));
      }
    }
    if (
      object.acksReceivedStartErasmus !== undefined &&
      object.acksReceivedStartErasmus !== null
    ) {
      message.acksReceivedStartErasmus = Number(
        object.acksReceivedStartErasmus
      );
    } else {
      message.acksReceivedStartErasmus = 0;
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
    if (
      object.errorAcksOrTimeoutsReceived !== undefined &&
      object.errorAcksOrTimeoutsReceived !== null
    ) {
      message.errorAcksOrTimeoutsReceived = Number(
        object.errorAcksOrTimeoutsReceived
      );
    } else {
      message.errorAcksOrTimeoutsReceived = 0;
    }
    if (
      object.revertErasmusCareerCompleted !== undefined &&
      object.revertErasmusCareerCompleted !== null
    ) {
      message.revertErasmusCareerCompleted = Boolean(
        object.revertErasmusCareerCompleted
      );
    } else {
      message.revertErasmusCareerCompleted = false;
    }
    return message;
  },

  toJSON(message: CountersInfo): unknown {
    const obj: any = {};
    if (message.packetsRetriesStartErasmus) {
      obj.packetsRetriesStartErasmus = message.packetsRetriesStartErasmus.map(
        (e) => e
      );
    } else {
      obj.packetsRetriesStartErasmus = [];
    }
    message.acksReceivedStartErasmus !== undefined &&
      (obj.acksReceivedStartErasmus = message.acksReceivedStartErasmus);
    message.retryNumberOperations !== undefined &&
      (obj.retryNumberOperations = message.retryNumberOperations);
    message.maximumNumberRetries !== undefined &&
      (obj.maximumNumberRetries = message.maximumNumberRetries);
    message.errorAcksOrTimeoutsReceived !== undefined &&
      (obj.errorAcksOrTimeoutsReceived = message.errorAcksOrTimeoutsReceived);
    message.revertErasmusCareerCompleted !== undefined &&
      (obj.revertErasmusCareerCompleted = message.revertErasmusCareerCompleted);
    return obj;
  },

  fromPartial(object: DeepPartial<CountersInfo>): CountersInfo {
    const message = { ...baseCountersInfo } as CountersInfo;
    message.packetsRetriesStartErasmus = [];
    if (
      object.packetsRetriesStartErasmus !== undefined &&
      object.packetsRetriesStartErasmus !== null
    ) {
      for (const e of object.packetsRetriesStartErasmus) {
        message.packetsRetriesStartErasmus.push(e);
      }
    }
    if (
      object.acksReceivedStartErasmus !== undefined &&
      object.acksReceivedStartErasmus !== null
    ) {
      message.acksReceivedStartErasmus = object.acksReceivedStartErasmus;
    } else {
      message.acksReceivedStartErasmus = 0;
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
    if (
      object.errorAcksOrTimeoutsReceived !== undefined &&
      object.errorAcksOrTimeoutsReceived !== null
    ) {
      message.errorAcksOrTimeoutsReceived = object.errorAcksOrTimeoutsReceived;
    } else {
      message.errorAcksOrTimeoutsReceived = 0;
    }
    if (
      object.revertErasmusCareerCompleted !== undefined &&
      object.revertErasmusCareerCompleted !== null
    ) {
      message.revertErasmusCareerCompleted =
        object.revertErasmusCareerCompleted;
    } else {
      message.revertErasmusCareerCompleted = false;
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
