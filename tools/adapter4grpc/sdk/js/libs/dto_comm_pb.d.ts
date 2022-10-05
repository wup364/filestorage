// package: dto
// file: dto_comm.proto

import * as jspb from 'google-protobuf';

export class QryOfString extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QryOfString.AsObject;
  static toObject(includeInstance: boolean, msg: QryOfString): QryOfString.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QryOfString, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QryOfString;
  static deserializeBinaryFromReader(message: QryOfString, reader: jspb.BinaryReader): QryOfString;
}

export namespace QryOfString {
  export type AsObject = {
    query: string,
  }
}

export class QueryLimitOfString extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): void;

  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryLimitOfString.AsObject;
  static toObject(includeInstance: boolean, msg: QueryLimitOfString): QueryLimitOfString.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QueryLimitOfString, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryLimitOfString;
  static deserializeBinaryFromReader(message: QueryLimitOfString, reader: jspb.BinaryReader): QueryLimitOfString;
}

export namespace QueryLimitOfString {
  export type AsObject = {
    query: string,
    limit: number,
    offset: number,
  }
}

export class QryOfStrings extends jspb.Message {
  clearQueryList(): void;
  getQueryList(): Array<string>;
  setQueryList(value: Array<string>): void;
  addQuery(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QryOfStrings.AsObject;
  static toObject(includeInstance: boolean, msg: QryOfStrings): QryOfStrings.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QryOfStrings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QryOfStrings;
  static deserializeBinaryFromReader(message: QryOfStrings, reader: jspb.BinaryReader): QryOfStrings;
}

export namespace QryOfStrings {
  export type AsObject = {
    queryList: Array<string>,
  }
}

export class ResultOfBool extends jspb.Message {
  getResult(): boolean;
  setResult(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResultOfBool.AsObject;
  static toObject(includeInstance: boolean, msg: ResultOfBool): ResultOfBool.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResultOfBool, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResultOfBool;
  static deserializeBinaryFromReader(message: ResultOfBool, reader: jspb.BinaryReader): ResultOfBool;
}

export namespace ResultOfBool {
  export type AsObject = {
    result: boolean,
    message: string,
  }
}

export class ResultOfInt64 extends jspb.Message {
  getResult(): number;
  setResult(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResultOfInt64.AsObject;
  static toObject(includeInstance: boolean, msg: ResultOfInt64): ResultOfInt64.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResultOfInt64, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResultOfInt64;
  static deserializeBinaryFromReader(message: ResultOfInt64, reader: jspb.BinaryReader): ResultOfInt64;
}

export namespace ResultOfInt64 {
  export type AsObject = {
    result: number,
  }
}

export class ResultOfString extends jspb.Message {
  getResult(): string;
  setResult(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResultOfString.AsObject;
  static toObject(includeInstance: boolean, msg: ResultOfString): ResultOfString.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResultOfString, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResultOfString;
  static deserializeBinaryFromReader(message: ResultOfString, reader: jspb.BinaryReader): ResultOfString;
}

export namespace ResultOfString {
  export type AsObject = {
    result: string,
  }
}

