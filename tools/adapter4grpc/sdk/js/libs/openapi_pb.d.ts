// package: service
// file: openapi.proto

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

export class RenameCmd extends jspb.Message {
  getSrc(): string;
  setSrc(value: string): void;

  getDst(): string;
  setDst(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RenameCmd.AsObject;
  static toObject(includeInstance: boolean, msg: RenameCmd): RenameCmd.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RenameCmd, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RenameCmd;
  static deserializeBinaryFromReader(message: RenameCmd, reader: jspb.BinaryReader): RenameCmd;
}

export namespace RenameCmd {
  export type AsObject = {
    src: string,
    dst: string,
  }
}

export class MoveCmd extends jspb.Message {
  getSrc(): string;
  setSrc(value: string): void;

  getDst(): string;
  setDst(value: string): void;

  getOverride(): boolean;
  setOverride(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveCmd.AsObject;
  static toObject(includeInstance: boolean, msg: MoveCmd): MoveCmd.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveCmd, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveCmd;
  static deserializeBinaryFromReader(message: MoveCmd, reader: jspb.BinaryReader): MoveCmd;
}

export namespace MoveCmd {
  export type AsObject = {
    src: string,
    dst: string,
    override: boolean,
  }
}

export class CopyCmd extends jspb.Message {
  getSrc(): string;
  setSrc(value: string): void;

  getDst(): string;
  setDst(value: string): void;

  getOverride(): boolean;
  setOverride(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CopyCmd.AsObject;
  static toObject(includeInstance: boolean, msg: CopyCmd): CopyCmd.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CopyCmd, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CopyCmd;
  static deserializeBinaryFromReader(message: CopyCmd, reader: jspb.BinaryReader): CopyCmd;
}

export namespace CopyCmd {
  export type AsObject = {
    src: string,
    dst: string,
    override: boolean,
  }
}

export class SubmitTokenCmd extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  getOverride(): boolean;
  setOverride(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubmitTokenCmd.AsObject;
  static toObject(includeInstance: boolean, msg: SubmitTokenCmd): SubmitTokenCmd.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SubmitTokenCmd, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubmitTokenCmd;
  static deserializeBinaryFromReader(message: SubmitTokenCmd, reader: jspb.BinaryReader): SubmitTokenCmd;
}

export namespace SubmitTokenCmd {
  export type AsObject = {
    token: string,
    override: boolean,
  }
}

export class QryStreamURLCmd extends jspb.Message {
  getNodeno(): string;
  setNodeno(value: string): void;

  getToken(): string;
  setToken(value: string): void;

  getEndpoint(): string;
  setEndpoint(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QryStreamURLCmd.AsObject;
  static toObject(includeInstance: boolean, msg: QryStreamURLCmd): QryStreamURLCmd.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QryStreamURLCmd, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QryStreamURLCmd;
  static deserializeBinaryFromReader(message: QryStreamURLCmd, reader: jspb.BinaryReader): QryStreamURLCmd;
}

export namespace QryStreamURLCmd {
  export type AsObject = {
    nodeno: string,
    token: string,
    endpoint: string,
  }
}

export class TNode extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getPid(): string;
  setPid(value: string): void;

  getAddr(): string;
  setAddr(value: string): void;

  getFlag(): number;
  setFlag(value: number): void;

  getName(): string;
  setName(value: string): void;

  getSize(): number;
  setSize(value: number): void;

  getCtime(): number;
  setCtime(value: number): void;

  getMtime(): number;
  setMtime(value: number): void;

  getProps(): string;
  setProps(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TNode.AsObject;
  static toObject(includeInstance: boolean, msg: TNode): TNode.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TNode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TNode;
  static deserializeBinaryFromReader(message: TNode, reader: jspb.BinaryReader): TNode;
}

export namespace TNode {
  export type AsObject = {
    id: string,
    pid: string,
    addr: string,
    flag: number,
    name: string,
    size: number,
    ctime: number,
    mtime: number,
    props: string,
  }
}

export class StreamToken extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  getNodeno(): string;
  setNodeno(value: string): void;

  getFileid(): string;
  setFileid(value: string): void;

  getFilepath(): string;
  setFilepath(value: string): void;

  getFilesize(): number;
  setFilesize(value: number): void;

  getCtime(): number;
  setCtime(value: number): void;

  getMtime(): number;
  setMtime(value: number): void;

  getEndpoint(): string;
  setEndpoint(value: string): void;

  getType(): number;
  setType(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamToken.AsObject;
  static toObject(includeInstance: boolean, msg: StreamToken): StreamToken.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StreamToken, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamToken;
  static deserializeBinaryFromReader(message: StreamToken, reader: jspb.BinaryReader): StreamToken;
}

export namespace StreamToken {
  export type AsObject = {
    token: string,
    nodeno: string,
    fileid: string,
    filepath: string,
    filesize: number,
    ctime: number,
    mtime: number,
    endpoint: string,
    type: number,
  }
}

export class SrcAndDstBo extends jspb.Message {
  getSrc(): string;
  setSrc(value: string): void;

  getDst(): string;
  setDst(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SrcAndDstBo.AsObject;
  static toObject(includeInstance: boolean, msg: SrcAndDstBo): SrcAndDstBo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SrcAndDstBo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SrcAndDstBo;
  static deserializeBinaryFromReader(message: SrcAndDstBo, reader: jspb.BinaryReader): SrcAndDstBo;
}

export namespace SrcAndDstBo {
  export type AsObject = {
    src: string,
    dst: string,
  }
}

export class CopyNodeBo extends jspb.Message {
  getSrc(): string;
  setSrc(value: string): void;

  getDst(): string;
  setDst(value: string): void;

  getOverride(): boolean;
  setOverride(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CopyNodeBo.AsObject;
  static toObject(includeInstance: boolean, msg: CopyNodeBo): CopyNodeBo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CopyNodeBo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CopyNodeBo;
  static deserializeBinaryFromReader(message: CopyNodeBo, reader: jspb.BinaryReader): CopyNodeBo;
}

export namespace CopyNodeBo {
  export type AsObject = {
    src: string,
    dst: string,
    override: boolean,
  }
}

export class MoveNodeBo extends jspb.Message {
  getSrc(): string;
  setSrc(value: string): void;

  getDst(): string;
  setDst(value: string): void;

  getOverride(): boolean;
  setOverride(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveNodeBo.AsObject;
  static toObject(includeInstance: boolean, msg: MoveNodeBo): MoveNodeBo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveNodeBo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveNodeBo;
  static deserializeBinaryFromReader(message: MoveNodeBo, reader: jspb.BinaryReader): MoveNodeBo;
}

export namespace MoveNodeBo {
  export type AsObject = {
    src: string,
    dst: string,
    override: boolean,
  }
}

export class WriteTokenBo extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  getOverride(): boolean;
  setOverride(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WriteTokenBo.AsObject;
  static toObject(includeInstance: boolean, msg: WriteTokenBo): WriteTokenBo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WriteTokenBo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WriteTokenBo;
  static deserializeBinaryFromReader(message: WriteTokenBo, reader: jspb.BinaryReader): WriteTokenBo;
}

export namespace WriteTokenBo {
  export type AsObject = {
    token: string,
    override: boolean,
  }
}

export class LimitQueryBo extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): void;

  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LimitQueryBo.AsObject;
  static toObject(includeInstance: boolean, msg: LimitQueryBo): LimitQueryBo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LimitQueryBo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LimitQueryBo;
  static deserializeBinaryFromReader(message: LimitQueryBo, reader: jspb.BinaryReader): LimitQueryBo;
}

export namespace LimitQueryBo {
  export type AsObject = {
    query: string,
    limit: number,
    offset: number,
  }
}

export class CreateUserBo extends jspb.Message {
  getUsertype(): number;
  setUsertype(value: number): void;

  getUserid(): string;
  setUserid(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getUserpwd(): string;
  setUserpwd(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserBo.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserBo): CreateUserBo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUserBo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserBo;
  static deserializeBinaryFromReader(message: CreateUserBo, reader: jspb.BinaryReader): CreateUserBo;
}

export namespace CreateUserBo {
  export type AsObject = {
    usertype: number,
    userid: string,
    username: string,
    userpwd: string,
  }
}

export class PwdUpdateBo extends jspb.Message {
  getUser(): string;
  setUser(value: string): void;

  getPwd(): string;
  setPwd(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PwdUpdateBo.AsObject;
  static toObject(includeInstance: boolean, msg: PwdUpdateBo): PwdUpdateBo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PwdUpdateBo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PwdUpdateBo;
  static deserializeBinaryFromReader(message: PwdUpdateBo, reader: jspb.BinaryReader): PwdUpdateBo;
}

export namespace PwdUpdateBo {
  export type AsObject = {
    user: string,
    pwd: string,
  }
}

export class UserListDto extends jspb.Message {
  clearDatasList(): void;
  getDatasList(): Array<UserInfo>;
  setDatasList(value: Array<UserInfo>): void;
  addDatas(value?: UserInfo, index?: number): UserInfo;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserListDto.AsObject;
  static toObject(includeInstance: boolean, msg: UserListDto): UserListDto.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserListDto, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserListDto;
  static deserializeBinaryFromReader(message: UserListDto, reader: jspb.BinaryReader): UserListDto;
}

export namespace UserListDto {
  export type AsObject = {
    datasList: Array<UserInfo.AsObject>,
    total: number,
  }
}

export class UserInfo extends jspb.Message {
  getUsertype(): number;
  setUsertype(value: number): void;

  getUserid(): string;
  setUserid(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getCttime(): number;
  setCttime(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserInfo.AsObject;
  static toObject(includeInstance: boolean, msg: UserInfo): UserInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserInfo;
  static deserializeBinaryFromReader(message: UserInfo, reader: jspb.BinaryReader): UserInfo;
}

export namespace UserInfo {
  export type AsObject = {
    usertype: number,
    userid: string,
    username: string,
    cttime: number,
  }
}

export class DirNameListDto extends jspb.Message {
  clearDatasList(): void;
  getDatasList(): Array<string>;
  setDatasList(value: Array<string>): void;
  addDatas(value: string, index?: number): string;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DirNameListDto.AsObject;
  static toObject(includeInstance: boolean, msg: DirNameListDto): DirNameListDto.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DirNameListDto, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DirNameListDto;
  static deserializeBinaryFromReader(message: DirNameListDto, reader: jspb.BinaryReader): DirNameListDto;
}

export namespace DirNameListDto {
  export type AsObject = {
    datasList: Array<string>,
    total: number,
  }
}

export class DirNodeListDto extends jspb.Message {
  clearDatasList(): void;
  getDatasList(): Array<TNode>;
  setDatasList(value: Array<TNode>): void;
  addDatas(value?: TNode, index?: number): TNode;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DirNodeListDto.AsObject;
  static toObject(includeInstance: boolean, msg: DirNodeListDto): DirNodeListDto.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DirNodeListDto, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DirNodeListDto;
  static deserializeBinaryFromReader(message: DirNodeListDto, reader: jspb.BinaryReader): DirNodeListDto;
}

export namespace DirNodeListDto {
  export type AsObject = {
    datasList: Array<TNode.AsObject>,
    total: number,
  }
}

