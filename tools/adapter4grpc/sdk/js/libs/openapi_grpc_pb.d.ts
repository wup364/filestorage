// package: service
// file: openapi.proto

import * as grpc from '@grpc/grpc-js';
import * as openapi_pb from './openapi_pb';
import * as dto_comm_pb from './dto_comm_pb';

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

interface IOpenApiService_IIsDir extends grpc.MethodDefinition<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool> {
  path: '/service.OpenApi/IsDir'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfBool>;
}

interface IOpenApiService_IIsFile extends grpc.MethodDefinition<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool> {
  path: '/service.OpenApi/IsFile'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfBool>;
}

interface IOpenApiService_IIsExist extends grpc.MethodDefinition<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool> {
  path: '/service.OpenApi/IsExist'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfBool>;
}

interface IOpenApiService_IGetFileSize extends grpc.MethodDefinition<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfInt64> {
  path: '/service.OpenApi/GetFileSize'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfInt64>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfInt64>;
}

interface IOpenApiService_IGetNode extends grpc.MethodDefinition<dto_comm_pb.QryOfString, openapi_pb.TNode> {
  path: '/service.OpenApi/GetNode'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.TNode>;
  responseDeserialize: grpc.deserialize<openapi_pb.TNode>;
}

interface IOpenApiService_IGetNodes extends grpc.MethodDefinition<dto_comm_pb.QryOfStrings, openapi_pb.DirNodeListDto> {
  path: '/service.OpenApi/GetNodes'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfStrings>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfStrings>;
  responseSerialize: grpc.serialize<openapi_pb.DirNodeListDto>;
  responseDeserialize: grpc.deserialize<openapi_pb.DirNodeListDto>;
}

interface IOpenApiService_IGetDirNameList extends grpc.MethodDefinition<dto_comm_pb.QueryLimitOfString, openapi_pb.DirNameListDto> {
  path: '/service.OpenApi/GetDirNameList'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QueryLimitOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QueryLimitOfString>;
  responseSerialize: grpc.serialize<openapi_pb.DirNameListDto>;
  responseDeserialize: grpc.deserialize<openapi_pb.DirNameListDto>;
}

interface IOpenApiService_IGetDirNodeList extends grpc.MethodDefinition<dto_comm_pb.QueryLimitOfString, openapi_pb.DirNodeListDto> {
  path: '/service.OpenApi/GetDirNodeList'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QueryLimitOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QueryLimitOfString>;
  responseSerialize: grpc.serialize<openapi_pb.DirNodeListDto>;
  responseDeserialize: grpc.deserialize<openapi_pb.DirNodeListDto>;
}

interface IOpenApiService_IDoMkDir extends grpc.MethodDefinition<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfString> {
  path: '/service.OpenApi/DoMkDir'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfString>;
}

interface IOpenApiService_IDoDelete extends grpc.MethodDefinition<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool> {
  path: '/service.OpenApi/DoDelete'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfBool>;
}

interface IOpenApiService_IDoRename extends grpc.MethodDefinition<openapi_pb.RenameCmd, dto_comm_pb.ResultOfBool> {
  path: '/service.OpenApi/DoRename'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.RenameCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.RenameCmd>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfBool>;
}

interface IOpenApiService_IDoCopy extends grpc.MethodDefinition<openapi_pb.MoveCmd, dto_comm_pb.ResultOfString> {
  path: '/service.OpenApi/DoCopy'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.MoveCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.MoveCmd>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfString>;
}

interface IOpenApiService_IDoMove extends grpc.MethodDefinition<openapi_pb.CopyCmd, dto_comm_pb.ResultOfBool> {
  path: '/service.OpenApi/DoMove'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.CopyCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.CopyCmd>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfBool>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfBool>;
}

interface IOpenApiService_IDoQueryToken extends grpc.MethodDefinition<dto_comm_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoQueryToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.StreamToken>;
  responseDeserialize: grpc.deserialize<openapi_pb.StreamToken>;
}

interface IOpenApiService_IDoAskReadToken extends grpc.MethodDefinition<dto_comm_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoAskReadToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.StreamToken>;
  responseDeserialize: grpc.deserialize<openapi_pb.StreamToken>;
}

interface IOpenApiService_IDoAskWriteToken extends grpc.MethodDefinition<dto_comm_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoAskWriteToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
  responseSerialize: grpc.serialize<openapi_pb.StreamToken>;
  responseDeserialize: grpc.deserialize<openapi_pb.StreamToken>;
}

interface IOpenApiService_IDoRefreshToken extends grpc.MethodDefinition<dto_comm_pb.QryOfString, openapi_pb.StreamToken> {
  path: '/service.OpenApi/DoRefreshToken'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<dto_comm_pb.QryOfString>;
  requestDeserialize: grpc.deserialize<dto_comm_pb.QryOfString>;
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

interface IOpenApiService_IGetReadStreamURL extends grpc.MethodDefinition<openapi_pb.QryStreamURLCmd, dto_comm_pb.ResultOfString> {
  path: '/service.OpenApi/GetReadStreamURL'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryStreamURLCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryStreamURLCmd>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfString>;
}

interface IOpenApiService_IGetWriteStreamURL extends grpc.MethodDefinition<openapi_pb.QryStreamURLCmd, dto_comm_pb.ResultOfString> {
  path: '/service.OpenApi/GetWriteStreamURL'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<openapi_pb.QryStreamURLCmd>;
  requestDeserialize: grpc.deserialize<openapi_pb.QryStreamURLCmd>;
  responseSerialize: grpc.serialize<dto_comm_pb.ResultOfString>;
  responseDeserialize: grpc.deserialize<dto_comm_pb.ResultOfString>;
}

export const OpenApiService: IOpenApiService;
export interface IOpenApiServer extends grpc.UntypedServiceImplementation {
  isDir: grpc.handleUnaryCall<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool>;
  isFile: grpc.handleUnaryCall<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool>;
  isExist: grpc.handleUnaryCall<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool>;
  getFileSize: grpc.handleUnaryCall<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfInt64>;
  getNode: grpc.handleUnaryCall<dto_comm_pb.QryOfString, openapi_pb.TNode>;
  getNodes: grpc.handleUnaryCall<dto_comm_pb.QryOfStrings, openapi_pb.DirNodeListDto>;
  getDirNameList: grpc.handleUnaryCall<dto_comm_pb.QueryLimitOfString, openapi_pb.DirNameListDto>;
  getDirNodeList: grpc.handleUnaryCall<dto_comm_pb.QueryLimitOfString, openapi_pb.DirNodeListDto>;
  doMkDir: grpc.handleUnaryCall<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfString>;
  doDelete: grpc.handleUnaryCall<dto_comm_pb.QryOfString, dto_comm_pb.ResultOfBool>;
  doRename: grpc.handleUnaryCall<openapi_pb.RenameCmd, dto_comm_pb.ResultOfBool>;
  doCopy: grpc.handleUnaryCall<openapi_pb.MoveCmd, dto_comm_pb.ResultOfString>;
  doMove: grpc.handleUnaryCall<openapi_pb.CopyCmd, dto_comm_pb.ResultOfBool>;
  doQueryToken: grpc.handleUnaryCall<dto_comm_pb.QryOfString, openapi_pb.StreamToken>;
  doAskReadToken: grpc.handleUnaryCall<dto_comm_pb.QryOfString, openapi_pb.StreamToken>;
  doAskWriteToken: grpc.handleUnaryCall<dto_comm_pb.QryOfString, openapi_pb.StreamToken>;
  doRefreshToken: grpc.handleUnaryCall<dto_comm_pb.QryOfString, openapi_pb.StreamToken>;
  doSubmitWriteToken: grpc.handleUnaryCall<openapi_pb.SubmitTokenCmd, openapi_pb.TNode>;
  getReadStreamURL: grpc.handleUnaryCall<openapi_pb.QryStreamURLCmd, dto_comm_pb.ResultOfString>;
  getWriteStreamURL: grpc.handleUnaryCall<openapi_pb.QryStreamURLCmd, dto_comm_pb.ResultOfString>;
}

export interface IOpenApiClient {
  isDir(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isFile(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isFile(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isFile(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isExist(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isExist(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  isExist(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  getFileSize(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  getFileSize(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  getFileSize(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  getNode(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getNode(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getNode(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getNodes(request: dto_comm_pb.QryOfStrings, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getNodes(request: dto_comm_pb.QryOfStrings, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getNodes(request: dto_comm_pb.QryOfStrings, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getDirNameList(request: dto_comm_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  getDirNameList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  getDirNameList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  getDirNodeList(request: dto_comm_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getDirNodeList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  getDirNodeList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  doMkDir(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doMkDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doMkDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doDelete(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doDelete(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doDelete(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doRename(request: openapi_pb.RenameCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doCopy(request: openapi_pb.MoveCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  doMove(request: openapi_pb.CopyCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  doQueryToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doQueryToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doQueryToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskReadToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskReadToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskReadToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskWriteToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskWriteToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doAskWriteToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doRefreshToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doRefreshToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doRefreshToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  getReadStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
}

export class OpenApiClient extends grpc.Client implements IOpenApiClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
  public isDir(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isFile(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isFile(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isFile(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isExist(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isExist(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public isExist(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public getFileSize(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  public getFileSize(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  public getFileSize(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfInt64) => void): grpc.ClientUnaryCall;
  public getNode(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getNode(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getNode(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getNodes(request: dto_comm_pb.QryOfStrings, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getNodes(request: dto_comm_pb.QryOfStrings, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getNodes(request: dto_comm_pb.QryOfStrings, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getDirNameList(request: dto_comm_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  public getDirNameList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  public getDirNameList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNameListDto) => void): grpc.ClientUnaryCall;
  public getDirNodeList(request: dto_comm_pb.QueryLimitOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getDirNodeList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public getDirNodeList(request: dto_comm_pb.QueryLimitOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.DirNodeListDto) => void): grpc.ClientUnaryCall;
  public doMkDir(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doMkDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doMkDir(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doDelete(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doDelete(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doDelete(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doRename(request: openapi_pb.RenameCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doRename(request: openapi_pb.RenameCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doCopy(request: openapi_pb.MoveCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doCopy(request: openapi_pb.MoveCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public doMove(request: openapi_pb.CopyCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doMove(request: openapi_pb.CopyCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfBool) => void): grpc.ClientUnaryCall;
  public doQueryToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doQueryToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doQueryToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskReadToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskReadToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskReadToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskWriteToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskWriteToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doAskWriteToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doRefreshToken(request: dto_comm_pb.QryOfString, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doRefreshToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doRefreshToken(request: dto_comm_pb.QryOfString, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.StreamToken) => void): grpc.ClientUnaryCall;
  public doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public doSubmitWriteToken(request: openapi_pb.SubmitTokenCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: openapi_pb.TNode) => void): grpc.ClientUnaryCall;
  public getReadStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getReadStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
  public getWriteStreamURL(request: openapi_pb.QryStreamURLCmd, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: dto_comm_pb.ResultOfString) => void): grpc.ClientUnaryCall;
}

