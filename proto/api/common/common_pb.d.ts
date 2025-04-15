import * as jspb from 'google-protobuf'



export class PaginationResponse extends jspb.Message {
  getPage(): number;
  setPage(value: number): PaginationResponse;

  getTotal(): number;
  setTotal(value: number): PaginationResponse;

  getOffset(): number;
  setOffset(value: number): PaginationResponse;

  getLimit(): number;
  setLimit(value: number): PaginationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PaginationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PaginationResponse): PaginationResponse.AsObject;
  static serializeBinaryToWriter(message: PaginationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PaginationResponse;
  static deserializeBinaryFromReader(message: PaginationResponse, reader: jspb.BinaryReader): PaginationResponse;
}

export namespace PaginationResponse {
  export type AsObject = {
    page: number,
    total: number,
    offset: number,
    limit: number,
  }
}

