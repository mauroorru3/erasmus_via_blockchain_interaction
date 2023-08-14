/* eslint-disable */
import { ProfessorsExams } from "../universitychainit/professors_exams";
import { UniversityInfo } from "../universitychainit/university_info";
import { StoredStudent } from "../universitychainit/stored_student";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "university_chain_it.universitychainit";

export interface UniversityDetails {
  universityName: string;
  professorsInfo: ProfessorsExams | undefined;
  universityData: UniversityInfo | undefined;
  studentDetails: StoredStudent | undefined;
}

const baseUniversityDetails: object = { universityName: "" };

export const UniversityDetails = {
  encode(message: UniversityDetails, writer: Writer = Writer.create()): Writer {
    if (message.universityName !== "") {
      writer.uint32(10).string(message.universityName);
    }
    if (message.professorsInfo !== undefined) {
      ProfessorsExams.encode(
        message.professorsInfo,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.universityData !== undefined) {
      UniversityInfo.encode(
        message.universityData,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.studentDetails !== undefined) {
      StoredStudent.encode(
        message.studentDetails,
        writer.uint32(34).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UniversityDetails {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUniversityDetails } as UniversityDetails;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityName = reader.string();
          break;
        case 2:
          message.professorsInfo = ProfessorsExams.decode(
            reader,
            reader.uint32()
          );
          break;
        case 3:
          message.universityData = UniversityInfo.decode(
            reader,
            reader.uint32()
          );
          break;
        case 4:
          message.studentDetails = StoredStudent.decode(
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

  fromJSON(object: any): UniversityDetails {
    const message = { ...baseUniversityDetails } as UniversityDetails;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = String(object.universityName);
    } else {
      message.universityName = "";
    }
    if (object.professorsInfo !== undefined && object.professorsInfo !== null) {
      message.professorsInfo = ProfessorsExams.fromJSON(object.professorsInfo);
    } else {
      message.professorsInfo = undefined;
    }
    if (object.universityData !== undefined && object.universityData !== null) {
      message.universityData = UniversityInfo.fromJSON(object.universityData);
    } else {
      message.universityData = undefined;
    }
    if (object.studentDetails !== undefined && object.studentDetails !== null) {
      message.studentDetails = StoredStudent.fromJSON(object.studentDetails);
    } else {
      message.studentDetails = undefined;
    }
    return message;
  },

  toJSON(message: UniversityDetails): unknown {
    const obj: any = {};
    message.universityName !== undefined &&
      (obj.universityName = message.universityName);
    message.professorsInfo !== undefined &&
      (obj.professorsInfo = message.professorsInfo
        ? ProfessorsExams.toJSON(message.professorsInfo)
        : undefined);
    message.universityData !== undefined &&
      (obj.universityData = message.universityData
        ? UniversityInfo.toJSON(message.universityData)
        : undefined);
    message.studentDetails !== undefined &&
      (obj.studentDetails = message.studentDetails
        ? StoredStudent.toJSON(message.studentDetails)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<UniversityDetails>): UniversityDetails {
    const message = { ...baseUniversityDetails } as UniversityDetails;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = object.universityName;
    } else {
      message.universityName = "";
    }
    if (object.professorsInfo !== undefined && object.professorsInfo !== null) {
      message.professorsInfo = ProfessorsExams.fromPartial(
        object.professorsInfo
      );
    } else {
      message.professorsInfo = undefined;
    }
    if (object.universityData !== undefined && object.universityData !== null) {
      message.universityData = UniversityInfo.fromPartial(
        object.universityData
      );
    } else {
      message.universityData = undefined;
    }
    if (object.studentDetails !== undefined && object.studentDetails !== null) {
      message.studentDetails = StoredStudent.fromPartial(object.studentDetails);
    } else {
      message.studentDetails = undefined;
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
