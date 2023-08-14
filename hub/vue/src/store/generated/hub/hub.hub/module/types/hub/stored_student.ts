/* eslint-disable */
import { StudentInfo } from "../hub/student_info";
import { TranscriptOfRecords } from "../hub/transcript_of_records";
import { PersonalInfo } from "../hub/personal_info";
import { ResidenceInfo } from "../hub/residence_info";
import { ContactInfo } from "../hub/contact_info";
import { TaxesInfo } from "../hub/taxes_info";
import { ErasmusInfo } from "../hub/erasmus_info";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface StoredStudent {
  index: string;
  studentData: StudentInfo | undefined;
  transcriptData: TranscriptOfRecords | undefined;
  personalData: PersonalInfo | undefined;
  residenceData: ResidenceInfo | undefined;
  contactData: ContactInfo | undefined;
  taxesData: TaxesInfo | undefined;
  erasmusData: ErasmusInfo | undefined;
}

const baseStoredStudent: object = { index: "" };

export const StoredStudent = {
  encode(message: StoredStudent, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.studentData !== undefined) {
      StudentInfo.encode(
        message.studentData,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.transcriptData !== undefined) {
      TranscriptOfRecords.encode(
        message.transcriptData,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.personalData !== undefined) {
      PersonalInfo.encode(
        message.personalData,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.residenceData !== undefined) {
      ResidenceInfo.encode(
        message.residenceData,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.contactData !== undefined) {
      ContactInfo.encode(
        message.contactData,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.taxesData !== undefined) {
      TaxesInfo.encode(message.taxesData, writer.uint32(58).fork()).ldelim();
    }
    if (message.erasmusData !== undefined) {
      ErasmusInfo.encode(
        message.erasmusData,
        writer.uint32(66).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredStudent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredStudent } as StoredStudent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.studentData = StudentInfo.decode(reader, reader.uint32());
          break;
        case 3:
          message.transcriptData = TranscriptOfRecords.decode(
            reader,
            reader.uint32()
          );
          break;
        case 4:
          message.personalData = PersonalInfo.decode(reader, reader.uint32());
          break;
        case 5:
          message.residenceData = ResidenceInfo.decode(reader, reader.uint32());
          break;
        case 6:
          message.contactData = ContactInfo.decode(reader, reader.uint32());
          break;
        case 7:
          message.taxesData = TaxesInfo.decode(reader, reader.uint32());
          break;
        case 8:
          message.erasmusData = ErasmusInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredStudent {
    const message = { ...baseStoredStudent } as StoredStudent;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.studentData !== undefined && object.studentData !== null) {
      message.studentData = StudentInfo.fromJSON(object.studentData);
    } else {
      message.studentData = undefined;
    }
    if (object.transcriptData !== undefined && object.transcriptData !== null) {
      message.transcriptData = TranscriptOfRecords.fromJSON(
        object.transcriptData
      );
    } else {
      message.transcriptData = undefined;
    }
    if (object.personalData !== undefined && object.personalData !== null) {
      message.personalData = PersonalInfo.fromJSON(object.personalData);
    } else {
      message.personalData = undefined;
    }
    if (object.residenceData !== undefined && object.residenceData !== null) {
      message.residenceData = ResidenceInfo.fromJSON(object.residenceData);
    } else {
      message.residenceData = undefined;
    }
    if (object.contactData !== undefined && object.contactData !== null) {
      message.contactData = ContactInfo.fromJSON(object.contactData);
    } else {
      message.contactData = undefined;
    }
    if (object.taxesData !== undefined && object.taxesData !== null) {
      message.taxesData = TaxesInfo.fromJSON(object.taxesData);
    } else {
      message.taxesData = undefined;
    }
    if (object.erasmusData !== undefined && object.erasmusData !== null) {
      message.erasmusData = ErasmusInfo.fromJSON(object.erasmusData);
    } else {
      message.erasmusData = undefined;
    }
    return message;
  },

  toJSON(message: StoredStudent): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.studentData !== undefined &&
      (obj.studentData = message.studentData
        ? StudentInfo.toJSON(message.studentData)
        : undefined);
    message.transcriptData !== undefined &&
      (obj.transcriptData = message.transcriptData
        ? TranscriptOfRecords.toJSON(message.transcriptData)
        : undefined);
    message.personalData !== undefined &&
      (obj.personalData = message.personalData
        ? PersonalInfo.toJSON(message.personalData)
        : undefined);
    message.residenceData !== undefined &&
      (obj.residenceData = message.residenceData
        ? ResidenceInfo.toJSON(message.residenceData)
        : undefined);
    message.contactData !== undefined &&
      (obj.contactData = message.contactData
        ? ContactInfo.toJSON(message.contactData)
        : undefined);
    message.taxesData !== undefined &&
      (obj.taxesData = message.taxesData
        ? TaxesInfo.toJSON(message.taxesData)
        : undefined);
    message.erasmusData !== undefined &&
      (obj.erasmusData = message.erasmusData
        ? ErasmusInfo.toJSON(message.erasmusData)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredStudent>): StoredStudent {
    const message = { ...baseStoredStudent } as StoredStudent;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.studentData !== undefined && object.studentData !== null) {
      message.studentData = StudentInfo.fromPartial(object.studentData);
    } else {
      message.studentData = undefined;
    }
    if (object.transcriptData !== undefined && object.transcriptData !== null) {
      message.transcriptData = TranscriptOfRecords.fromPartial(
        object.transcriptData
      );
    } else {
      message.transcriptData = undefined;
    }
    if (object.personalData !== undefined && object.personalData !== null) {
      message.personalData = PersonalInfo.fromPartial(object.personalData);
    } else {
      message.personalData = undefined;
    }
    if (object.residenceData !== undefined && object.residenceData !== null) {
      message.residenceData = ResidenceInfo.fromPartial(object.residenceData);
    } else {
      message.residenceData = undefined;
    }
    if (object.contactData !== undefined && object.contactData !== null) {
      message.contactData = ContactInfo.fromPartial(object.contactData);
    } else {
      message.contactData = undefined;
    }
    if (object.taxesData !== undefined && object.taxesData !== null) {
      message.taxesData = TaxesInfo.fromPartial(object.taxesData);
    } else {
      message.taxesData = undefined;
    }
    if (object.erasmusData !== undefined && object.erasmusData !== null) {
      message.erasmusData = ErasmusInfo.fromPartial(object.erasmusData);
    } else {
      message.erasmusData = undefined;
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
