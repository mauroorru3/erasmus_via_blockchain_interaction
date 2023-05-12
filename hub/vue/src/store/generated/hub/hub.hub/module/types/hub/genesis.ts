/* eslint-disable */
import { Params } from "../hub/params";
import { StudentInfo } from "../hub/student_info";
import { TranscriptOfRecords } from "../hub/transcript_of_records";
import { PersonalInfo } from "../hub/personal_info";
import { ResidenceInfo } from "../hub/residence_info";
import { ContactInfo } from "../hub/contact_info";
import { TaxesInfo } from "../hub/taxes_info";
import { ErasmusInfo } from "../hub/erasmus_info";
import { StoredStudent } from "../hub/stored_student";
import { ChainInfo } from "../hub/chain_info";
import { Universities } from "../hub/universities";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

/** GenesisState defines the hub module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  port_id: string;
  studentInfo: StudentInfo | undefined;
  transcriptOfRecords: TranscriptOfRecords | undefined;
  personalInfo: PersonalInfo | undefined;
  residenceInfo: ResidenceInfo | undefined;
  contactInfo: ContactInfo | undefined;
  taxesInfo: TaxesInfo | undefined;
  erasmusInfo: ErasmusInfo | undefined;
  storedStudentList: StoredStudent[];
  chainInfo: ChainInfo | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  universitiesList: Universities[];
}

const baseGenesisState: object = { port_id: "" };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.port_id !== "") {
      writer.uint32(18).string(message.port_id);
    }
    if (message.studentInfo !== undefined) {
      StudentInfo.encode(
        message.studentInfo,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.transcriptOfRecords !== undefined) {
      TranscriptOfRecords.encode(
        message.transcriptOfRecords,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.personalInfo !== undefined) {
      PersonalInfo.encode(
        message.personalInfo,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.residenceInfo !== undefined) {
      ResidenceInfo.encode(
        message.residenceInfo,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.contactInfo !== undefined) {
      ContactInfo.encode(
        message.contactInfo,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.taxesInfo !== undefined) {
      TaxesInfo.encode(message.taxesInfo, writer.uint32(66).fork()).ldelim();
    }
    if (message.erasmusInfo !== undefined) {
      ErasmusInfo.encode(
        message.erasmusInfo,
        writer.uint32(74).fork()
      ).ldelim();
    }
    for (const v of message.storedStudentList) {
      StoredStudent.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    if (message.chainInfo !== undefined) {
      ChainInfo.encode(message.chainInfo, writer.uint32(90).fork()).ldelim();
    }
    for (const v of message.universitiesList) {
      Universities.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.storedStudentList = [];
    message.universitiesList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.port_id = reader.string();
          break;
        case 3:
          message.studentInfo = StudentInfo.decode(reader, reader.uint32());
          break;
        case 4:
          message.transcriptOfRecords = TranscriptOfRecords.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.personalInfo = PersonalInfo.decode(reader, reader.uint32());
          break;
        case 6:
          message.residenceInfo = ResidenceInfo.decode(reader, reader.uint32());
          break;
        case 7:
          message.contactInfo = ContactInfo.decode(reader, reader.uint32());
          break;
        case 8:
          message.taxesInfo = TaxesInfo.decode(reader, reader.uint32());
          break;
        case 9:
          message.erasmusInfo = ErasmusInfo.decode(reader, reader.uint32());
          break;
        case 10:
          message.storedStudentList.push(
            StoredStudent.decode(reader, reader.uint32())
          );
          break;
        case 11:
          message.chainInfo = ChainInfo.decode(reader, reader.uint32());
          break;
        case 12:
          message.universitiesList.push(
            Universities.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedStudentList = [];
    message.universitiesList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = String(object.port_id);
    } else {
      message.port_id = "";
    }
    if (object.studentInfo !== undefined && object.studentInfo !== null) {
      message.studentInfo = StudentInfo.fromJSON(object.studentInfo);
    } else {
      message.studentInfo = undefined;
    }
    if (
      object.transcriptOfRecords !== undefined &&
      object.transcriptOfRecords !== null
    ) {
      message.transcriptOfRecords = TranscriptOfRecords.fromJSON(
        object.transcriptOfRecords
      );
    } else {
      message.transcriptOfRecords = undefined;
    }
    if (object.personalInfo !== undefined && object.personalInfo !== null) {
      message.personalInfo = PersonalInfo.fromJSON(object.personalInfo);
    } else {
      message.personalInfo = undefined;
    }
    if (object.residenceInfo !== undefined && object.residenceInfo !== null) {
      message.residenceInfo = ResidenceInfo.fromJSON(object.residenceInfo);
    } else {
      message.residenceInfo = undefined;
    }
    if (object.contactInfo !== undefined && object.contactInfo !== null) {
      message.contactInfo = ContactInfo.fromJSON(object.contactInfo);
    } else {
      message.contactInfo = undefined;
    }
    if (object.taxesInfo !== undefined && object.taxesInfo !== null) {
      message.taxesInfo = TaxesInfo.fromJSON(object.taxesInfo);
    } else {
      message.taxesInfo = undefined;
    }
    if (object.erasmusInfo !== undefined && object.erasmusInfo !== null) {
      message.erasmusInfo = ErasmusInfo.fromJSON(object.erasmusInfo);
    } else {
      message.erasmusInfo = undefined;
    }
    if (
      object.storedStudentList !== undefined &&
      object.storedStudentList !== null
    ) {
      for (const e of object.storedStudentList) {
        message.storedStudentList.push(StoredStudent.fromJSON(e));
      }
    }
    if (object.chainInfo !== undefined && object.chainInfo !== null) {
      message.chainInfo = ChainInfo.fromJSON(object.chainInfo);
    } else {
      message.chainInfo = undefined;
    }
    if (
      object.universitiesList !== undefined &&
      object.universitiesList !== null
    ) {
      for (const e of object.universitiesList) {
        message.universitiesList.push(Universities.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.port_id !== undefined && (obj.port_id = message.port_id);
    message.studentInfo !== undefined &&
      (obj.studentInfo = message.studentInfo
        ? StudentInfo.toJSON(message.studentInfo)
        : undefined);
    message.transcriptOfRecords !== undefined &&
      (obj.transcriptOfRecords = message.transcriptOfRecords
        ? TranscriptOfRecords.toJSON(message.transcriptOfRecords)
        : undefined);
    message.personalInfo !== undefined &&
      (obj.personalInfo = message.personalInfo
        ? PersonalInfo.toJSON(message.personalInfo)
        : undefined);
    message.residenceInfo !== undefined &&
      (obj.residenceInfo = message.residenceInfo
        ? ResidenceInfo.toJSON(message.residenceInfo)
        : undefined);
    message.contactInfo !== undefined &&
      (obj.contactInfo = message.contactInfo
        ? ContactInfo.toJSON(message.contactInfo)
        : undefined);
    message.taxesInfo !== undefined &&
      (obj.taxesInfo = message.taxesInfo
        ? TaxesInfo.toJSON(message.taxesInfo)
        : undefined);
    message.erasmusInfo !== undefined &&
      (obj.erasmusInfo = message.erasmusInfo
        ? ErasmusInfo.toJSON(message.erasmusInfo)
        : undefined);
    if (message.storedStudentList) {
      obj.storedStudentList = message.storedStudentList.map((e) =>
        e ? StoredStudent.toJSON(e) : undefined
      );
    } else {
      obj.storedStudentList = [];
    }
    message.chainInfo !== undefined &&
      (obj.chainInfo = message.chainInfo
        ? ChainInfo.toJSON(message.chainInfo)
        : undefined);
    if (message.universitiesList) {
      obj.universitiesList = message.universitiesList.map((e) =>
        e ? Universities.toJSON(e) : undefined
      );
    } else {
      obj.universitiesList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedStudentList = [];
    message.universitiesList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = object.port_id;
    } else {
      message.port_id = "";
    }
    if (object.studentInfo !== undefined && object.studentInfo !== null) {
      message.studentInfo = StudentInfo.fromPartial(object.studentInfo);
    } else {
      message.studentInfo = undefined;
    }
    if (
      object.transcriptOfRecords !== undefined &&
      object.transcriptOfRecords !== null
    ) {
      message.transcriptOfRecords = TranscriptOfRecords.fromPartial(
        object.transcriptOfRecords
      );
    } else {
      message.transcriptOfRecords = undefined;
    }
    if (object.personalInfo !== undefined && object.personalInfo !== null) {
      message.personalInfo = PersonalInfo.fromPartial(object.personalInfo);
    } else {
      message.personalInfo = undefined;
    }
    if (object.residenceInfo !== undefined && object.residenceInfo !== null) {
      message.residenceInfo = ResidenceInfo.fromPartial(object.residenceInfo);
    } else {
      message.residenceInfo = undefined;
    }
    if (object.contactInfo !== undefined && object.contactInfo !== null) {
      message.contactInfo = ContactInfo.fromPartial(object.contactInfo);
    } else {
      message.contactInfo = undefined;
    }
    if (object.taxesInfo !== undefined && object.taxesInfo !== null) {
      message.taxesInfo = TaxesInfo.fromPartial(object.taxesInfo);
    } else {
      message.taxesInfo = undefined;
    }
    if (object.erasmusInfo !== undefined && object.erasmusInfo !== null) {
      message.erasmusInfo = ErasmusInfo.fromPartial(object.erasmusInfo);
    } else {
      message.erasmusInfo = undefined;
    }
    if (
      object.storedStudentList !== undefined &&
      object.storedStudentList !== null
    ) {
      for (const e of object.storedStudentList) {
        message.storedStudentList.push(StoredStudent.fromPartial(e));
      }
    }
    if (object.chainInfo !== undefined && object.chainInfo !== null) {
      message.chainInfo = ChainInfo.fromPartial(object.chainInfo);
    } else {
      message.chainInfo = undefined;
    }
    if (
      object.universitiesList !== undefined &&
      object.universitiesList !== null
    ) {
      for (const e of object.universitiesList) {
        message.universitiesList.push(Universities.fromPartial(e));
      }
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
