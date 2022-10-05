// package: service
// file: openapi.proto

import * as grpc from '@grpc/grpc-js';
import * as openapi_pb from './openapi_pb';

interface IOpenApiService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  isDir: IOpenApiService_IIsDir;
  isFile: IOpenApiService_IIsFile;
  isExist: IOpenApiService_IIsExist;
  getFileSize: IOpenApiService_IGetFileSize;
  getNode: IOpenApiService_IGetNode;
  getNodes: IOpenApiService_IGetNodes;
  getDirNameList: IOpenApiService_IGetDirNameList;
  getDirNodeList: IOpenApiService_IGetDirNodeList;
  doMkDir: IOpenApiService_IDoMkDir;
  doDelete: IOpenApiService_IDoDelete;
  doRename: IOpenApiService_IDoRename;
  doCopy: IOpenApiService_IDoCopy;
  doMove: IOpenApiService_IDoMove;
  doQueryToken: IOpenApiService_IDoQueryToken;
  doAskReadToken: IOpenApiService_IDoAskReadToken;
  doAskWriteToken: IOpenApiService_IDoAskWriteToken;
  doRefreshToken: IOpenApiService_IDoRefreshToken;
  doSubmitWriteToken: IOpenApiService_IDoSubmitWriteToken;
  getReadStreamURL: IOpenApiService_IGetReadStreamURL;
  getWriteStreamURL: IOpenApiService_IGetWriteStreamURL;
}

interface IOpenApiService_IIsDir extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.ResultOfBool> {
  path: '/service.OpenApi/IsDir'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfBool>;
}

interface IOpenApiService_IIsFile extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.ResultOfBool> {
  path: '/service.OpenApi/IsFile'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfBool>;
}

interface IOpenApiService_IIsExist extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.ResultOfBool> {
  path: '/service.OpenApi/IsExist'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfBool>;
}

interface IOpenApiService_IGetFileSize extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.ResultOfInt64> {
  path: '/service.OpenApi/GetFileSize'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfInt64>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfInt64>;
}

interface IOpenApiService_IGetNode extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.TNode> {
  path: '/service.OpenApi/GetNode'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.TNode>;
  responseDeserialize: grpc.deserialize<openapi_pb.TNode>;
}

interface IOpenApiService_IGetNodes extends grpc.MethodDefinition<openapi_pb.QryOfStrings, openapi_pb.DirNodeListDto> {
  path: '/service.OpenApi/GetNodes'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfStrings>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfStrings>;
  responseSerialize: grpc.serialize<openapi_pb.DirNodeListDto>;
  responseDeserialize: grpc.deserialize<openapi_pb.DirNodeListDto>;
}

interface IOpenApiService_IGetDirNameList extends grpc.MethodDefinition<openapi_pb.QueryLimitOfString, openapi_pb.DirNameListDto> {
  path: '/service.OpenApi/GetDirNameList'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QueryLimitOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QueryLimitOfString>;
  responseSerialize: grpc.serialize<openapi_pb.DirNameListDto>;
  responseDeserialize: grpc.deserialize<openapi_pb.DirNameListDto>;
}

interface IOpenApiService_IGetDirNodeList extends grpc.MethodDefinition<openapi_pb.QueryLimitOfString, openapi_pb.DirNodeListDto> {
  path: '/service.OpenApi/GetDirNodeList'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QueryLimitOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QueryLimitOfString>;
  responseSerialize: grpc.serialize<openapi_pb.DirNodeListDto>;
  responseDeserialize: grpc.deserialize<openapi_pb.DirNodeListDto>;
}

interface IOpenApiService_IDoMkDir extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.ResultOfString> {
  path: '/service.OpenApi/DoMkDir'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfString>;
}

interface IOpenApiService_IDoDelete extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.ResultOfBool> {
  path: '/service.OpenApi/DoDelete'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfBool>;
}

interface IOpenApiService_IDoRename extends grpc.MethodDefinition<openapi_pb.RenameCmd, openapi_pb.ResultOfBool> {
  path: '/service.OpenApi/DoRename'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.RenameCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.RenameCmd>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfBool>;
}

interface IOpenApiService_IDoCopy extends grpc.MethodDefinition<openapi_pb.MoveCmd, openapi_pb.ResultOfString> {
  path: '/service.OpenApi/DoCopy'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.MoveCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.MoveCmd>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfString>;
}

interface IOpenApiService_IDoMove extends grpc.MethodDefinition<openapi_pb.CopyCmd, openapi_pb.ResultOfBool> {
  path: '/service.OpenApi/DoMove'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.CopyCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.CopyCmd>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfBool>;
}

interface IOpenApiService_IDoQueryToken extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoQueryToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.StreamToken>;
  responseDeserialize: grpc.deserialize<openapi_pb.StreamToken>;
}

interface IOpenApiService_IDoAskReadToken extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoAskReadToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.StreamToken>;
  responseDeserialize: grpc.deserialize<openapi_pb.StreamToken>;
}

interface IOpenApiService_IDoAskWriteToken extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoAskWriteToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.StreamToken>;
  responseDeserialize: grpc.deserialize<openapi_pb.StreamToken>;
}

interface IOpenApiService_IDoRefreshToken extends grpc.MethodDefinition<openapi_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoRefreshToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.StreamToken>;
  responseDeserialize: grpc.deserialize<openapi_pb.StreamToken>;
}

interface IOpenApiService_IDoSubmitWriteToken extends grpc.MethodDefinition<openapi_pb.SubmitTokenCmd, openapi_pb.TNode> {
  path: '/service.OpenApi/DoSubmitWriteToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.SubmitTokenCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.SubmitTokenCmd>;
  responseSerialize: grpc.serialize<openapi_pb.TNode>;
  responseDeserialize: grpc.deserialize<openapi_pb.TNode>;
}

interface IOpenApiService_IGetReadStreamURL extends grpc.MethodDefinition<openapi_pb.QryStreamURLCmd, openapi_pb.ResultOfString> {
  path: '/service.OpenApi/GetReadStreamURL'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryStreamURLCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryStreamURLCmd>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfString>;
}

interface IOpenApiService_IGetWriteStreamURL extends grpc.MethodDefinition<openapi_pb.QryStreamURLCmd, openapi_pb.ResultOfString> {
  path: '/service.OpenApi/GetWriteStreamURL'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryStreamURLCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryStreamURLCmd>;
  responseSerialize: grpc.serialize<openapi_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<openapi_pb.ResultOfString>;
}

export const OpenApiService: IOpenApiService;
export interface IOpenApiServer extends grpc.UntypedServiceImplementation {
  isDir: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.ResultOfBool>;
  isFile: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.ResultOfBool>;
  isExist: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.ResultOfBool>;
  getFileSize: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.ResultOfInt64>;
  getNode: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.TNode>;
  getNodes: grpc.handleUnaryCall<openapi_pb.QryOfStrings, openapi_pb.DirNodeListDto>;
  getDirNameList: grpc.handleUnaryCall<openapi_pb.QueryLimitOfString, openapi_pb.DirNameListDto>;
  getDirNodeList: grpc.handleUnaryCall<openapi_pb.QueryLimitOfString, openapi_pb.DirNodeListDto>;
  doMkDir: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.ResultOfString>;
  doDelete: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.ResultOfBool>;
  doRename: grpc.handleUnaryCall<openapi_pb.RenameCmd, openapi_pb.ResultOfBool>;
  doCopy: grpc.handleUnaryCall<openapi_pb.MoveCmd, openapi_pb.ResultOfString>;
  doMove: grpc.handleUnaryCall<openapi_pb.CopyCmd, openapi_pb.ResultOfBool>;
  doQueryToken: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.StreamToken>;
  doAskReadToken: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.StreamToken>;
  doAskWriteToken: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.StreamToken>;
  doRefreshToken: grpc.handleUnaryCall<openapi_pb.QryOfString, openapi_pb.StreamToken>;
  doSubmitWriteToken: grpc.handleUnaryCall<openapi_pb.SubmitTokenCmd, openapi_pb.TNode>;
  getReadStreamURL: grpc.handleUnaryCall<openapi_pb.QryStreamURLCmd, openapi_pb.ResultOfString>;
  getWriteStreamURL: grpc.handleUnaryCall<openapi_pb.QryStreamURLCmd, openapi_pb.ResultOfString>;
}

export interface IOpenApiClient {
  isDir(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isFile(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isFile(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isFile(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isExist(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isExist(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isExist(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  getFileSize(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  getFileSize(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  getFileSize(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  getNode(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getNode(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getNode(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getNodes(request: openapi_pb.QryOfStrings, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getNodes(request: openapi_pb.QryOfStrings, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getNodes(request: openapi_pb.QryOfStrings, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getDirNameList(request: openapi_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  getDirNameList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  getDirNameList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  getDirNodeList(request: openapi_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getDirNodeList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getDirNodeList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  doMkDir(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doMkDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doMkDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doDelete(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doDelete(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doDelete(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doRename(request: openapi_pb.RenameCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doCopy(request: openapi_pb.MoveCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doMove(request: openapi_pb.CopyCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doQueryToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doQueryToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doQueryToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskReadToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskReadToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskReadToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskWriteToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskWriteToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskWriteToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doRefreshToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doRefreshToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doRefreshToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getReadStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
}

export class OpenApiClient extends grpc.Client implements IOpenApiClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
  public isDir(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isFile(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isFile(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isFile(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isExist(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isExist(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isExist(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public getFileSize(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  public getFileSize(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  public getFileSize(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  public getNode(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getNode(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getNode(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getNodes(request: openapi_pb.QryOfStrings, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getNodes(request: openapi_pb.QryOfStrings, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getNodes(request: openapi_pb.QryOfStrings, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getDirNameList(request: openapi_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  public getDirNameList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  public getDirNameList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  public getDirNodeList(request: openapi_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getDirNodeList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getDirNodeList(request: openapi_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public doMkDir(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doMkDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doMkDir(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doDelete(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doDelete(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doDelete(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doRename(request: openapi_pb.RenameCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doCopy(request: openapi_pb.MoveCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doMove(request: openapi_pb.CopyCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doQueryToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doQueryToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doQueryToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskReadToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskReadToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskReadToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskWriteToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskWriteToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskWriteToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doRefreshToken(request: openapi_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doRefreshToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doRefreshToken(request: openapi_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getReadStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.ResultOfString) => void): grpc.ClientUnaryCall;
}

