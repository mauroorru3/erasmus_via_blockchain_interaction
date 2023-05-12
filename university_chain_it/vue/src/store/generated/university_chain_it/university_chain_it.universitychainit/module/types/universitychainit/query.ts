/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../universitychainit/params";
import { StudentInfo } from "../universitychainit/student_info";
import { TranscriptOfRecords } from "../universitychainit/transcript_of_records";
import { PersonalInfo } from "../universitychainit/personal_info";
import { ResidenceInfo } from "../universitychainit/residence_info";
import { ContactInfo } from "../universitychainit/contact_info";
import { TaxesInfo } from "../universitychainit/taxes_info";
import { ErasmusInfo } from "../universitychainit/erasmus_info";
import { ChainInfo } from "../universitychainit/chain_info";
import { ForeignUniversities } from "../universitychainit/foreign_universities";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { UniversityInfo } from "../universitychainit/university_info";
import { ProfessorsExams } from "../universitychainit/professors_exams";
import { StoredStudent } from "../universitychainit/stored_student";

export const protobufPackage = "university_chain_it.universitychainit";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetStudentInfoRequest {}

export interface QueryGetStudentInfoResponse {
  StudentInfo: StudentInfo | undefined;
}

export interface QueryGetTranscriptOfRecordsRequest {}

export interface QueryGetTranscriptOfRecordsResponse {
  TranscriptOfRecords: TranscriptOfRecords | undefined;
}

export interface QueryGetPersonalInfoRequest {}

export interface QueryGetPersonalInfoResponse {
  PersonalInfo: PersonalInfo | undefined;
}

export interface QueryGetResidenceInfoRequest {}

export interface QueryGetResidenceInfoResponse {
  ResidenceInfo: ResidenceInfo | undefined;
}

export interface QueryGetContactInfoRequest {}

export interface QueryGetContactInfoResponse {
  ContactInfo: ContactInfo | undefined;
}

export interface QueryGetTaxesInfoRequest {}

export interface QueryGetTaxesInfoResponse {
  TaxesInfo: TaxesInfo | undefined;
}

export interface QueryGetErasmusInfoRequest {}

export interface QueryGetErasmusInfoResponse {
  ErasmusInfo: ErasmusInfo | undefined;
}

export interface QueryGetChainInfoRequest {}

export interface QueryGetChainInfoResponse {
  ChainInfo: ChainInfo | undefined;
}

export interface QueryGetForeignUniversitiesRequest {
  universityName: string;
}

export interface QueryGetForeignUniversitiesResponse {
  foreignUniversities: ForeignUniversities | undefined;
}

export interface QueryAllForeignUniversitiesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllForeignUniversitiesResponse {
  foreignUniversities: ForeignUniversities[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUniversityInfoRequest {
  universityName: string;
}

export interface QueryGetUniversityInfoResponse {
  universityInfo: UniversityInfo | undefined;
}

export interface QueryAllUniversityInfoRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUniversityInfoResponse {
  universityInfo: UniversityInfo[];
  pagination: PageResponse | undefined;
}

export interface QueryGetProfessorsExamsRequest {
  examName: string;
}

export interface QueryGetProfessorsExamsResponse {
  professorsExams: ProfessorsExams | undefined;
}

export interface QueryAllProfessorsExamsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllProfessorsExamsResponse {
  professorsExams: ProfessorsExams[];
  pagination: PageResponse | undefined;
}

export interface QueryGetStoredStudentRequest {
  index: string;
}

export interface QueryGetStoredStudentResponse {
  storedStudent: StoredStudent | undefined;
}

export interface QueryAllStoredStudentRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredStudentResponse {
  storedStudent: StoredStudent[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetStudentInfoRequest: object = {};

export const QueryGetStudentInfoRequest = {
  encode(
    _: QueryGetStudentInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStudentInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStudentInfoRequest,
    } as QueryGetStudentInfoRequest;
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

  fromJSON(_: any): QueryGetStudentInfoRequest {
    const message = {
      ...baseQueryGetStudentInfoRequest,
    } as QueryGetStudentInfoRequest;
    return message;
  },

  toJSON(_: QueryGetStudentInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetStudentInfoRequest>
  ): QueryGetStudentInfoRequest {
    const message = {
      ...baseQueryGetStudentInfoRequest,
    } as QueryGetStudentInfoRequest;
    return message;
  },
};

const baseQueryGetStudentInfoResponse: object = {};

export const QueryGetStudentInfoResponse = {
  encode(
    message: QueryGetStudentInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.StudentInfo !== undefined) {
      StudentInfo.encode(
        message.StudentInfo,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStudentInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStudentInfoResponse,
    } as QueryGetStudentInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.StudentInfo = StudentInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStudentInfoResponse {
    const message = {
      ...baseQueryGetStudentInfoResponse,
    } as QueryGetStudentInfoResponse;
    if (object.StudentInfo !== undefined && object.StudentInfo !== null) {
      message.StudentInfo = StudentInfo.fromJSON(object.StudentInfo);
    } else {
      message.StudentInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStudentInfoResponse): unknown {
    const obj: any = {};
    message.StudentInfo !== undefined &&
      (obj.StudentInfo = message.StudentInfo
        ? StudentInfo.toJSON(message.StudentInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStudentInfoResponse>
  ): QueryGetStudentInfoResponse {
    const message = {
      ...baseQueryGetStudentInfoResponse,
    } as QueryGetStudentInfoResponse;
    if (object.StudentInfo !== undefined && object.StudentInfo !== null) {
      message.StudentInfo = StudentInfo.fromPartial(object.StudentInfo);
    } else {
      message.StudentInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetTranscriptOfRecordsRequest: object = {};

export const QueryGetTranscriptOfRecordsRequest = {
  encode(
    _: QueryGetTranscriptOfRecordsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTranscriptOfRecordsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTranscriptOfRecordsRequest,
    } as QueryGetTranscriptOfRecordsRequest;
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

  fromJSON(_: any): QueryGetTranscriptOfRecordsRequest {
    const message = {
      ...baseQueryGetTranscriptOfRecordsRequest,
    } as QueryGetTranscriptOfRecordsRequest;
    return message;
  },

  toJSON(_: QueryGetTranscriptOfRecordsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetTranscriptOfRecordsRequest>
  ): QueryGetTranscriptOfRecordsRequest {
    const message = {
      ...baseQueryGetTranscriptOfRecordsRequest,
    } as QueryGetTranscriptOfRecordsRequest;
    return message;
  },
};

const baseQueryGetTranscriptOfRecordsResponse: object = {};

export const QueryGetTranscriptOfRecordsResponse = {
  encode(
    message: QueryGetTranscriptOfRecordsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TranscriptOfRecords !== undefined) {
      TranscriptOfRecords.encode(
        message.TranscriptOfRecords,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTranscriptOfRecordsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTranscriptOfRecordsResponse,
    } as QueryGetTranscriptOfRecordsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TranscriptOfRecords = TranscriptOfRecords.decode(
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

  fromJSON(object: any): QueryGetTranscriptOfRecordsResponse {
    const message = {
      ...baseQueryGetTranscriptOfRecordsResponse,
    } as QueryGetTranscriptOfRecordsResponse;
    if (
      object.TranscriptOfRecords !== undefined &&
      object.TranscriptOfRecords !== null
    ) {
      message.TranscriptOfRecords = TranscriptOfRecords.fromJSON(
        object.TranscriptOfRecords
      );
    } else {
      message.TranscriptOfRecords = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTranscriptOfRecordsResponse): unknown {
    const obj: any = {};
    message.TranscriptOfRecords !== undefined &&
      (obj.TranscriptOfRecords = message.TranscriptOfRecords
        ? TranscriptOfRecords.toJSON(message.TranscriptOfRecords)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTranscriptOfRecordsResponse>
  ): QueryGetTranscriptOfRecordsResponse {
    const message = {
      ...baseQueryGetTranscriptOfRecordsResponse,
    } as QueryGetTranscriptOfRecordsResponse;
    if (
      object.TranscriptOfRecords !== undefined &&
      object.TranscriptOfRecords !== null
    ) {
      message.TranscriptOfRecords = TranscriptOfRecords.fromPartial(
        object.TranscriptOfRecords
      );
    } else {
      message.TranscriptOfRecords = undefined;
    }
    return message;
  },
};

const baseQueryGetPersonalInfoRequest: object = {};

export const QueryGetPersonalInfoRequest = {
  encode(
    _: QueryGetPersonalInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPersonalInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPersonalInfoRequest,
    } as QueryGetPersonalInfoRequest;
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

  fromJSON(_: any): QueryGetPersonalInfoRequest {
    const message = {
      ...baseQueryGetPersonalInfoRequest,
    } as QueryGetPersonalInfoRequest;
    return message;
  },

  toJSON(_: QueryGetPersonalInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetPersonalInfoRequest>
  ): QueryGetPersonalInfoRequest {
    const message = {
      ...baseQueryGetPersonalInfoRequest,
    } as QueryGetPersonalInfoRequest;
    return message;
  },
};

const baseQueryGetPersonalInfoResponse: object = {};

export const QueryGetPersonalInfoResponse = {
  encode(
    message: QueryGetPersonalInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.PersonalInfo !== undefined) {
      PersonalInfo.encode(
        message.PersonalInfo,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPersonalInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPersonalInfoResponse,
    } as QueryGetPersonalInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.PersonalInfo = PersonalInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPersonalInfoResponse {
    const message = {
      ...baseQueryGetPersonalInfoResponse,
    } as QueryGetPersonalInfoResponse;
    if (object.PersonalInfo !== undefined && object.PersonalInfo !== null) {
      message.PersonalInfo = PersonalInfo.fromJSON(object.PersonalInfo);
    } else {
      message.PersonalInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPersonalInfoResponse): unknown {
    const obj: any = {};
    message.PersonalInfo !== undefined &&
      (obj.PersonalInfo = message.PersonalInfo
        ? PersonalInfo.toJSON(message.PersonalInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPersonalInfoResponse>
  ): QueryGetPersonalInfoResponse {
    const message = {
      ...baseQueryGetPersonalInfoResponse,
    } as QueryGetPersonalInfoResponse;
    if (object.PersonalInfo !== undefined && object.PersonalInfo !== null) {
      message.PersonalInfo = PersonalInfo.fromPartial(object.PersonalInfo);
    } else {
      message.PersonalInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetResidenceInfoRequest: object = {};

export const QueryGetResidenceInfoRequest = {
  encode(
    _: QueryGetResidenceInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetResidenceInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetResidenceInfoRequest,
    } as QueryGetResidenceInfoRequest;
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

  fromJSON(_: any): QueryGetResidenceInfoRequest {
    const message = {
      ...baseQueryGetResidenceInfoRequest,
    } as QueryGetResidenceInfoRequest;
    return message;
  },

  toJSON(_: QueryGetResidenceInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetResidenceInfoRequest>
  ): QueryGetResidenceInfoRequest {
    const message = {
      ...baseQueryGetResidenceInfoRequest,
    } as QueryGetResidenceInfoRequest;
    return message;
  },
};

const baseQueryGetResidenceInfoResponse: object = {};

export const QueryGetResidenceInfoResponse = {
  encode(
    message: QueryGetResidenceInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ResidenceInfo !== undefined) {
      ResidenceInfo.encode(
        message.ResidenceInfo,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetResidenceInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetResidenceInfoResponse,
    } as QueryGetResidenceInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ResidenceInfo = ResidenceInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetResidenceInfoResponse {
    const message = {
      ...baseQueryGetResidenceInfoResponse,
    } as QueryGetResidenceInfoResponse;
    if (object.ResidenceInfo !== undefined && object.ResidenceInfo !== null) {
      message.ResidenceInfo = ResidenceInfo.fromJSON(object.ResidenceInfo);
    } else {
      message.ResidenceInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetResidenceInfoResponse): unknown {
    const obj: any = {};
    message.ResidenceInfo !== undefined &&
      (obj.ResidenceInfo = message.ResidenceInfo
        ? ResidenceInfo.toJSON(message.ResidenceInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetResidenceInfoResponse>
  ): QueryGetResidenceInfoResponse {
    const message = {
      ...baseQueryGetResidenceInfoResponse,
    } as QueryGetResidenceInfoResponse;
    if (object.ResidenceInfo !== undefined && object.ResidenceInfo !== null) {
      message.ResidenceInfo = ResidenceInfo.fromPartial(object.ResidenceInfo);
    } else {
      message.ResidenceInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetContactInfoRequest: object = {};

export const QueryGetContactInfoRequest = {
  encode(
    _: QueryGetContactInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetContactInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContactInfoRequest,
    } as QueryGetContactInfoRequest;
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

  fromJSON(_: any): QueryGetContactInfoRequest {
    const message = {
      ...baseQueryGetContactInfoRequest,
    } as QueryGetContactInfoRequest;
    return message;
  },

  toJSON(_: QueryGetContactInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetContactInfoRequest>
  ): QueryGetContactInfoRequest {
    const message = {
      ...baseQueryGetContactInfoRequest,
    } as QueryGetContactInfoRequest;
    return message;
  },
};

const baseQueryGetContactInfoResponse: object = {};

export const QueryGetContactInfoResponse = {
  encode(
    message: QueryGetContactInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ContactInfo !== undefined) {
      ContactInfo.encode(
        message.ContactInfo,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetContactInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContactInfoResponse,
    } as QueryGetContactInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ContactInfo = ContactInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContactInfoResponse {
    const message = {
      ...baseQueryGetContactInfoResponse,
    } as QueryGetContactInfoResponse;
    if (object.ContactInfo !== undefined && object.ContactInfo !== null) {
      message.ContactInfo = ContactInfo.fromJSON(object.ContactInfo);
    } else {
      message.ContactInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetContactInfoResponse): unknown {
    const obj: any = {};
    message.ContactInfo !== undefined &&
      (obj.ContactInfo = message.ContactInfo
        ? ContactInfo.toJSON(message.ContactInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetContactInfoResponse>
  ): QueryGetContactInfoResponse {
    const message = {
      ...baseQueryGetContactInfoResponse,
    } as QueryGetContactInfoResponse;
    if (object.ContactInfo !== undefined && object.ContactInfo !== null) {
      message.ContactInfo = ContactInfo.fromPartial(object.ContactInfo);
    } else {
      message.ContactInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetTaxesInfoRequest: object = {};

export const QueryGetTaxesInfoRequest = {
  encode(
    _: QueryGetTaxesInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTaxesInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTaxesInfoRequest,
    } as QueryGetTaxesInfoRequest;
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

  fromJSON(_: any): QueryGetTaxesInfoRequest {
    const message = {
      ...baseQueryGetTaxesInfoRequest,
    } as QueryGetTaxesInfoRequest;
    return message;
  },

  toJSON(_: QueryGetTaxesInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetTaxesInfoRequest>
  ): QueryGetTaxesInfoRequest {
    const message = {
      ...baseQueryGetTaxesInfoRequest,
    } as QueryGetTaxesInfoRequest;
    return message;
  },
};

const baseQueryGetTaxesInfoResponse: object = {};

export const QueryGetTaxesInfoResponse = {
  encode(
    message: QueryGetTaxesInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TaxesInfo !== undefined) {
      TaxesInfo.encode(message.TaxesInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTaxesInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTaxesInfoResponse,
    } as QueryGetTaxesInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TaxesInfo = TaxesInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTaxesInfoResponse {
    const message = {
      ...baseQueryGetTaxesInfoResponse,
    } as QueryGetTaxesInfoResponse;
    if (object.TaxesInfo !== undefined && object.TaxesInfo !== null) {
      message.TaxesInfo = TaxesInfo.fromJSON(object.TaxesInfo);
    } else {
      message.TaxesInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTaxesInfoResponse): unknown {
    const obj: any = {};
    message.TaxesInfo !== undefined &&
      (obj.TaxesInfo = message.TaxesInfo
        ? TaxesInfo.toJSON(message.TaxesInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTaxesInfoResponse>
  ): QueryGetTaxesInfoResponse {
    const message = {
      ...baseQueryGetTaxesInfoResponse,
    } as QueryGetTaxesInfoResponse;
    if (object.TaxesInfo !== undefined && object.TaxesInfo !== null) {
      message.TaxesInfo = TaxesInfo.fromPartial(object.TaxesInfo);
    } else {
      message.TaxesInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetErasmusInfoRequest: object = {};

export const QueryGetErasmusInfoRequest = {
  encode(
    _: QueryGetErasmusInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetErasmusInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetErasmusInfoRequest,
    } as QueryGetErasmusInfoRequest;
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

  fromJSON(_: any): QueryGetErasmusInfoRequest {
    const message = {
      ...baseQueryGetErasmusInfoRequest,
    } as QueryGetErasmusInfoRequest;
    return message;
  },

  toJSON(_: QueryGetErasmusInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetErasmusInfoRequest>
  ): QueryGetErasmusInfoRequest {
    const message = {
      ...baseQueryGetErasmusInfoRequest,
    } as QueryGetErasmusInfoRequest;
    return message;
  },
};

const baseQueryGetErasmusInfoResponse: object = {};

export const QueryGetErasmusInfoResponse = {
  encode(
    message: QueryGetErasmusInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ErasmusInfo !== undefined) {
      ErasmusInfo.encode(
        message.ErasmusInfo,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetErasmusInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetErasmusInfoResponse,
    } as QueryGetErasmusInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ErasmusInfo = ErasmusInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetErasmusInfoResponse {
    const message = {
      ...baseQueryGetErasmusInfoResponse,
    } as QueryGetErasmusInfoResponse;
    if (object.ErasmusInfo !== undefined && object.ErasmusInfo !== null) {
      message.ErasmusInfo = ErasmusInfo.fromJSON(object.ErasmusInfo);
    } else {
      message.ErasmusInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetErasmusInfoResponse): unknown {
    const obj: any = {};
    message.ErasmusInfo !== undefined &&
      (obj.ErasmusInfo = message.ErasmusInfo
        ? ErasmusInfo.toJSON(message.ErasmusInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetErasmusInfoResponse>
  ): QueryGetErasmusInfoResponse {
    const message = {
      ...baseQueryGetErasmusInfoResponse,
    } as QueryGetErasmusInfoResponse;
    if (object.ErasmusInfo !== undefined && object.ErasmusInfo !== null) {
      message.ErasmusInfo = ErasmusInfo.fromPartial(object.ErasmusInfo);
    } else {
      message.ErasmusInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetChainInfoRequest: object = {};

export const QueryGetChainInfoRequest = {
  encode(
    _: QueryGetChainInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetChainInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetChainInfoRequest,
    } as QueryGetChainInfoRequest;
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

  fromJSON(_: any): QueryGetChainInfoRequest {
    const message = {
      ...baseQueryGetChainInfoRequest,
    } as QueryGetChainInfoRequest;
    return message;
  },

  toJSON(_: QueryGetChainInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetChainInfoRequest>
  ): QueryGetChainInfoRequest {
    const message = {
      ...baseQueryGetChainInfoRequest,
    } as QueryGetChainInfoRequest;
    return message;
  },
};

const baseQueryGetChainInfoResponse: object = {};

export const QueryGetChainInfoResponse = {
  encode(
    message: QueryGetChainInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ChainInfo !== undefined) {
      ChainInfo.encode(message.ChainInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetChainInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetChainInfoResponse,
    } as QueryGetChainInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ChainInfo = ChainInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetChainInfoResponse {
    const message = {
      ...baseQueryGetChainInfoResponse,
    } as QueryGetChainInfoResponse;
    if (object.ChainInfo !== undefined && object.ChainInfo !== null) {
      message.ChainInfo = ChainInfo.fromJSON(object.ChainInfo);
    } else {
      message.ChainInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetChainInfoResponse): unknown {
    const obj: any = {};
    message.ChainInfo !== undefined &&
      (obj.ChainInfo = message.ChainInfo
        ? ChainInfo.toJSON(message.ChainInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetChainInfoResponse>
  ): QueryGetChainInfoResponse {
    const message = {
      ...baseQueryGetChainInfoResponse,
    } as QueryGetChainInfoResponse;
    if (object.ChainInfo !== undefined && object.ChainInfo !== null) {
      message.ChainInfo = ChainInfo.fromPartial(object.ChainInfo);
    } else {
      message.ChainInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetForeignUniversitiesRequest: object = { universityName: "" };

export const QueryGetForeignUniversitiesRequest = {
  encode(
    message: QueryGetForeignUniversitiesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.universityName !== "") {
      writer.uint32(10).string(message.universityName);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetForeignUniversitiesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetForeignUniversitiesRequest,
    } as QueryGetForeignUniversitiesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetForeignUniversitiesRequest {
    const message = {
      ...baseQueryGetForeignUniversitiesRequest,
    } as QueryGetForeignUniversitiesRequest;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = String(object.universityName);
    } else {
      message.universityName = "";
    }
    return message;
  },

  toJSON(message: QueryGetForeignUniversitiesRequest): unknown {
    const obj: any = {};
    message.universityName !== undefined &&
      (obj.universityName = message.universityName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetForeignUniversitiesRequest>
  ): QueryGetForeignUniversitiesRequest {
    const message = {
      ...baseQueryGetForeignUniversitiesRequest,
    } as QueryGetForeignUniversitiesRequest;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = object.universityName;
    } else {
      message.universityName = "";
    }
    return message;
  },
};

const baseQueryGetForeignUniversitiesResponse: object = {};

export const QueryGetForeignUniversitiesResponse = {
  encode(
    message: QueryGetForeignUniversitiesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.foreignUniversities !== undefined) {
      ForeignUniversities.encode(
        message.foreignUniversities,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetForeignUniversitiesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetForeignUniversitiesResponse,
    } as QueryGetForeignUniversitiesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.foreignUniversities = ForeignUniversities.decode(
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

  fromJSON(object: any): QueryGetForeignUniversitiesResponse {
    const message = {
      ...baseQueryGetForeignUniversitiesResponse,
    } as QueryGetForeignUniversitiesResponse;
    if (
      object.foreignUniversities !== undefined &&
      object.foreignUniversities !== null
    ) {
      message.foreignUniversities = ForeignUniversities.fromJSON(
        object.foreignUniversities
      );
    } else {
      message.foreignUniversities = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetForeignUniversitiesResponse): unknown {
    const obj: any = {};
    message.foreignUniversities !== undefined &&
      (obj.foreignUniversities = message.foreignUniversities
        ? ForeignUniversities.toJSON(message.foreignUniversities)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetForeignUniversitiesResponse>
  ): QueryGetForeignUniversitiesResponse {
    const message = {
      ...baseQueryGetForeignUniversitiesResponse,
    } as QueryGetForeignUniversitiesResponse;
    if (
      object.foreignUniversities !== undefined &&
      object.foreignUniversities !== null
    ) {
      message.foreignUniversities = ForeignUniversities.fromPartial(
        object.foreignUniversities
      );
    } else {
      message.foreignUniversities = undefined;
    }
    return message;
  },
};

const baseQueryAllForeignUniversitiesRequest: object = {};

export const QueryAllForeignUniversitiesRequest = {
  encode(
    message: QueryAllForeignUniversitiesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllForeignUniversitiesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllForeignUniversitiesRequest,
    } as QueryAllForeignUniversitiesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllForeignUniversitiesRequest {
    const message = {
      ...baseQueryAllForeignUniversitiesRequest,
    } as QueryAllForeignUniversitiesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllForeignUniversitiesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllForeignUniversitiesRequest>
  ): QueryAllForeignUniversitiesRequest {
    const message = {
      ...baseQueryAllForeignUniversitiesRequest,
    } as QueryAllForeignUniversitiesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllForeignUniversitiesResponse: object = {};

export const QueryAllForeignUniversitiesResponse = {
  encode(
    message: QueryAllForeignUniversitiesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.foreignUniversities) {
      ForeignUniversities.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllForeignUniversitiesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllForeignUniversitiesResponse,
    } as QueryAllForeignUniversitiesResponse;
    message.foreignUniversities = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.foreignUniversities.push(
            ForeignUniversities.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllForeignUniversitiesResponse {
    const message = {
      ...baseQueryAllForeignUniversitiesResponse,
    } as QueryAllForeignUniversitiesResponse;
    message.foreignUniversities = [];
    if (
      object.foreignUniversities !== undefined &&
      object.foreignUniversities !== null
    ) {
      for (const e of object.foreignUniversities) {
        message.foreignUniversities.push(ForeignUniversities.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllForeignUniversitiesResponse): unknown {
    const obj: any = {};
    if (message.foreignUniversities) {
      obj.foreignUniversities = message.foreignUniversities.map((e) =>
        e ? ForeignUniversities.toJSON(e) : undefined
      );
    } else {
      obj.foreignUniversities = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllForeignUniversitiesResponse>
  ): QueryAllForeignUniversitiesResponse {
    const message = {
      ...baseQueryAllForeignUniversitiesResponse,
    } as QueryAllForeignUniversitiesResponse;
    message.foreignUniversities = [];
    if (
      object.foreignUniversities !== undefined &&
      object.foreignUniversities !== null
    ) {
      for (const e of object.foreignUniversities) {
        message.foreignUniversities.push(ForeignUniversities.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetUniversityInfoRequest: object = { universityName: "" };

export const QueryGetUniversityInfoRequest = {
  encode(
    message: QueryGetUniversityInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.universityName !== "") {
      writer.uint32(10).string(message.universityName);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUniversityInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUniversityInfoRequest,
    } as QueryGetUniversityInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUniversityInfoRequest {
    const message = {
      ...baseQueryGetUniversityInfoRequest,
    } as QueryGetUniversityInfoRequest;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = String(object.universityName);
    } else {
      message.universityName = "";
    }
    return message;
  },

  toJSON(message: QueryGetUniversityInfoRequest): unknown {
    const obj: any = {};
    message.universityName !== undefined &&
      (obj.universityName = message.universityName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUniversityInfoRequest>
  ): QueryGetUniversityInfoRequest {
    const message = {
      ...baseQueryGetUniversityInfoRequest,
    } as QueryGetUniversityInfoRequest;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = object.universityName;
    } else {
      message.universityName = "";
    }
    return message;
  },
};

const baseQueryGetUniversityInfoResponse: object = {};

export const QueryGetUniversityInfoResponse = {
  encode(
    message: QueryGetUniversityInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.universityInfo !== undefined) {
      UniversityInfo.encode(
        message.universityInfo,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUniversityInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUniversityInfoResponse,
    } as QueryGetUniversityInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityInfo = UniversityInfo.decode(
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

  fromJSON(object: any): QueryGetUniversityInfoResponse {
    const message = {
      ...baseQueryGetUniversityInfoResponse,
    } as QueryGetUniversityInfoResponse;
    if (object.universityInfo !== undefined && object.universityInfo !== null) {
      message.universityInfo = UniversityInfo.fromJSON(object.universityInfo);
    } else {
      message.universityInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetUniversityInfoResponse): unknown {
    const obj: any = {};
    message.universityInfo !== undefined &&
      (obj.universityInfo = message.universityInfo
        ? UniversityInfo.toJSON(message.universityInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUniversityInfoResponse>
  ): QueryGetUniversityInfoResponse {
    const message = {
      ...baseQueryGetUniversityInfoResponse,
    } as QueryGetUniversityInfoResponse;
    if (object.universityInfo !== undefined && object.universityInfo !== null) {
      message.universityInfo = UniversityInfo.fromPartial(
        object.universityInfo
      );
    } else {
      message.universityInfo = undefined;
    }
    return message;
  },
};

const baseQueryAllUniversityInfoRequest: object = {};

export const QueryAllUniversityInfoRequest = {
  encode(
    message: QueryAllUniversityInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllUniversityInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUniversityInfoRequest,
    } as QueryAllUniversityInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUniversityInfoRequest {
    const message = {
      ...baseQueryAllUniversityInfoRequest,
    } as QueryAllUniversityInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUniversityInfoRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUniversityInfoRequest>
  ): QueryAllUniversityInfoRequest {
    const message = {
      ...baseQueryAllUniversityInfoRequest,
    } as QueryAllUniversityInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllUniversityInfoResponse: object = {};

export const QueryAllUniversityInfoResponse = {
  encode(
    message: QueryAllUniversityInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.universityInfo) {
      UniversityInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllUniversityInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUniversityInfoResponse,
    } as QueryAllUniversityInfoResponse;
    message.universityInfo = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityInfo.push(
            UniversityInfo.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUniversityInfoResponse {
    const message = {
      ...baseQueryAllUniversityInfoResponse,
    } as QueryAllUniversityInfoResponse;
    message.universityInfo = [];
    if (object.universityInfo !== undefined && object.universityInfo !== null) {
      for (const e of object.universityInfo) {
        message.universityInfo.push(UniversityInfo.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUniversityInfoResponse): unknown {
    const obj: any = {};
    if (message.universityInfo) {
      obj.universityInfo = message.universityInfo.map((e) =>
        e ? UniversityInfo.toJSON(e) : undefined
      );
    } else {
      obj.universityInfo = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUniversityInfoResponse>
  ): QueryAllUniversityInfoResponse {
    const message = {
      ...baseQueryAllUniversityInfoResponse,
    } as QueryAllUniversityInfoResponse;
    message.universityInfo = [];
    if (object.universityInfo !== undefined && object.universityInfo !== null) {
      for (const e of object.universityInfo) {
        message.universityInfo.push(UniversityInfo.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetProfessorsExamsRequest: object = { examName: "" };

export const QueryGetProfessorsExamsRequest = {
  encode(
    message: QueryGetProfessorsExamsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.examName !== "") {
      writer.uint32(10).string(message.examName);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetProfessorsExamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetProfessorsExamsRequest,
    } as QueryGetProfessorsExamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.examName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProfessorsExamsRequest {
    const message = {
      ...baseQueryGetProfessorsExamsRequest,
    } as QueryGetProfessorsExamsRequest;
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = String(object.examName);
    } else {
      message.examName = "";
    }
    return message;
  },

  toJSON(message: QueryGetProfessorsExamsRequest): unknown {
    const obj: any = {};
    message.examName !== undefined && (obj.examName = message.examName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfessorsExamsRequest>
  ): QueryGetProfessorsExamsRequest {
    const message = {
      ...baseQueryGetProfessorsExamsRequest,
    } as QueryGetProfessorsExamsRequest;
    if (object.examName !== undefined && object.examName !== null) {
      message.examName = object.examName;
    } else {
      message.examName = "";
    }
    return message;
  },
};

const baseQueryGetProfessorsExamsResponse: object = {};

export const QueryGetProfessorsExamsResponse = {
  encode(
    message: QueryGetProfessorsExamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.professorsExams !== undefined) {
      ProfessorsExams.encode(
        message.professorsExams,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetProfessorsExamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetProfessorsExamsResponse,
    } as QueryGetProfessorsExamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.professorsExams = ProfessorsExams.decode(
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

  fromJSON(object: any): QueryGetProfessorsExamsResponse {
    const message = {
      ...baseQueryGetProfessorsExamsResponse,
    } as QueryGetProfessorsExamsResponse;
    if (
      object.professorsExams !== undefined &&
      object.professorsExams !== null
    ) {
      message.professorsExams = ProfessorsExams.fromJSON(
        object.professorsExams
      );
    } else {
      message.professorsExams = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetProfessorsExamsResponse): unknown {
    const obj: any = {};
    message.professorsExams !== undefined &&
      (obj.professorsExams = message.professorsExams
        ? ProfessorsExams.toJSON(message.professorsExams)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfessorsExamsResponse>
  ): QueryGetProfessorsExamsResponse {
    const message = {
      ...baseQueryGetProfessorsExamsResponse,
    } as QueryGetProfessorsExamsResponse;
    if (
      object.professorsExams !== undefined &&
      object.professorsExams !== null
    ) {
      message.professorsExams = ProfessorsExams.fromPartial(
        object.professorsExams
      );
    } else {
      message.professorsExams = undefined;
    }
    return message;
  },
};

const baseQueryAllProfessorsExamsRequest: object = {};

export const QueryAllProfessorsExamsRequest = {
  encode(
    message: QueryAllProfessorsExamsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllProfessorsExamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllProfessorsExamsRequest,
    } as QueryAllProfessorsExamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProfessorsExamsRequest {
    const message = {
      ...baseQueryAllProfessorsExamsRequest,
    } as QueryAllProfessorsExamsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllProfessorsExamsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllProfessorsExamsRequest>
  ): QueryAllProfessorsExamsRequest {
    const message = {
      ...baseQueryAllProfessorsExamsRequest,
    } as QueryAllProfessorsExamsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllProfessorsExamsResponse: object = {};

export const QueryAllProfessorsExamsResponse = {
  encode(
    message: QueryAllProfessorsExamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.professorsExams) {
      ProfessorsExams.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllProfessorsExamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllProfessorsExamsResponse,
    } as QueryAllProfessorsExamsResponse;
    message.professorsExams = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.professorsExams.push(
            ProfessorsExams.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProfessorsExamsResponse {
    const message = {
      ...baseQueryAllProfessorsExamsResponse,
    } as QueryAllProfessorsExamsResponse;
    message.professorsExams = [];
    if (
      object.professorsExams !== undefined &&
      object.professorsExams !== null
    ) {
      for (const e of object.professorsExams) {
        message.professorsExams.push(ProfessorsExams.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllProfessorsExamsResponse): unknown {
    const obj: any = {};
    if (message.professorsExams) {
      obj.professorsExams = message.professorsExams.map((e) =>
        e ? ProfessorsExams.toJSON(e) : undefined
      );
    } else {
      obj.professorsExams = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllProfessorsExamsResponse>
  ): QueryAllProfessorsExamsResponse {
    const message = {
      ...baseQueryAllProfessorsExamsResponse,
    } as QueryAllProfessorsExamsResponse;
    message.professorsExams = [];
    if (
      object.professorsExams !== undefined &&
      object.professorsExams !== null
    ) {
      for (const e of object.professorsExams) {
        message.professorsExams.push(ProfessorsExams.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetStoredStudentRequest: object = { index: "" };

export const QueryGetStoredStudentRequest = {
  encode(
    message: QueryGetStoredStudentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredStudentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredStudentRequest,
    } as QueryGetStoredStudentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredStudentRequest {
    const message = {
      ...baseQueryGetStoredStudentRequest,
    } as QueryGetStoredStudentRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetStoredStudentRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredStudentRequest>
  ): QueryGetStoredStudentRequest {
    const message = {
      ...baseQueryGetStoredStudentRequest,
    } as QueryGetStoredStudentRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetStoredStudentResponse: object = {};

export const QueryGetStoredStudentResponse = {
  encode(
    message: QueryGetStoredStudentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storedStudent !== undefined) {
      StoredStudent.encode(
        message.storedStudent,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredStudentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredStudentResponse,
    } as QueryGetStoredStudentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedStudent = StoredStudent.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredStudentResponse {
    const message = {
      ...baseQueryGetStoredStudentResponse,
    } as QueryGetStoredStudentResponse;
    if (object.storedStudent !== undefined && object.storedStudent !== null) {
      message.storedStudent = StoredStudent.fromJSON(object.storedStudent);
    } else {
      message.storedStudent = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStoredStudentResponse): unknown {
    const obj: any = {};
    message.storedStudent !== undefined &&
      (obj.storedStudent = message.storedStudent
        ? StoredStudent.toJSON(message.storedStudent)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredStudentResponse>
  ): QueryGetStoredStudentResponse {
    const message = {
      ...baseQueryGetStoredStudentResponse,
    } as QueryGetStoredStudentResponse;
    if (object.storedStudent !== undefined && object.storedStudent !== null) {
      message.storedStudent = StoredStudent.fromPartial(object.storedStudent);
    } else {
      message.storedStudent = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredStudentRequest: object = {};

export const QueryAllStoredStudentRequest = {
  encode(
    message: QueryAllStoredStudentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredStudentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredStudentRequest,
    } as QueryAllStoredStudentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredStudentRequest {
    const message = {
      ...baseQueryAllStoredStudentRequest,
    } as QueryAllStoredStudentRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredStudentRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredStudentRequest>
  ): QueryAllStoredStudentRequest {
    const message = {
      ...baseQueryAllStoredStudentRequest,
    } as QueryAllStoredStudentRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredStudentResponse: object = {};

export const QueryAllStoredStudentResponse = {
  encode(
    message: QueryAllStoredStudentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.storedStudent) {
      StoredStudent.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredStudentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredStudentResponse,
    } as QueryAllStoredStudentResponse;
    message.storedStudent = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedStudent.push(
            StoredStudent.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredStudentResponse {
    const message = {
      ...baseQueryAllStoredStudentResponse,
    } as QueryAllStoredStudentResponse;
    message.storedStudent = [];
    if (object.storedStudent !== undefined && object.storedStudent !== null) {
      for (const e of object.storedStudent) {
        message.storedStudent.push(StoredStudent.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredStudentResponse): unknown {
    const obj: any = {};
    if (message.storedStudent) {
      obj.storedStudent = message.storedStudent.map((e) =>
        e ? StoredStudent.toJSON(e) : undefined
      );
    } else {
      obj.storedStudent = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredStudentResponse>
  ): QueryAllStoredStudentResponse {
    const message = {
      ...baseQueryAllStoredStudentResponse,
    } as QueryAllStoredStudentResponse;
    message.storedStudent = [];
    if (object.storedStudent !== undefined && object.storedStudent !== null) {
      for (const e of object.storedStudent) {
        message.storedStudent.push(StoredStudent.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a StudentInfo by index. */
  StudentInfo(
    request: QueryGetStudentInfoRequest
  ): Promise<QueryGetStudentInfoResponse>;
  /** Queries a TranscriptOfRecords by index. */
  TranscriptOfRecords(
    request: QueryGetTranscriptOfRecordsRequest
  ): Promise<QueryGetTranscriptOfRecordsResponse>;
  /** Queries a PersonalInfo by index. */
  PersonalInfo(
    request: QueryGetPersonalInfoRequest
  ): Promise<QueryGetPersonalInfoResponse>;
  /** Queries a ResidenceInfo by index. */
  ResidenceInfo(
    request: QueryGetResidenceInfoRequest
  ): Promise<QueryGetResidenceInfoResponse>;
  /** Queries a ContactInfo by index. */
  ContactInfo(
    request: QueryGetContactInfoRequest
  ): Promise<QueryGetContactInfoResponse>;
  /** Queries a TaxesInfo by index. */
  TaxesInfo(
    request: QueryGetTaxesInfoRequest
  ): Promise<QueryGetTaxesInfoResponse>;
  /** Queries a ErasmusInfo by index. */
  ErasmusInfo(
    request: QueryGetErasmusInfoRequest
  ): Promise<QueryGetErasmusInfoResponse>;
  /** Queries a ChainInfo by index. */
  ChainInfo(
    request: QueryGetChainInfoRequest
  ): Promise<QueryGetChainInfoResponse>;
  /** Queries a ForeignUniversities by index. */
  ForeignUniversities(
    request: QueryGetForeignUniversitiesRequest
  ): Promise<QueryGetForeignUniversitiesResponse>;
  /** Queries a list of ForeignUniversities items. */
  ForeignUniversitiesAll(
    request: QueryAllForeignUniversitiesRequest
  ): Promise<QueryAllForeignUniversitiesResponse>;
  /** Queries a UniversityInfo by index. */
  UniversityInfo(
    request: QueryGetUniversityInfoRequest
  ): Promise<QueryGetUniversityInfoResponse>;
  /** Queries a list of UniversityInfo items. */
  UniversityInfoAll(
    request: QueryAllUniversityInfoRequest
  ): Promise<QueryAllUniversityInfoResponse>;
  /** Queries a ProfessorsExams by index. */
  ProfessorsExams(
    request: QueryGetProfessorsExamsRequest
  ): Promise<QueryGetProfessorsExamsResponse>;
  /** Queries a list of ProfessorsExams items. */
  ProfessorsExamsAll(
    request: QueryAllProfessorsExamsRequest
  ): Promise<QueryAllProfessorsExamsResponse>;
  /** Queries a StoredStudent by index. */
  StoredStudent(
    request: QueryGetStoredStudentRequest
  ): Promise<QueryGetStoredStudentResponse>;
  /** Queries a list of StoredStudent items. */
  StoredStudentAll(
    request: QueryAllStoredStudentRequest
  ): Promise<QueryAllStoredStudentResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  StudentInfo(
    request: QueryGetStudentInfoRequest
  ): Promise<QueryGetStudentInfoResponse> {
    const data = QueryGetStudentInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "StudentInfo",
      data
    );
    return promise.then((data) =>
      QueryGetStudentInfoResponse.decode(new Reader(data))
    );
  }

  TranscriptOfRecords(
    request: QueryGetTranscriptOfRecordsRequest
  ): Promise<QueryGetTranscriptOfRecordsResponse> {
    const data = QueryGetTranscriptOfRecordsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "TranscriptOfRecords",
      data
    );
    return promise.then((data) =>
      QueryGetTranscriptOfRecordsResponse.decode(new Reader(data))
    );
  }

  PersonalInfo(
    request: QueryGetPersonalInfoRequest
  ): Promise<QueryGetPersonalInfoResponse> {
    const data = QueryGetPersonalInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "PersonalInfo",
      data
    );
    return promise.then((data) =>
      QueryGetPersonalInfoResponse.decode(new Reader(data))
    );
  }

  ResidenceInfo(
    request: QueryGetResidenceInfoRequest
  ): Promise<QueryGetResidenceInfoResponse> {
    const data = QueryGetResidenceInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ResidenceInfo",
      data
    );
    return promise.then((data) =>
      QueryGetResidenceInfoResponse.decode(new Reader(data))
    );
  }

  ContactInfo(
    request: QueryGetContactInfoRequest
  ): Promise<QueryGetContactInfoResponse> {
    const data = QueryGetContactInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ContactInfo",
      data
    );
    return promise.then((data) =>
      QueryGetContactInfoResponse.decode(new Reader(data))
    );
  }

  TaxesInfo(
    request: QueryGetTaxesInfoRequest
  ): Promise<QueryGetTaxesInfoResponse> {
    const data = QueryGetTaxesInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "TaxesInfo",
      data
    );
    return promise.then((data) =>
      QueryGetTaxesInfoResponse.decode(new Reader(data))
    );
  }

  ErasmusInfo(
    request: QueryGetErasmusInfoRequest
  ): Promise<QueryGetErasmusInfoResponse> {
    const data = QueryGetErasmusInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ErasmusInfo",
      data
    );
    return promise.then((data) =>
      QueryGetErasmusInfoResponse.decode(new Reader(data))
    );
  }

  ChainInfo(
    request: QueryGetChainInfoRequest
  ): Promise<QueryGetChainInfoResponse> {
    const data = QueryGetChainInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ChainInfo",
      data
    );
    return promise.then((data) =>
      QueryGetChainInfoResponse.decode(new Reader(data))
    );
  }

  ForeignUniversities(
    request: QueryGetForeignUniversitiesRequest
  ): Promise<QueryGetForeignUniversitiesResponse> {
    const data = QueryGetForeignUniversitiesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ForeignUniversities",
      data
    );
    return promise.then((data) =>
      QueryGetForeignUniversitiesResponse.decode(new Reader(data))
    );
  }

  ForeignUniversitiesAll(
    request: QueryAllForeignUniversitiesRequest
  ): Promise<QueryAllForeignUniversitiesResponse> {
    const data = QueryAllForeignUniversitiesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ForeignUniversitiesAll",
      data
    );
    return promise.then((data) =>
      QueryAllForeignUniversitiesResponse.decode(new Reader(data))
    );
  }

  UniversityInfo(
    request: QueryGetUniversityInfoRequest
  ): Promise<QueryGetUniversityInfoResponse> {
    const data = QueryGetUniversityInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "UniversityInfo",
      data
    );
    return promise.then((data) =>
      QueryGetUniversityInfoResponse.decode(new Reader(data))
    );
  }

  UniversityInfoAll(
    request: QueryAllUniversityInfoRequest
  ): Promise<QueryAllUniversityInfoResponse> {
    const data = QueryAllUniversityInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "UniversityInfoAll",
      data
    );
    return promise.then((data) =>
      QueryAllUniversityInfoResponse.decode(new Reader(data))
    );
  }

  ProfessorsExams(
    request: QueryGetProfessorsExamsRequest
  ): Promise<QueryGetProfessorsExamsResponse> {
    const data = QueryGetProfessorsExamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ProfessorsExams",
      data
    );
    return promise.then((data) =>
      QueryGetProfessorsExamsResponse.decode(new Reader(data))
    );
  }

  ProfessorsExamsAll(
    request: QueryAllProfessorsExamsRequest
  ): Promise<QueryAllProfessorsExamsResponse> {
    const data = QueryAllProfessorsExamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "ProfessorsExamsAll",
      data
    );
    return promise.then((data) =>
      QueryAllProfessorsExamsResponse.decode(new Reader(data))
    );
  }

  StoredStudent(
    request: QueryGetStoredStudentRequest
  ): Promise<QueryGetStoredStudentResponse> {
    const data = QueryGetStoredStudentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "StoredStudent",
      data
    );
    return promise.then((data) =>
      QueryGetStoredStudentResponse.decode(new Reader(data))
    );
  }

  StoredStudentAll(
    request: QueryAllStoredStudentRequest
  ): Promise<QueryAllStoredStudentResponse> {
    const data = QueryAllStoredStudentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "university_chain_it.universitychainit.Query",
      "StoredStudentAll",
      data
    );
    return promise.then((data) =>
      QueryAllStoredStudentResponse.decode(new Reader(data))
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
