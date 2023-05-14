/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface HubChainInfo {
  chainKey?: string;
  chainAdministratorKey?: string;
  initStatus?: boolean;
}

export interface HubContactInfo {
  contactAddress?: string;
  email?: string;
  mobilePhone?: string;
}

export interface HubErasmusInfo {
  erasmusStudent?: string;

  /** @format int64 */
  numberTimes?: number;

  /** @format int64 */
  numberMonths?: number;

  /** @format int64 */
  totalExams?: number;

  /** @format int64 */
  examsPassed?: number;

  /** @format int64 */
  totalCredits?: number;

  /** @format int64 */
  achievedCredits?: number;
  career?: string;
  previousStudentFifo?: string;
  nextStudentFifo?: string;
}

export interface HubMsgConfigureChainResponse {
  /** @format int32 */
  status?: number;
}

export interface HubMsgSendErasmusStudentResponse {
  /** @format int32 */
  status?: number;
}

/**
 * Params defines the parameters for the module.
 */
export type HubParams = object;

export interface HubPersonalInfo {
  gender?: string;
  dateOfBirth?: string;
  primaryNationality?: string;
  countryOfBirth?: string;
  provinceOfBirth?: string;
  townOfBirth?: string;
  taxCode?: string;
}

export interface HubQueryAllStoredStudentResponse {
  storedStudent?: HubStoredStudent[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface HubQueryAllUniversitiesResponse {
  universities?: HubUniversities[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface HubQueryGetChainInfoResponse {
  ChainInfo?: HubChainInfo;
}

export interface HubQueryGetContactInfoResponse {
  ContactInfo?: HubContactInfo;
}

export interface HubQueryGetErasmusInfoResponse {
  ErasmusInfo?: HubErasmusInfo;
}

export interface HubQueryGetPersonalInfoResponse {
  PersonalInfo?: HubPersonalInfo;
}

export interface HubQueryGetResidenceInfoResponse {
  ResidenceInfo?: HubResidenceInfo;
}

export interface HubQueryGetStoredStudentResponse {
  storedStudent?: HubStoredStudent;
}

export interface HubQueryGetStudentInfoResponse {
  StudentInfo?: HubStudentInfo;
}

export interface HubQueryGetTaxesInfoResponse {
  TaxesInfo?: HubTaxesInfo;
}

export interface HubQueryGetTranscriptOfRecordsResponse {
  TranscriptOfRecords?: HubTranscriptOfRecords;
}

export interface HubQueryGetUniversitiesResponse {
  universities?: HubUniversities;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface HubQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: HubParams;
}

export interface HubResidenceInfo {
  country?: string;
  province?: string;
  town?: string;
  postCode?: string;
  address?: string;
  houseNumber?: string;
  homePhone?: string;
}

export interface HubStoredStudent {
  index?: string;
  studentData?: HubStudentInfo;
  transcriptData?: HubTranscriptOfRecords;
  personalData?: HubPersonalInfo;
  residenceData?: HubResidenceInfo;
  contactData?: HubContactInfo;
  taxesData?: HubTaxesInfo;
  erasmusData?: HubErasmusInfo;
}

export interface HubStudentInfo {
  name?: string;
  surname?: string;
  courseType?: string;
  courseOfStudy?: string;
  status?: string;

  /** @format int64 */
  currentYearOfStudy?: number;
  outOfCourse?: boolean;

  /** @format int64 */
  numberOfYearsOutOfCourse?: number;
  studentKey?: string;
  completeInformation?: number[];
}

export interface HubTaxesInfo {
  status?: boolean;

  /** @format int64 */
  totalAmount?: number;
  taxesHistory?: string;
}

export interface HubTranscriptOfRecords {
  examsData?: string;

  /** @format int64 */
  totalExams?: number;

  /** @format int64 */
  examsPassed?: number;

  /** @format int64 */
  totalCredits?: number;

  /** @format int64 */
  achievedCredits?: number;
}

export interface HubUniversities {
  universityName?: string;
  universitiesKey?: string;
  universitiesCountry?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title hub/chain_info.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryChainInfo
   * @summary Queries a ChainInfo by index.
   * @request GET:/hub/hub/chain_info
   */
  queryChainInfo = (params: RequestParams = {}) =>
    this.request<HubQueryGetChainInfoResponse, RpcStatus>({
      path: `/hub/hub/chain_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryContactInfo
   * @summary Queries a ContactInfo by index.
   * @request GET:/hub/hub/contact_info
   */
  queryContactInfo = (params: RequestParams = {}) =>
    this.request<HubQueryGetContactInfoResponse, RpcStatus>({
      path: `/hub/hub/contact_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryErasmusInfo
   * @summary Queries a ErasmusInfo by index.
   * @request GET:/hub/hub/erasmus_info
   */
  queryErasmusInfo = (params: RequestParams = {}) =>
    this.request<HubQueryGetErasmusInfoResponse, RpcStatus>({
      path: `/hub/hub/erasmus_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/hub/hub/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<HubQueryParamsResponse, RpcStatus>({
      path: `/hub/hub/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPersonalInfo
   * @summary Queries a PersonalInfo by index.
   * @request GET:/hub/hub/personal_info
   */
  queryPersonalInfo = (params: RequestParams = {}) =>
    this.request<HubQueryGetPersonalInfoResponse, RpcStatus>({
      path: `/hub/hub/personal_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryResidenceInfo
   * @summary Queries a ResidenceInfo by index.
   * @request GET:/hub/hub/residence_info
   */
  queryResidenceInfo = (params: RequestParams = {}) =>
    this.request<HubQueryGetResidenceInfoResponse, RpcStatus>({
      path: `/hub/hub/residence_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStoredStudentAll
   * @summary Queries a list of StoredStudent items.
   * @request GET:/hub/hub/stored_student
   */
  queryStoredStudentAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<HubQueryAllStoredStudentResponse, RpcStatus>({
      path: `/hub/hub/stored_student`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStoredStudent
   * @summary Queries a StoredStudent by index.
   * @request GET:/hub/hub/stored_student/{index}
   */
  queryStoredStudent = (index: string, params: RequestParams = {}) =>
    this.request<HubQueryGetStoredStudentResponse, RpcStatus>({
      path: `/hub/hub/stored_student/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStudentInfo
   * @summary Queries a StudentInfo by index.
   * @request GET:/hub/hub/student_info
   */
  queryStudentInfo = (params: RequestParams = {}) =>
    this.request<HubQueryGetStudentInfoResponse, RpcStatus>({
      path: `/hub/hub/student_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTaxesInfo
   * @summary Queries a TaxesInfo by index.
   * @request GET:/hub/hub/taxes_info
   */
  queryTaxesInfo = (params: RequestParams = {}) =>
    this.request<HubQueryGetTaxesInfoResponse, RpcStatus>({
      path: `/hub/hub/taxes_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTranscriptOfRecords
   * @summary Queries a TranscriptOfRecords by index.
   * @request GET:/hub/hub/transcript_of_records
   */
  queryTranscriptOfRecords = (params: RequestParams = {}) =>
    this.request<HubQueryGetTranscriptOfRecordsResponse, RpcStatus>({
      path: `/hub/hub/transcript_of_records`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUniversitiesAll
   * @summary Queries a list of Universities items.
   * @request GET:/hub/hub/universities
   */
  queryUniversitiesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<HubQueryAllUniversitiesResponse, RpcStatus>({
      path: `/hub/hub/universities`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUniversities
   * @summary Queries a Universities by index.
   * @request GET:/hub/hub/universities/{universityName}
   */
  queryUniversities = (universityName: string, params: RequestParams = {}) =>
    this.request<HubQueryGetUniversitiesResponse, RpcStatus>({
      path: `/hub/hub/universities/${universityName}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
