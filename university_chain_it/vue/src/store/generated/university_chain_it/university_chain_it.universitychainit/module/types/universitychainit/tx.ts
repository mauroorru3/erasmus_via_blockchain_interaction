/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "university_chain_it.universitychainit";

export interface MsgConfigureChain {
  creator: string;
}

export interface MsgConfigureChainResponse {
  status: number;
}

export interface MsgRegisterNewStudent {
  creator: string;
  university: string;
  name: string;
  surname: string;
  courseType: string;
  courseOfStudy: string;
  departmentName: string;
}

export interface MsgRegisterNewStudentResponse {
  studentIndex: string;
}

export interface MsgInsertStudentPersonalInfo {
  creator: string;
  university: string;
  studentIndex: string;
  gender: string;
  dateOfBirth: string;
  primaryNationality: string;
  countryOfBirth: string;
  provinceOfBirth: string;
  townOfBirth: string;
  taxCode: string;
  incomeBracket: number;
}

export interface MsgInsertStudentPersonalInfoResponse {
  status: number;
}

export interface MsgInsertStudentContactInfo {
  creator: string;
  university: string;
  studentIndex: string;
  contactAddress: string;
  email: string;
  mobilePhone: string;
}

export interface MsgInsertStudentContactInfoResponse {
  status: number;
}

export interface MsgInsertStudentResidenceInfo {
  creator: string;
  university: string;
  studentIndex: string;
  country: string;
  province: string;
  town: string;
  postCode: string;
  address: string;
  houseNumber: string;
  homePhone: string;
}

export interface MsgInsertStudentResidenceInfoResponse {
  status: number;
}

export interface MsgInsertExamGrade {
  creator: string;
  university: string;
  studentIndex: string;
  examName: string;
  grade: string;
}

export interface MsgInsertExamGradeResponse {
  status: number;
}

export interface MsgPayTaxes {
  creator: string;
  university: string;
  studentIndex: string;
}

export interface MsgPayTaxesResponse {
  status: number;
}

export interface MsgInsertErasmusRequest {
  creator: string;
  university: string;
  studentIndex: string;
  durationInMonths: string;
  foreignUniversityName: string;
  erasmusType: string;
}

export interface MsgInsertErasmusRequestResponse {
  status: number;
}

export interface MsgInsertErasmusExam {
  creator: string;
  university: string;
  studentIndex: string;
  examName: string;
}

export interface MsgInsertErasmusExamResponse {
  status: number;
}

export interface MsgStartErasmus {
  creator: string;
  university: string;
  studentIndex: string;
}

export interface MsgStartErasmusResponse {
  status: number;
}

export interface MsgSendErasmusStudent {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  index: string;
}

export interface MsgSendErasmusStudentResponse {
  status: number;
}

export interface MsgSendEndErasmusPeriodRequest {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  startingUniversityName: string;
  destinationUniversityName: string;
  index: string;
  foreignIndex: string;
}

export interface MsgSendEndErasmusPeriodRequestResponse {
  status: number;
}

const baseMsgConfigureChain: object = { creator: "" };

export const MsgConfigureChain = {
  encode(message: MsgConfigureChain, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConfigureChain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgConfigureChain } as MsgConfigureChain;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfigureChain {
    const message = { ...baseMsgConfigureChain } as MsgConfigureChain;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgConfigureChain): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgConfigureChain>): MsgConfigureChain {
    const message = { ...baseMsgConfigureChain } as MsgConfigureChain;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgConfigureChainResponse: object = { status: 0 };

export const MsgConfigureChainResponse = {
  encode(
    message: MsgConfigureChainResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgConfigureChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgConfigureChainResponse,
    } as MsgConfigureChainResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfigureChainResponse {
    const message = {
      ...baseMsgConfigureChainResponse,
    } as MsgConfigureChainResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgConfigureChainResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgConfigureChainResponse>
  ): MsgConfigureChainResponse {
    const message = {
      ...baseMsgConfigureChainResponse,
    } as MsgConfigureChainResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgRegisterNewStudent: object = {
  creator: "",
  university: "",
  name: "",
  surname: "",
  courseType: "",
  courseOfStudy: "",
  departmentName: "",
};

export const MsgRegisterNewStudent = {
  encode(
    message: MsgRegisterNewStudent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.surname !== "") {
      writer.uint32(34).string(message.surname);
    }
    if (message.courseType !== "") {
      writer.uint32(42).string(message.courseType);
    }
    if (message.courseOfStudy !== "") {
      writer.uint32(50).string(message.courseOfStudy);
    }
    if (message.departmentName !== "") {
      writer.uint32(58).string(message.departmentName);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterNewStudent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterNewStudent } as MsgRegisterNewStudent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.surname = reader.string();
          break;
        case 5:
          message.courseType = reader.string();
          break;
        case 6:
          message.courseOfStudy = reader.string();
          break;
        case 7:
          message.departmentName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterNewStudent {
    const message = { ...baseMsgRegisterNewStudent } as MsgRegisterNewStudent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.surname !== undefined && object.surname !== null) {
      message.surname = String(object.surname);
    } else {
      message.surname = "";
    }
    if (object.courseType !== undefined && object.courseType !== null) {
      message.courseType = String(object.courseType);
    } else {
      message.courseType = "";
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = String(object.courseOfStudy);
    } else {
      message.courseOfStudy = "";
    }
    if (object.departmentName !== undefined && object.departmentName !== null) {
      message.departmentName = String(object.departmentName);
    } else {
      message.departmentName = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterNewStudent): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.name !== undefined && (obj.name = message.name);
    message.surname !== undefined && (obj.surname = message.surname);
    message.courseType !== undefined && (obj.courseType = message.courseType);
    message.courseOfStudy !== undefined &&
      (obj.courseOfStudy = message.courseOfStudy);
    message.departmentName !== undefined &&
      (obj.departmentName = message.departmentName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterNewStudent>
  ): MsgRegisterNewStudent {
    const message = { ...baseMsgRegisterNewStudent } as MsgRegisterNewStudent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.surname !== undefined && object.surname !== null) {
      message.surname = object.surname;
    } else {
      message.surname = "";
    }
    if (object.courseType !== undefined && object.courseType !== null) {
      message.courseType = object.courseType;
    } else {
      message.courseType = "";
    }
    if (object.courseOfStudy !== undefined && object.courseOfStudy !== null) {
      message.courseOfStudy = object.courseOfStudy;
    } else {
      message.courseOfStudy = "";
    }
    if (object.departmentName !== undefined && object.departmentName !== null) {
      message.departmentName = object.departmentName;
    } else {
      message.departmentName = "";
    }
    return message;
  },
};

const baseMsgRegisterNewStudentResponse: object = { studentIndex: "" };

export const MsgRegisterNewStudentResponse = {
  encode(
    message: MsgRegisterNewStudentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.studentIndex !== "") {
      writer.uint32(10).string(message.studentIndex);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterNewStudentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterNewStudentResponse,
    } as MsgRegisterNewStudentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.studentIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterNewStudentResponse {
    const message = {
      ...baseMsgRegisterNewStudentResponse,
    } as MsgRegisterNewStudentResponse;
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterNewStudentResponse): unknown {
    const obj: any = {};
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterNewStudentResponse>
  ): MsgRegisterNewStudentResponse {
    const message = {
      ...baseMsgRegisterNewStudentResponse,
    } as MsgRegisterNewStudentResponse;
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    return message;
  },
};

const baseMsgInsertStudentPersonalInfo: object = {
  creator: "",
  university: "",
  studentIndex: "",
  gender: "",
  dateOfBirth: "",
  primaryNationality: "",
  countryOfBirth: "",
  provinceOfBirth: "",
  townOfBirth: "",
  taxCode: "",
  incomeBracket: 0,
};

export const MsgInsertStudentPersonalInfo = {
  encode(
    message: MsgInsertStudentPersonalInfo,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    if (message.gender !== "") {
      writer.uint32(34).string(message.gender);
    }
    if (message.dateOfBirth !== "") {
      writer.uint32(42).string(message.dateOfBirth);
    }
    if (message.primaryNationality !== "") {
      writer.uint32(50).string(message.primaryNationality);
    }
    if (message.countryOfBirth !== "") {
      writer.uint32(58).string(message.countryOfBirth);
    }
    if (message.provinceOfBirth !== "") {
      writer.uint32(66).string(message.provinceOfBirth);
    }
    if (message.townOfBirth !== "") {
      writer.uint32(74).string(message.townOfBirth);
    }
    if (message.taxCode !== "") {
      writer.uint32(82).string(message.taxCode);
    }
    if (message.incomeBracket !== 0) {
      writer.uint32(88).uint32(message.incomeBracket);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertStudentPersonalInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertStudentPersonalInfo,
    } as MsgInsertStudentPersonalInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        case 4:
          message.gender = reader.string();
          break;
        case 5:
          message.dateOfBirth = reader.string();
          break;
        case 6:
          message.primaryNationality = reader.string();
          break;
        case 7:
          message.countryOfBirth = reader.string();
          break;
        case 8:
          message.provinceOfBirth = reader.string();
          break;
        case 9:
          message.townOfBirth = reader.string();
          break;
        case 10:
          message.taxCode = reader.string();
          break;
        case 11:
          message.incomeBracket = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertStudentPersonalInfo {
    const message = {
      ...baseMsgInsertStudentPersonalInfo,
    } as MsgInsertStudentPersonalInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = String(object.gender);
    } else {
      message.gender = "";
    }
    if (object.dateOfBirth !== undefined && object.dateOfBirth !== null) {
      message.dateOfBirth = String(object.dateOfBirth);
    } else {
      message.dateOfBirth = "";
    }
    if (
      object.primaryNationality !== undefined &&
      object.primaryNationality !== null
    ) {
      message.primaryNationality = String(object.primaryNationality);
    } else {
      message.primaryNationality = "";
    }
    if (object.countryOfBirth !== undefined && object.countryOfBirth !== null) {
      message.countryOfBirth = String(object.countryOfBirth);
    } else {
      message.countryOfBirth = "";
    }
    if (
      object.provinceOfBirth !== undefined &&
      object.provinceOfBirth !== null
    ) {
      message.provinceOfBirth = String(object.provinceOfBirth);
    } else {
      message.provinceOfBirth = "";
    }
    if (object.townOfBirth !== undefined && object.townOfBirth !== null) {
      message.townOfBirth = String(object.townOfBirth);
    } else {
      message.townOfBirth = "";
    }
    if (object.taxCode !== undefined && object.taxCode !== null) {
      message.taxCode = String(object.taxCode);
    } else {
      message.taxCode = "";
    }
    if (object.incomeBracket !== undefined && object.incomeBracket !== null) {
      message.incomeBracket = Number(object.incomeBracket);
    } else {
      message.incomeBracket = 0;
    }
    return message;
  },

  toJSON(message: MsgInsertStudentPersonalInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    message.gender !== undefined && (obj.gender = message.gender);
    message.dateOfBirth !== undefined &&
      (obj.dateOfBirth = message.dateOfBirth);
    message.primaryNationality !== undefined &&
      (obj.primaryNationality = message.primaryNationality);
    message.countryOfBirth !== undefined &&
      (obj.countryOfBirth = message.countryOfBirth);
    message.provinceOfBirth !== undefined &&
      (obj.provinceOfBirth = message.provinceOfBirth);
    message.townOfBirth !== undefined &&
      (obj.townOfBirth = message.townOfBirth);
    message.taxCode !== undefined && (obj.taxCode = message.taxCode);
    message.incomeBracket !== undefined &&
      (obj.incomeBracket = message.incomeBracket);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertStudentPersonalInfo>
  ): MsgInsertStudentPersonalInfo {
    const message = {
      ...baseMsgInsertStudentPersonalInfo,
    } as MsgInsertStudentPersonalInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    if (object.gender !== undefined && object.gender !== null) {
      message.gender = object.gender;
    } else {
      message.gender = "";
    }
    if (object.dateOfBirth !== undefined && object.dateOfBirth !== null) {
      message.dateOfBirth = object.dateOfBirth;
    } else {
      message.dateOfBirth = "";
    }
    if (
      object.primaryNationality !== undefined &&
      object.primaryNationality !== null
    ) {
      message.primaryNationality = object.primaryNationality;
    } else {
      message.primaryNationality = "";
    }
    if (object.countryOfBirth !== undefined && object.countryOfBirth !== null) {
      message.countryOfBirth = object.countryOfBirth;
    } else {
      message.countryOfBirth = "";
    }
    if (
      object.provinceOfBirth !== undefined &&
      object.provinceOfBirth !== null
    ) {
      message.provinceOfBirth = object.provinceOfBirth;
    } else {
      message.provinceOfBirth = "";
    }
    if (object.townOfBirth !== undefined && object.townOfBirth !== null) {
      message.townOfBirth = object.townOfBirth;
    } else {
      message.townOfBirth = "";
    }
    if (object.taxCode !== undefined && object.taxCode !== null) {
      message.taxCode = object.taxCode;
    } else {
      message.taxCode = "";
    }
    if (object.incomeBracket !== undefined && object.incomeBracket !== null) {
      message.incomeBracket = object.incomeBracket;
    } else {
      message.incomeBracket = 0;
    }
    return message;
  },
};

const baseMsgInsertStudentPersonalInfoResponse: object = { status: 0 };

export const MsgInsertStudentPersonalInfoResponse = {
  encode(
    message: MsgInsertStudentPersonalInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertStudentPersonalInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertStudentPersonalInfoResponse,
    } as MsgInsertStudentPersonalInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertStudentPersonalInfoResponse {
    const message = {
      ...baseMsgInsertStudentPersonalInfoResponse,
    } as MsgInsertStudentPersonalInfoResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgInsertStudentPersonalInfoResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertStudentPersonalInfoResponse>
  ): MsgInsertStudentPersonalInfoResponse {
    const message = {
      ...baseMsgInsertStudentPersonalInfoResponse,
    } as MsgInsertStudentPersonalInfoResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgInsertStudentContactInfo: object = {
  creator: "",
  university: "",
  studentIndex: "",
  contactAddress: "",
  email: "",
  mobilePhone: "",
};

export const MsgInsertStudentContactInfo = {
  encode(
    message: MsgInsertStudentContactInfo,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    if (message.contactAddress !== "") {
      writer.uint32(34).string(message.contactAddress);
    }
    if (message.email !== "") {
      writer.uint32(42).string(message.email);
    }
    if (message.mobilePhone !== "") {
      writer.uint32(50).string(message.mobilePhone);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertStudentContactInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertStudentContactInfo,
    } as MsgInsertStudentContactInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        case 4:
          message.contactAddress = reader.string();
          break;
        case 5:
          message.email = reader.string();
          break;
        case 6:
          message.mobilePhone = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertStudentContactInfo {
    const message = {
      ...baseMsgInsertStudentContactInfo,
    } as MsgInsertStudentContactInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    if (object.contactAddress !== undefined && object.contactAddress !== null) {
      message.contactAddress = String(object.contactAddress);
    } else {
      message.contactAddress = "";
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = String(object.email);
    } else {
      message.email = "";
    }
    if (object.mobilePhone !== undefined && object.mobilePhone !== null) {
      message.mobilePhone = String(object.mobilePhone);
    } else {
      message.mobilePhone = "";
    }
    return message;
  },

  toJSON(message: MsgInsertStudentContactInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    message.contactAddress !== undefined &&
      (obj.contactAddress = message.contactAddress);
    message.email !== undefined && (obj.email = message.email);
    message.mobilePhone !== undefined &&
      (obj.mobilePhone = message.mobilePhone);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertStudentContactInfo>
  ): MsgInsertStudentContactInfo {
    const message = {
      ...baseMsgInsertStudentContactInfo,
    } as MsgInsertStudentContactInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    if (object.contactAddress !== undefined && object.contactAddress !== null) {
      message.contactAddress = object.contactAddress;
    } else {
      message.contactAddress = "";
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = object.email;
    } else {
      message.email = "";
    }
    if (object.mobilePhone !== undefined && object.mobilePhone !== null) {
      message.mobilePhone = object.mobilePhone;
    } else {
      message.mobilePhone = "";
    }
    return message;
  },
};

const baseMsgInsertStudentContactInfoResponse: object = { status: 0 };

export const MsgInsertStudentContactInfoResponse = {
  encode(
    message: MsgInsertStudentContactInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertStudentContactInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertStudentContactInfoResponse,
    } as MsgInsertStudentContactInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertStudentContactInfoResponse {
    const message = {
      ...baseMsgInsertStudentContactInfoResponse,
    } as MsgInsertStudentContactInfoResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgInsertStudentContactInfoResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertStudentContactInfoResponse>
  ): MsgInsertStudentContactInfoResponse {
    const message = {
      ...baseMsgInsertStudentContactInfoResponse,
    } as MsgInsertStudentContactInfoResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgInsertStudentResidenceInfo: object = {
  creator: "",
  university: "",
  studentIndex: "",
  country: "",
  province: "",
  town: "",
  postCode: "",
  address: "",
  houseNumber: "",
  homePhone: "",
};

export const MsgInsertStudentResidenceInfo = {
  encode(
    message: MsgInsertStudentResidenceInfo,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    if (message.country !== "") {
      writer.uint32(34).string(message.country);
    }
    if (message.province !== "") {
      writer.uint32(42).string(message.province);
    }
    if (message.town !== "") {
      writer.uint32(50).string(message.town);
    }
    if (message.postCode !== "") {
      writer.uint32(58).string(message.postCode);
    }
    if (message.address !== "") {
      writer.uint32(66).string(message.address);
    }
    if (message.houseNumber !== "") {
      writer.uint32(74).string(message.houseNumber);
    }
    if (message.homePhone !== "") {
      writer.uint32(82).string(message.homePhone);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertStudentResidenceInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertStudentResidenceInfo,
    } as MsgInsertStudentResidenceInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        case 4:
          message.country = reader.string();
          break;
        case 5:
          message.province = reader.string();
          break;
        case 6:
          message.town = reader.string();
          break;
        case 7:
          message.postCode = reader.string();
          break;
        case 8:
          message.address = reader.string();
          break;
        case 9:
          message.houseNumber = reader.string();
          break;
        case 10:
          message.homePhone = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertStudentResidenceInfo {
    const message = {
      ...baseMsgInsertStudentResidenceInfo,
    } as MsgInsertStudentResidenceInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = String(object.country);
    } else {
      message.country = "";
    }
    if (object.province !== undefined && object.province !== null) {
      message.province = String(object.province);
    } else {
      message.province = "";
    }
    if (object.town !== undefined && object.town !== null) {
      message.town = String(object.town);
    } else {
      message.town = "";
    }
    if (object.postCode !== undefined && object.postCode !== null) {
      message.postCode = String(object.postCode);
    } else {
      message.postCode = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.houseNumber !== undefined && object.houseNumber !== null) {
      message.houseNumber = String(object.houseNumber);
    } else {
      message.houseNumber = "";
    }
    if (object.homePhone !== undefined && object.homePhone !== null) {
      message.homePhone = String(object.homePhone);
    } else {
      message.homePhone = "";
    }
    return message;
  },

  toJSON(message: MsgInsertStudentResidenceInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    message.country !== undefined && (obj.country = message.country);
    message.province !== undefined && (obj.province = message.province);
    message.town !== undefined && (obj.town = message.town);
    message.postCode !== undefined && (obj.postCode = message.postCode);
    message.address !== undefined && (obj.address = message.address);
    message.houseNumber !== undefined &&
      (obj.houseNumber = message.houseNumber);
    message.homePhone !== undefined && (obj.homePhone = message.homePhone);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertStudentResidenceInfo>
  ): MsgInsertStudentResidenceInfo {
    const message = {
      ...baseMsgInsertStudentResidenceInfo,
    } as MsgInsertStudentResidenceInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = object.country;
    } else {
      message.country = "";
    }
    if (object.province !== undefined && object.province !== null) {
      message.province = object.province;
    } else {
      message.province = "";
    }
    if (object.town !== undefined && object.town !== null) {
      message.town = object.town;
    } else {
      message.town = "";
    }
    if (object.postCode !== undefined && object.postCode !== null) {
      message.postCode = object.postCode;
    } else {
      message.postCode = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.houseNumber !== undefined && object.houseNumber !== null) {
      message.houseNumber = object.houseNumber;
    } else {
      message.houseNumber = "";
    }
    if (object.homePhone !== undefined && object.homePhone !== null) {
      message.homePhone = object.homePhone;
    } else {
      message.homePhone = "";
    }
    return message;
  },
};

const baseMsgInsertStudentResidenceInfoResponse: object = { status: 0 };

export const MsgInsertStudentResidenceInfoResponse = {
  encode(
    message: MsgInsertStudentResidenceInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertStudentResidenceInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertStudentResidenceInfoResponse,
    } as MsgInsertStudentResidenceInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertStudentResidenceInfoResponse {
    const message = {
      ...baseMsgInsertStudentResidenceInfoResponse,
    } as MsgInsertStudentResidenceInfoResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgInsertStudentResidenceInfoResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertStudentResidenceInfoResponse>
  ): MsgInsertStudentResidenceInfoResponse {
    const message = {
      ...baseMsgInsertStudentResidenceInfoResponse,
    } as MsgInsertStudentResidenceInfoResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgInsertExamGrade: object = {
  creator: "",
  university: "",
  studentIndex: "",
  examName: "",
  grade: "",
};

export const MsgInsertExamGrade = {
  encode(
    message: MsgInsertExamGrade,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    if (message.examName !== "") {
      writer.uint32(34).string(message.examName);
    }
    if (message.grade !== "") {
      writer.uint32(42).string(message.grade);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInsertExamGrade {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInsertExamGrade } as MsgInsertExamGrade;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        case 4:
          message.examName = reader.string();
          break;
        case 5:
          message.grade = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertExamGrade {
    const message = { ...baseMsgInsertExamGrade } as MsgInsertExamGrade;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = String(object.examName);
    } else {
      message.examName = "";
    }
    if (object.grade !== undefined && object.grade !== null) {
      message.grade = String(object.grade);
    } else {
      message.grade = "";
    }
    return message;
  },

  toJSON(message: MsgInsertExamGrade): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    message.examName !== undefined && (obj.examName = message.examName);
    message.grade !== undefined && (obj.grade = message.grade);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInsertExamGrade>): MsgInsertExamGrade {
    const message = { ...baseMsgInsertExamGrade } as MsgInsertExamGrade;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = object.examName;
    } else {
      message.examName = "";
    }
    if (object.grade !== undefined && object.grade !== null) {
      message.grade = object.grade;
    } else {
      message.grade = "";
    }
    return message;
  },
};

const baseMsgInsertExamGradeResponse: object = { status: 0 };

export const MsgInsertExamGradeResponse = {
  encode(
    message: MsgInsertExamGradeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertExamGradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertExamGradeResponse,
    } as MsgInsertExamGradeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertExamGradeResponse {
    const message = {
      ...baseMsgInsertExamGradeResponse,
    } as MsgInsertExamGradeResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgInsertExamGradeResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertExamGradeResponse>
  ): MsgInsertExamGradeResponse {
    const message = {
      ...baseMsgInsertExamGradeResponse,
    } as MsgInsertExamGradeResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgPayTaxes: object = {
  creator: "",
  university: "",
  studentIndex: "",
};

export const MsgPayTaxes = {
  encode(message: MsgPayTaxes, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPayTaxes {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPayTaxes } as MsgPayTaxes;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPayTaxes {
    const message = { ...baseMsgPayTaxes } as MsgPayTaxes;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    return message;
  },

  toJSON(message: MsgPayTaxes): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPayTaxes>): MsgPayTaxes {
    const message = { ...baseMsgPayTaxes } as MsgPayTaxes;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    return message;
  },
};

const baseMsgPayTaxesResponse: object = { status: 0 };

export const MsgPayTaxesResponse = {
  encode(
    message: MsgPayTaxesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPayTaxesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPayTaxesResponse } as MsgPayTaxesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPayTaxesResponse {
    const message = { ...baseMsgPayTaxesResponse } as MsgPayTaxesResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgPayTaxesResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPayTaxesResponse>): MsgPayTaxesResponse {
    const message = { ...baseMsgPayTaxesResponse } as MsgPayTaxesResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgInsertErasmusRequest: object = {
  creator: "",
  university: "",
  studentIndex: "",
  durationInMonths: "",
  foreignUniversityName: "",
  erasmusType: "",
};

export const MsgInsertErasmusRequest = {
  encode(
    message: MsgInsertErasmusRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    if (message.durationInMonths !== "") {
      writer.uint32(34).string(message.durationInMonths);
    }
    if (message.foreignUniversityName !== "") {
      writer.uint32(42).string(message.foreignUniversityName);
    }
    if (message.erasmusType !== "") {
      writer.uint32(50).string(message.erasmusType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInsertErasmusRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertErasmusRequest,
    } as MsgInsertErasmusRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        case 4:
          message.durationInMonths = reader.string();
          break;
        case 5:
          message.foreignUniversityName = reader.string();
          break;
        case 6:
          message.erasmusType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertErasmusRequest {
    const message = {
      ...baseMsgInsertErasmusRequest,
    } as MsgInsertErasmusRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    if (
      object.durationInMonths !== undefined &&
      object.durationInMonths !== null
    ) {
      message.durationInMonths = String(object.durationInMonths);
    } else {
      message.durationInMonths = "";
    }
    if (
      object.foreignUniversityName !== undefined &&
      object.foreignUniversityName !== null
    ) {
      message.foreignUniversityName = String(object.foreignUniversityName);
    } else {
      message.foreignUniversityName = "";
    }
    if (object.erasmusType !== undefined && object.erasmusType !== null) {
      message.erasmusType = String(object.erasmusType);
    } else {
      message.erasmusType = "";
    }
    return message;
  },

  toJSON(message: MsgInsertErasmusRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    message.durationInMonths !== undefined &&
      (obj.durationInMonths = message.durationInMonths);
    message.foreignUniversityName !== undefined &&
      (obj.foreignUniversityName = message.foreignUniversityName);
    message.erasmusType !== undefined &&
      (obj.erasmusType = message.erasmusType);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertErasmusRequest>
  ): MsgInsertErasmusRequest {
    const message = {
      ...baseMsgInsertErasmusRequest,
    } as MsgInsertErasmusRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    if (
      object.durationInMonths !== undefined &&
      object.durationInMonths !== null
    ) {
      message.durationInMonths = object.durationInMonths;
    } else {
      message.durationInMonths = "";
    }
    if (
      object.foreignUniversityName !== undefined &&
      object.foreignUniversityName !== null
    ) {
      message.foreignUniversityName = object.foreignUniversityName;
    } else {
      message.foreignUniversityName = "";
    }
    if (object.erasmusType !== undefined && object.erasmusType !== null) {
      message.erasmusType = object.erasmusType;
    } else {
      message.erasmusType = "";
    }
    return message;
  },
};

const baseMsgInsertErasmusRequestResponse: object = { status: 0 };

export const MsgInsertErasmusRequestResponse = {
  encode(
    message: MsgInsertErasmusRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertErasmusRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertErasmusRequestResponse,
    } as MsgInsertErasmusRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertErasmusRequestResponse {
    const message = {
      ...baseMsgInsertErasmusRequestResponse,
    } as MsgInsertErasmusRequestResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgInsertErasmusRequestResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertErasmusRequestResponse>
  ): MsgInsertErasmusRequestResponse {
    const message = {
      ...baseMsgInsertErasmusRequestResponse,
    } as MsgInsertErasmusRequestResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgInsertErasmusExam: object = {
  creator: "",
  university: "",
  studentIndex: "",
  examName: "",
};

export const MsgInsertErasmusExam = {
  encode(
    message: MsgInsertErasmusExam,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    if (message.examName !== "") {
      writer.uint32(34).string(message.examName);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInsertErasmusExam {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInsertErasmusExam } as MsgInsertErasmusExam;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        case 4:
          message.examName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertErasmusExam {
    const message = { ...baseMsgInsertErasmusExam } as MsgInsertErasmusExam;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = String(object.examName);
    } else {
      message.examName = "";
    }
    return message;
  },

  toJSON(message: MsgInsertErasmusExam): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    message.examName !== undefined && (obj.examName = message.examName);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInsertErasmusExam>): MsgInsertErasmusExam {
    const message = { ...baseMsgInsertErasmusExam } as MsgInsertErasmusExam;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = object.examName;
    } else {
      message.examName = "";
    }
    return message;
  },
};

const baseMsgInsertErasmusExamResponse: object = { status: 0 };

export const MsgInsertErasmusExamResponse = {
  encode(
    message: MsgInsertErasmusExamResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgInsertErasmusExamResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInsertErasmusExamResponse,
    } as MsgInsertErasmusExamResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInsertErasmusExamResponse {
    const message = {
      ...baseMsgInsertErasmusExamResponse,
    } as MsgInsertErasmusExamResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgInsertErasmusExamResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInsertErasmusExamResponse>
  ): MsgInsertErasmusExamResponse {
    const message = {
      ...baseMsgInsertErasmusExamResponse,
    } as MsgInsertErasmusExamResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgStartErasmus: object = {
  creator: "",
  university: "",
  studentIndex: "",
};

export const MsgStartErasmus = {
  encode(message: MsgStartErasmus, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.university !== "") {
      writer.uint32(18).string(message.university);
    }
    if (message.studentIndex !== "") {
      writer.uint32(26).string(message.studentIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStartErasmus {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgStartErasmus } as MsgStartErasmus;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.university = reader.string();
          break;
        case 3:
          message.studentIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStartErasmus {
    const message = { ...baseMsgStartErasmus } as MsgStartErasmus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = String(object.university);
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = String(object.studentIndex);
    } else {
      message.studentIndex = "";
    }
    return message;
  },

  toJSON(message: MsgStartErasmus): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.university !== undefined && (obj.university = message.university);
    message.studentIndex !== undefined &&
      (obj.studentIndex = message.studentIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgStartErasmus>): MsgStartErasmus {
    const message = { ...baseMsgStartErasmus } as MsgStartErasmus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.university !== undefined && object.university !== null) {
      message.university = object.university;
    } else {
      message.university = "";
    }
    if (object.studentIndex !== undefined && object.studentIndex !== null) {
      message.studentIndex = object.studentIndex;
    } else {
      message.studentIndex = "";
    }
    return message;
  },
};

const baseMsgStartErasmusResponse: object = { status: 0 };

export const MsgStartErasmusResponse = {
  encode(
    message: MsgStartErasmusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStartErasmusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgStartErasmusResponse,
    } as MsgStartErasmusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStartErasmusResponse {
    const message = {
      ...baseMsgStartErasmusResponse,
    } as MsgStartErasmusResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgStartErasmusResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgStartErasmusResponse>
  ): MsgStartErasmusResponse {
    const message = {
      ...baseMsgStartErasmusResponse,
    } as MsgStartErasmusResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgSendErasmusStudent: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  index: "",
};

export const MsgSendErasmusStudent = {
  encode(
    message: MsgSendErasmusStudent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.index !== "") {
      writer.uint32(42).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendErasmusStudent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendErasmusStudent } as MsgSendErasmusStudent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendErasmusStudent {
    const message = { ...baseMsgSendErasmusStudent } as MsgSendErasmusStudent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgSendErasmusStudent): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendErasmusStudent>
  ): MsgSendErasmusStudent {
    const message = { ...baseMsgSendErasmusStudent } as MsgSendErasmusStudent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgSendErasmusStudentResponse: object = { status: 0 };

export const MsgSendErasmusStudentResponse = {
  encode(
    message: MsgSendErasmusStudentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendErasmusStudentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendErasmusStudentResponse {
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgSendErasmusStudentResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendErasmusStudentResponse>
  ): MsgSendErasmusStudentResponse {
    const message = {
      ...baseMsgSendErasmusStudentResponse,
    } as MsgSendErasmusStudentResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseMsgSendEndErasmusPeriodRequest: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  startingUniversityName: "",
  destinationUniversityName: "",
  index: "",
  foreignIndex: "",
};

export const MsgSendEndErasmusPeriodRequest = {
  encode(
    message: MsgSendEndErasmusPeriodRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.startingUniversityName !== "") {
      writer.uint32(42).string(message.startingUniversityName);
    }
    if (message.destinationUniversityName !== "") {
      writer.uint32(50).string(message.destinationUniversityName);
    }
    if (message.index !== "") {
      writer.uint32(58).string(message.index);
    }
    if (message.foreignIndex !== "") {
      writer.uint32(66).string(message.foreignIndex);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendEndErasmusPeriodRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendEndErasmusPeriodRequest,
    } as MsgSendEndErasmusPeriodRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.startingUniversityName = reader.string();
          break;
        case 6:
          message.destinationUniversityName = reader.string();
          break;
        case 7:
          message.index = reader.string();
          break;
        case 8:
          message.foreignIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendEndErasmusPeriodRequest {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequest,
    } as MsgSendEndErasmusPeriodRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
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

  toJSON(message: MsgSendEndErasmusPeriodRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
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
    object: DeepPartial<MsgSendEndErasmusPeriodRequest>
  ): MsgSendEndErasmusPeriodRequest {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequest,
    } as MsgSendEndErasmusPeriodRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
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

const baseMsgSendEndErasmusPeriodRequestResponse: object = { status: 0 };

export const MsgSendEndErasmusPeriodRequestResponse = {
  encode(
    message: MsgSendEndErasmusPeriodRequestResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendEndErasmusPeriodRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendEndErasmusPeriodRequestResponse,
    } as MsgSendEndErasmusPeriodRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendEndErasmusPeriodRequestResponse {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequestResponse,
    } as MsgSendEndErasmusPeriodRequestResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: MsgSendEndErasmusPeriodRequestResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendEndErasmusPeriodRequestResponse>
  ): MsgSendEndErasmusPeriodRequestResponse {
    const message = {
      ...baseMsgSendEndErasmusPeriodRequestResponse,
    } as MsgSendEndErasmusPeriodRequestResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ConfigureChain(
    request: MsgConfigureChain
  ): Promise<MsgConfigureChainResponse>;
  RegisterNewStudent(
    request: MsgRegisterNewStudent
  ): Promise<MsgRegisterNewStudentResponse>;
  InsertStudentPersonalInfo(
    request: MsgInsertStudentPersonalInfo
  ): Promise<MsgInsertStudentPersonalInfoResponse>;
  InsertStudentContactInfo(
    request: MsgInsertStudentContactInfo
  ): Promise<MsgInsertStudentContactInfoResponse>;
  InsertStudentResidenceInfo(
    request: MsgInsertStudentResidenceInfo
  ): Promise<MsgInsertStudentResidenceInfoResponse>;
  InsertExamGrade(
    request: MsgInsertExamGrade
  ): Promise<MsgInsertExamGradeResponse>;
  PayTaxes(request: MsgPayTaxes): Promise<MsgPayTaxesResponse>;
  InsertErasmusRequest(
    request: MsgInsertErasmusRequest
  ): Promise<MsgInsertErasmusRequestResponse>;
  InsertErasmusExam(
    request: MsgInsertErasmusExam
  ): Promise<MsgInsertErasmusExamResponse>;
  StartErasmus(request: MsgStartErasmus): Promise<MsgStartErasmusResponse>;
  SendErasmusStudent(
    request: MsgSendErasmusStudent
  ): Promise<MsgSendErasmusStudentResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendEndErasmusPeriodRequest(
    request: MsgSendEndErasmusPeriodRequest
  ): Promise<MsgSendEndErasmusPeriodRequestResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ConfigureChain(
    request: MsgConfigureChain
  ): Promise<MsgConfigureChainResponse> {
    const data = MsgConfigureChain.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "ConfigureChain",
      data
    );
    return promise.then((data) =>
      MsgConfigureChainResponse.decode(new Reader(data))
    );
  }

  RegisterNewStudent(
    request: MsgRegisterNewStudent
  ): Promise<MsgRegisterNewStudentResponse> {
    const data = MsgRegisterNewStudent.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "RegisterNewStudent",
      data
    );
    return promise.then((data) =>
      MsgRegisterNewStudentResponse.decode(new Reader(data))
    );
  }

  InsertStudentPersonalInfo(
    request: MsgInsertStudentPersonalInfo
  ): Promise<MsgInsertStudentPersonalInfoResponse> {
    const data = MsgInsertStudentPersonalInfo.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "InsertStudentPersonalInfo",
      data
    );
    return promise.then((data) =>
      MsgInsertStudentPersonalInfoResponse.decode(new Reader(data))
    );
  }

  InsertStudentContactInfo(
    request: MsgInsertStudentContactInfo
  ): Promise<MsgInsertStudentContactInfoResponse> {
    const data = MsgInsertStudentContactInfo.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "InsertStudentContactInfo",
      data
    );
    return promise.then((data) =>
      MsgInsertStudentContactInfoResponse.decode(new Reader(data))
    );
  }

  InsertStudentResidenceInfo(
    request: MsgInsertStudentResidenceInfo
  ): Promise<MsgInsertStudentResidenceInfoResponse> {
    const data = MsgInsertStudentResidenceInfo.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "InsertStudentResidenceInfo",
      data
    );
    return promise.then((data) =>
      MsgInsertStudentResidenceInfoResponse.decode(new Reader(data))
    );
  }

  InsertExamGrade(
    request: MsgInsertExamGrade
  ): Promise<MsgInsertExamGradeResponse> {
    const data = MsgInsertExamGrade.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "InsertExamGrade",
      data
    );
    return promise.then((data) =>
      MsgInsertExamGradeResponse.decode(new Reader(data))
    );
  }

  PayTaxes(request: MsgPayTaxes): Promise<MsgPayTaxesResponse> {
    const data = MsgPayTaxes.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "PayTaxes",
      data
    );
    return promise.then((data) => MsgPayTaxesResponse.decode(new Reader(data)));
  }

  InsertErasmusRequest(
    request: MsgInsertErasmusRequest
  ): Promise<MsgInsertErasmusRequestResponse> {
    const data = MsgInsertErasmusRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "InsertErasmusRequest",
      data
    );
    return promise.then((data) =>
      MsgInsertErasmusRequestResponse.decode(new Reader(data))
    );
  }

  InsertErasmusExam(
    request: MsgInsertErasmusExam
  ): Promise<MsgInsertErasmusExamResponse> {
    const data = MsgInsertErasmusExam.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "InsertErasmusExam",
      data
    );
    return promise.then((data) =>
      MsgInsertErasmusExamResponse.decode(new Reader(data))
    );
  }

  StartErasmus(request: MsgStartErasmus): Promise<MsgStartErasmusResponse> {
    const data = MsgStartErasmus.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "StartErasmus",
      data
    );
    return promise.then((data) =>
      MsgStartErasmusResponse.decode(new Reader(data))
    );
  }

  SendErasmusStudent(
    request: MsgSendErasmusStudent
  ): Promise<MsgSendErasmusStudentResponse> {
    const data = MsgSendErasmusStudent.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "SendErasmusStudent",
      data
    );
    return promise.then((data) =>
      MsgSendErasmusStudentResponse.decode(new Reader(data))
    );
  }

  SendEndErasmusPeriodRequest(
    request: MsgSendEndErasmusPeriodRequest
  ): Promise<MsgSendEndErasmusPeriodRequestResponse> {
    const data = MsgSendEndErasmusPeriodRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Msg",
      "SendEndErasmusPeriodRequest",
      data
    );
    return promise.then((data) =>
      MsgSendEndErasmusPeriodRequestResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
