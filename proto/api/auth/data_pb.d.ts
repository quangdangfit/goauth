import * as jspb from 'google-protobuf'

import * as validate_validate_pb from '../validate/validate_pb'; // proto import: "validate/validate.proto"
import * as common_common_pb from '../common/common_pb'; // proto import: "common/common.proto"
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb'; // proto import: "google/protobuf/struct.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class User extends jspb.Message {
  getId(): number;
  setId(value: number): User;

  getUsername(): string;
  setUsername(value: string): User;

  getName(): string;
  setName(value: string): User;

  getPhone(): string;
  setPhone(value: string): User;

  getEmail(): string;
  setEmail(value: string): User;

  getVerifiedPhone(): boolean;
  setVerifiedPhone(value: boolean): User;

  getVerifiedEmail(): boolean;
  setVerifiedEmail(value: boolean): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    username: string,
    name: string,
    phone: string,
    email: string,
    verifiedPhone: boolean,
    verifiedEmail: boolean,
  }
}

export class Permission extends jspb.Message {
  getName(): string;
  setName(value: string): Permission;

  getGroup(): string;
  setGroup(value: string): Permission;

  getParent(): string;
  setParent(value: string): Permission;

  getApiMethod(): string;
  setApiMethod(value: string): Permission;

  getApiEndpoint(): string;
  setApiEndpoint(value: string): Permission;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Permission.AsObject;
  static toObject(includeInstance: boolean, msg: Permission): Permission.AsObject;
  static serializeBinaryToWriter(message: Permission, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Permission;
  static deserializeBinaryFromReader(message: Permission, reader: jspb.BinaryReader): Permission;
}

export namespace Permission {
  export type AsObject = {
    name: string,
    group: string,
    parent: string,
    apiMethod: string,
    apiEndpoint: string,
  }
}

export class RegisterRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): RegisterRequest;

  getPassword(): string;
  setPassword(value: string): RegisterRequest;

  getName(): string;
  setName(value: string): RegisterRequest;

  getPhone(): string;
  setPhone(value: string): RegisterRequest;

  getEmail(): string;
  setEmail(value: string): RegisterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterRequest;
  static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
  export type AsObject = {
    username: string,
    password: string,
    name: string,
    phone: string,
    email: string,
  }
}

export class RegisterResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): RegisterResponse;

  getMessage(): string;
  setMessage(value: string): RegisterResponse;

  getData(): User | undefined;
  setData(value?: User): RegisterResponse;
  hasData(): boolean;
  clearData(): RegisterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterResponse): RegisterResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterResponse;
  static deserializeBinaryFromReader(message: RegisterResponse, reader: jspb.BinaryReader): RegisterResponse;
}

export namespace RegisterResponse {
  export type AsObject = {
    code: number,
    message: string,
    data?: User.AsObject,
  }
}

export class LoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): LoginResponse;

  getMessage(): string;
  setMessage(value: string): LoginResponse;

  getData(): LoginResponse.Data | undefined;
  setData(value?: LoginResponse.Data): LoginResponse;
  hasData(): boolean;
  clearData(): LoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
  static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginResponse;
  static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
  export type AsObject = {
    code: number,
    message: string,
    data?: LoginResponse.Data.AsObject,
  }

  export class Data extends jspb.Message {
    getAccessToken(): string;
    setAccessToken(value: string): Data;

    getRefreshToken(): string;
    setRefreshToken(value: string): Data;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Data.AsObject;
    static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
    static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Data;
    static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
  }

  export namespace Data {
    export type AsObject = {
      accessToken: string,
      refreshToken: string,
    }
  }

}

export class RefreshTokenRequest extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): RefreshTokenRequest;

  getRefreshToken(): string;
  setRefreshToken(value: string): RefreshTokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshTokenRequest): RefreshTokenRequest.AsObject;
  static serializeBinaryToWriter(message: RefreshTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshTokenRequest;
  static deserializeBinaryFromReader(message: RefreshTokenRequest, reader: jspb.BinaryReader): RefreshTokenRequest;
}

export namespace RefreshTokenRequest {
  export type AsObject = {
    accessToken: string,
    refreshToken: string,
  }
}

export class RefreshTokenResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): RefreshTokenResponse;

  getMessage(): string;
  setMessage(value: string): RefreshTokenResponse;

  getData(): RefreshTokenResponse.Data | undefined;
  setData(value?: RefreshTokenResponse.Data): RefreshTokenResponse;
  hasData(): boolean;
  clearData(): RefreshTokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshTokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshTokenResponse): RefreshTokenResponse.AsObject;
  static serializeBinaryToWriter(message: RefreshTokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshTokenResponse;
  static deserializeBinaryFromReader(message: RefreshTokenResponse, reader: jspb.BinaryReader): RefreshTokenResponse;
}

export namespace RefreshTokenResponse {
  export type AsObject = {
    code: number,
    message: string,
    data?: RefreshTokenResponse.Data.AsObject,
  }

  export class Data extends jspb.Message {
    getAccessToken(): string;
    setAccessToken(value: string): Data;

    getRefreshToken(): string;
    setRefreshToken(value: string): Data;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Data.AsObject;
    static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
    static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Data;
    static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
  }

  export namespace Data {
    export type AsObject = {
      accessToken: string,
      refreshToken: string,
    }
  }

}

export class ChangePasswordRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): ChangePasswordRequest;

  getPassword(): string;
  setPassword(value: string): ChangePasswordRequest;

  getNewPassword(): string;
  setNewPassword(value: string): ChangePasswordRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChangePasswordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChangePasswordRequest): ChangePasswordRequest.AsObject;
  static serializeBinaryToWriter(message: ChangePasswordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChangePasswordRequest;
  static deserializeBinaryFromReader(message: ChangePasswordRequest, reader: jspb.BinaryReader): ChangePasswordRequest;
}

export namespace ChangePasswordRequest {
  export type AsObject = {
    username: string,
    password: string,
    newPassword: string,
  }
}

export class ChangePasswordResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): ChangePasswordResponse;

  getMessage(): string;
  setMessage(value: string): ChangePasswordResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChangePasswordResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChangePasswordResponse): ChangePasswordResponse.AsObject;
  static serializeBinaryToWriter(message: ChangePasswordResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChangePasswordResponse;
  static deserializeBinaryFromReader(message: ChangePasswordResponse, reader: jspb.BinaryReader): ChangePasswordResponse;
}

export namespace ChangePasswordResponse {
  export type AsObject = {
    code: number,
    message: string,
  }
}

export class ResetPasswordRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): ResetPasswordRequest;

  getMethod(): VerificationMethod;
  setMethod(value: VerificationMethod): ResetPasswordRequest;

  getRefCode(): string;
  setRefCode(value: string): ResetPasswordRequest;

  getVerificationCode(): string;
  setVerificationCode(value: string): ResetPasswordRequest;

  getNewPassword(): string;
  setNewPassword(value: string): ResetPasswordRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResetPasswordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ResetPasswordRequest): ResetPasswordRequest.AsObject;
  static serializeBinaryToWriter(message: ResetPasswordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResetPasswordRequest;
  static deserializeBinaryFromReader(message: ResetPasswordRequest, reader: jspb.BinaryReader): ResetPasswordRequest;
}

export namespace ResetPasswordRequest {
  export type AsObject = {
    username: string,
    method: VerificationMethod,
    refCode: string,
    verificationCode: string,
    newPassword: string,
  }
}

export class ResetPasswordResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): ResetPasswordResponse;

  getMessage(): string;
  setMessage(value: string): ResetPasswordResponse;

  getData(): ResetPasswordResponse.Data | undefined;
  setData(value?: ResetPasswordResponse.Data): ResetPasswordResponse;
  hasData(): boolean;
  clearData(): ResetPasswordResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResetPasswordResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ResetPasswordResponse): ResetPasswordResponse.AsObject;
  static serializeBinaryToWriter(message: ResetPasswordResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResetPasswordResponse;
  static deserializeBinaryFromReader(message: ResetPasswordResponse, reader: jspb.BinaryReader): ResetPasswordResponse;
}

export namespace ResetPasswordResponse {
  export type AsObject = {
    code: number,
    message: string,
    data?: ResetPasswordResponse.Data.AsObject,
  }

  export class Data extends jspb.Message {
    getRefCode(): string;
    setRefCode(value: string): Data;
    hasRefCode(): boolean;
    clearRefCode(): Data;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Data.AsObject;
    static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
    static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Data;
    static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
  }

  export namespace Data {
    export type AsObject = {
      refCode?: string,
    }

    export enum RefCodeCase { 
      _REF_CODE_NOT_SET = 0,
      REF_CODE = 1,
    }
  }

}

export class UpdateUserProfileRequest extends jspb.Message {
  getName(): string;
  setName(value: string): UpdateUserProfileRequest;

  getPhone(): string;
  setPhone(value: string): UpdateUserProfileRequest;

  getEmail(): string;
  setEmail(value: string): UpdateUserProfileRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserProfileRequest): UpdateUserProfileRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateUserProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserProfileRequest;
  static deserializeBinaryFromReader(message: UpdateUserProfileRequest, reader: jspb.BinaryReader): UpdateUserProfileRequest;
}

export namespace UpdateUserProfileRequest {
  export type AsObject = {
    name: string,
    phone: string,
    email: string,
  }
}

export class UpdateUserProfileResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): UpdateUserProfileResponse;

  getMessage(): string;
  setMessage(value: string): UpdateUserProfileResponse;

  getData(): User | undefined;
  setData(value?: User): UpdateUserProfileResponse;
  hasData(): boolean;
  clearData(): UpdateUserProfileResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserProfileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserProfileResponse): UpdateUserProfileResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateUserProfileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserProfileResponse;
  static deserializeBinaryFromReader(message: UpdateUserProfileResponse, reader: jspb.BinaryReader): UpdateUserProfileResponse;
}

export namespace UpdateUserProfileResponse {
  export type AsObject = {
    code: number,
    message: string,
    data?: User.AsObject,
  }
}

export class GetUserProfileRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserProfileRequest): GetUserProfileRequest.AsObject;
  static serializeBinaryToWriter(message: GetUserProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserProfileRequest;
  static deserializeBinaryFromReader(message: GetUserProfileRequest, reader: jspb.BinaryReader): GetUserProfileRequest;
}

export namespace GetUserProfileRequest {
  export type AsObject = {
  }
}

export class GetUserProfileResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): GetUserProfileResponse;

  getMessage(): string;
  setMessage(value: string): GetUserProfileResponse;

  getData(): User | undefined;
  setData(value?: User): GetUserProfileResponse;
  hasData(): boolean;
  clearData(): GetUserProfileResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserProfileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserProfileResponse): GetUserProfileResponse.AsObject;
  static serializeBinaryToWriter(message: GetUserProfileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserProfileResponse;
  static deserializeBinaryFromReader(message: GetUserProfileResponse, reader: jspb.BinaryReader): GetUserProfileResponse;
}

export namespace GetUserProfileResponse {
  export type AsObject = {
    code: number,
    message: string,
    data?: User.AsObject,
  }
}

export class VerifyAccountRequest extends jspb.Message {
  getMethod(): VerificationMethod;
  setMethod(value: VerificationMethod): VerifyAccountRequest;

  getRefCode(): string;
  setRefCode(value: string): VerifyAccountRequest;

  getVerificationCode(): string;
  setVerificationCode(value: string): VerifyAccountRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VerifyAccountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: VerifyAccountRequest): VerifyAccountRequest.AsObject;
  static serializeBinaryToWriter(message: VerifyAccountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VerifyAccountRequest;
  static deserializeBinaryFromReader(message: VerifyAccountRequest, reader: jspb.BinaryReader): VerifyAccountRequest;
}

export namespace VerifyAccountRequest {
  export type AsObject = {
    method: VerificationMethod,
    refCode: string,
    verificationCode: string,
  }
}

export class VerifyAccountResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): VerifyAccountResponse;

  getMessage(): string;
  setMessage(value: string): VerifyAccountResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VerifyAccountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: VerifyAccountResponse): VerifyAccountResponse.AsObject;
  static serializeBinaryToWriter(message: VerifyAccountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VerifyAccountResponse;
  static deserializeBinaryFromReader(message: VerifyAccountResponse, reader: jspb.BinaryReader): VerifyAccountResponse;
}

export namespace VerifyAccountResponse {
  export type AsObject = {
    code: number,
    message: string,
  }

  export class Data extends jspb.Message {
    getRefCode(): string;
    setRefCode(value: string): Data;
    hasRefCode(): boolean;
    clearRefCode(): Data;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Data.AsObject;
    static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
    static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Data;
    static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
  }

  export namespace Data {
    export type AsObject = {
      refCode?: string,
    }

    export enum RefCodeCase { 
      _REF_CODE_NOT_SET = 0,
      REF_CODE = 1,
    }
  }

}

export class DeactivateAccountRequest extends jspb.Message {
  getMethod(): VerificationMethod;
  setMethod(value: VerificationMethod): DeactivateAccountRequest;

  getRefCode(): string;
  setRefCode(value: string): DeactivateAccountRequest;

  getVerificationCode(): string;
  setVerificationCode(value: string): DeactivateAccountRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeactivateAccountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeactivateAccountRequest): DeactivateAccountRequest.AsObject;
  static serializeBinaryToWriter(message: DeactivateAccountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeactivateAccountRequest;
  static deserializeBinaryFromReader(message: DeactivateAccountRequest, reader: jspb.BinaryReader): DeactivateAccountRequest;
}

export namespace DeactivateAccountRequest {
  export type AsObject = {
    method: VerificationMethod,
    refCode: string,
    verificationCode: string,
  }
}

export class DeactivateAccountResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): DeactivateAccountResponse;

  getMessage(): string;
  setMessage(value: string): DeactivateAccountResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeactivateAccountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeactivateAccountResponse): DeactivateAccountResponse.AsObject;
  static serializeBinaryToWriter(message: DeactivateAccountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeactivateAccountResponse;
  static deserializeBinaryFromReader(message: DeactivateAccountResponse, reader: jspb.BinaryReader): DeactivateAccountResponse;
}

export namespace DeactivateAccountResponse {
  export type AsObject = {
    code: number,
    message: string,
  }

  export class Data extends jspb.Message {
    getRefCode(): string;
    setRefCode(value: string): Data;
    hasRefCode(): boolean;
    clearRefCode(): Data;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Data.AsObject;
    static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
    static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Data;
    static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
  }

  export namespace Data {
    export type AsObject = {
      refCode?: string,
    }

    export enum RefCodeCase { 
      _REF_CODE_NOT_SET = 0,
      REF_CODE = 1,
    }
  }

}

export class CreatePermissionRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreatePermissionRequest;

  getGroup(): string;
  setGroup(value: string): CreatePermissionRequest;

  getParent(): string;
  setParent(value: string): CreatePermissionRequest;

  getApiMethod(): string;
  setApiMethod(value: string): CreatePermissionRequest;

  getApiEndpoint(): string;
  setApiEndpoint(value: string): CreatePermissionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePermissionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePermissionRequest): CreatePermissionRequest.AsObject;
  static serializeBinaryToWriter(message: CreatePermissionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePermissionRequest;
  static deserializeBinaryFromReader(message: CreatePermissionRequest, reader: jspb.BinaryReader): CreatePermissionRequest;
}

export namespace CreatePermissionRequest {
  export type AsObject = {
    name: string,
    group: string,
    parent: string,
    apiMethod: string,
    apiEndpoint: string,
  }
}

export class CreatePermissionResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): CreatePermissionResponse;

  getMessage(): string;
  setMessage(value: string): CreatePermissionResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePermissionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePermissionResponse): CreatePermissionResponse.AsObject;
  static serializeBinaryToWriter(message: CreatePermissionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePermissionResponse;
  static deserializeBinaryFromReader(message: CreatePermissionResponse, reader: jspb.BinaryReader): CreatePermissionResponse;
}

export namespace CreatePermissionResponse {
  export type AsObject = {
    code: number,
    message: string,
  }
}

export class ListPermissionRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPermissionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPermissionRequest): ListPermissionRequest.AsObject;
  static serializeBinaryToWriter(message: ListPermissionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPermissionRequest;
  static deserializeBinaryFromReader(message: ListPermissionRequest, reader: jspb.BinaryReader): ListPermissionRequest;
}

export namespace ListPermissionRequest {
  export type AsObject = {
  }
}

export class ListPermissionResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): ListPermissionResponse;

  getMessage(): string;
  setMessage(value: string): ListPermissionResponse;

  getDataList(): Array<Permission>;
  setDataList(value: Array<Permission>): ListPermissionResponse;
  clearDataList(): ListPermissionResponse;
  addData(value?: Permission, index?: number): Permission;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPermissionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPermissionResponse): ListPermissionResponse.AsObject;
  static serializeBinaryToWriter(message: ListPermissionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPermissionResponse;
  static deserializeBinaryFromReader(message: ListPermissionResponse, reader: jspb.BinaryReader): ListPermissionResponse;
}

export namespace ListPermissionResponse {
  export type AsObject = {
    code: number,
    message: string,
    dataList: Array<Permission.AsObject>,
  }
}

export class CreateUserRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): CreateUserRequest;

  getPassword(): string;
  setPassword(value: string): CreateUserRequest;

  getName(): string;
  setName(value: string): CreateUserRequest;

  getPhone(): string;
  setPhone(value: string): CreateUserRequest;

  getEmail(): string;
  setEmail(value: string): CreateUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    username: string,
    password: string,
    name: string,
    phone: string,
    email: string,
  }
}

export class CreateUserResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): CreateUserResponse;

  getMessage(): string;
  setMessage(value: string): CreateUserResponse;

  getData(): User | undefined;
  setData(value?: User): CreateUserResponse;
  hasData(): boolean;
  clearData(): CreateUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
  static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
  static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
  export type AsObject = {
    code: number,
    message: string,
    data?: User.AsObject,
  }
}

export class ListUserRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListUserRequest): ListUserRequest.AsObject;
  static serializeBinaryToWriter(message: ListUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListUserRequest;
  static deserializeBinaryFromReader(message: ListUserRequest, reader: jspb.BinaryReader): ListUserRequest;
}

export namespace ListUserRequest {
  export type AsObject = {
  }
}

export class ListUserResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): ListUserResponse;

  getMessage(): string;
  setMessage(value: string): ListUserResponse;

  getDataList(): Array<User>;
  setDataList(value: Array<User>): ListUserResponse;
  clearDataList(): ListUserResponse;
  addData(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListUserResponse): ListUserResponse.AsObject;
  static serializeBinaryToWriter(message: ListUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListUserResponse;
  static deserializeBinaryFromReader(message: ListUserResponse, reader: jspb.BinaryReader): ListUserResponse;
}

export namespace ListUserResponse {
  export type AsObject = {
    code: number,
    message: string,
    dataList: Array<User.AsObject>,
  }
}

export class RegisterDeviceRequest extends jspb.Message {
  getDeviceId(): string;
  setDeviceId(value: string): RegisterDeviceRequest;

  getDeviceName(): string;
  setDeviceName(value: string): RegisterDeviceRequest;

  getFcmToken(): string;
  setFcmToken(value: string): RegisterDeviceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterDeviceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterDeviceRequest): RegisterDeviceRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterDeviceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterDeviceRequest;
  static deserializeBinaryFromReader(message: RegisterDeviceRequest, reader: jspb.BinaryReader): RegisterDeviceRequest;
}

export namespace RegisterDeviceRequest {
  export type AsObject = {
    deviceId: string,
    deviceName: string,
    fcmToken: string,
  }
}

export class RegisterDeviceResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): RegisterDeviceResponse;

  getMessage(): string;
  setMessage(value: string): RegisterDeviceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterDeviceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterDeviceResponse): RegisterDeviceResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterDeviceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterDeviceResponse;
  static deserializeBinaryFromReader(message: RegisterDeviceResponse, reader: jspb.BinaryReader): RegisterDeviceResponse;
}

export namespace RegisterDeviceResponse {
  export type AsObject = {
    code: number,
    message: string,
  }
}

export class UnregisterDeviceRequest extends jspb.Message {
  getDeviceId(): string;
  setDeviceId(value: string): UnregisterDeviceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnregisterDeviceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnregisterDeviceRequest): UnregisterDeviceRequest.AsObject;
  static serializeBinaryToWriter(message: UnregisterDeviceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnregisterDeviceRequest;
  static deserializeBinaryFromReader(message: UnregisterDeviceRequest, reader: jspb.BinaryReader): UnregisterDeviceRequest;
}

export namespace UnregisterDeviceRequest {
  export type AsObject = {
    deviceId: string,
  }
}

export class UnregisterDeviceResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): UnregisterDeviceResponse;

  getMessage(): string;
  setMessage(value: string): UnregisterDeviceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnregisterDeviceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UnregisterDeviceResponse): UnregisterDeviceResponse.AsObject;
  static serializeBinaryToWriter(message: UnregisterDeviceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnregisterDeviceResponse;
  static deserializeBinaryFromReader(message: UnregisterDeviceResponse, reader: jspb.BinaryReader): UnregisterDeviceResponse;
}

export namespace UnregisterDeviceResponse {
  export type AsObject = {
    code: number,
    message: string,
  }
}

export class ListDeviceRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDeviceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDeviceRequest): ListDeviceRequest.AsObject;
  static serializeBinaryToWriter(message: ListDeviceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDeviceRequest;
  static deserializeBinaryFromReader(message: ListDeviceRequest, reader: jspb.BinaryReader): ListDeviceRequest;
}

export namespace ListDeviceRequest {
  export type AsObject = {
  }
}

export class ListDeviceResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): ListDeviceResponse;

  getMessage(): string;
  setMessage(value: string): ListDeviceResponse;

  getDataList(): Array<ListDeviceResponse.Data>;
  setDataList(value: Array<ListDeviceResponse.Data>): ListDeviceResponse;
  clearDataList(): ListDeviceResponse;
  addData(value?: ListDeviceResponse.Data, index?: number): ListDeviceResponse.Data;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDeviceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDeviceResponse): ListDeviceResponse.AsObject;
  static serializeBinaryToWriter(message: ListDeviceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDeviceResponse;
  static deserializeBinaryFromReader(message: ListDeviceResponse, reader: jspb.BinaryReader): ListDeviceResponse;
}

export namespace ListDeviceResponse {
  export type AsObject = {
    code: number,
    message: string,
    dataList: Array<ListDeviceResponse.Data.AsObject>,
  }

  export class Data extends jspb.Message {
    getDeviceId(): string;
    setDeviceId(value: string): Data;

    getDeviceName(): string;
    setDeviceName(value: string): Data;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Data.AsObject;
    static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
    static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Data;
    static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
  }

  export namespace Data {
    export type AsObject = {
      deviceId: string,
      deviceName: string,
    }
  }

}

export enum VerificationMethod { 
  VERIFICATIONMETHODEMAIL = 0,
  VERIFICATIONMETHODPHONE = 1,
}
