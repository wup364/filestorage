// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var openapi_pb = require('./openapi_pb.js');
var dto_comm_pb = require('./dto_comm_pb.js');

function serialize_dto_QryOfString(arg) {
  if (!(arg instanceof dto_comm_pb.QryOfString)) {
    throw new Error('Expected argument of type dto.QryOfString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dto_QryOfString(buffer_arg) {
  return dto_comm_pb.QryOfString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_dto_QryOfStrings(arg) {
  if (!(arg instanceof dto_comm_pb.QryOfStrings)) {
    throw new Error('Expected argument of type dto.QryOfStrings');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dto_QryOfStrings(buffer_arg) {
  return dto_comm_pb.QryOfStrings.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_dto_QueryLimitOfString(arg) {
  if (!(arg instanceof dto_comm_pb.QueryLimitOfString)) {
    throw new Error('Expected argument of type dto.QueryLimitOfString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dto_QueryLimitOfString(buffer_arg) {
  return dto_comm_pb.QueryLimitOfString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_dto_ResultOfBool(arg) {
  if (!(arg instanceof dto_comm_pb.ResultOfBool)) {
    throw new Error('Expected argument of type dto.ResultOfBool');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dto_ResultOfBool(buffer_arg) {
  return dto_comm_pb.ResultOfBool.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_dto_ResultOfInt64(arg) {
  if (!(arg instanceof dto_comm_pb.ResultOfInt64)) {
    throw new Error('Expected argument of type dto.ResultOfInt64');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dto_ResultOfInt64(buffer_arg) {
  return dto_comm_pb.ResultOfInt64.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_dto_ResultOfString(arg) {
  if (!(arg instanceof dto_comm_pb.ResultOfString)) {
    throw new Error('Expected argument of type dto.ResultOfString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_dto_ResultOfString(buffer_arg) {
  return dto_comm_pb.ResultOfString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_CopyCmd(arg) {
  if (!(arg instanceof openapi_pb.CopyCmd)) {
    throw new Error('Expected argument of type service.CopyCmd');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_CopyCmd(buffer_arg) {
  return openapi_pb.CopyCmd.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_DirNameListDto(arg) {
  if (!(arg instanceof openapi_pb.DirNameListDto)) {
    throw new Error('Expected argument of type service.DirNameListDto');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_DirNameListDto(buffer_arg) {
  return openapi_pb.DirNameListDto.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_DirNodeListDto(arg) {
  if (!(arg instanceof openapi_pb.DirNodeListDto)) {
    throw new Error('Expected argument of type service.DirNodeListDto');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_DirNodeListDto(buffer_arg) {
  return openapi_pb.DirNodeListDto.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_MoveCmd(arg) {
  if (!(arg instanceof openapi_pb.MoveCmd)) {
    throw new Error('Expected argument of type service.MoveCmd');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_MoveCmd(buffer_arg) {
  return openapi_pb.MoveCmd.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_QryStreamURLCmd(arg) {
  if (!(arg instanceof openapi_pb.QryStreamURLCmd)) {
    throw new Error('Expected argument of type service.QryStreamURLCmd');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_QryStreamURLCmd(buffer_arg) {
  return openapi_pb.QryStreamURLCmd.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_RenameCmd(arg) {
  if (!(arg instanceof openapi_pb.RenameCmd)) {
    throw new Error('Expected argument of type service.RenameCmd');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_RenameCmd(buffer_arg) {
  return openapi_pb.RenameCmd.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_StreamToken(arg) {
  if (!(arg instanceof openapi_pb.StreamToken)) {
    throw new Error('Expected argument of type service.StreamToken');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_StreamToken(buffer_arg) {
  return openapi_pb.StreamToken.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_SubmitTokenCmd(arg) {
  if (!(arg instanceof openapi_pb.SubmitTokenCmd)) {
    throw new Error('Expected argument of type service.SubmitTokenCmd');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_SubmitTokenCmd(buffer_arg) {
  return openapi_pb.SubmitTokenCmd.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_TNode(arg) {
  if (!(arg instanceof openapi_pb.TNode)) {
    throw new Error('Expected argument of type service.TNode');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_TNode(buffer_arg) {
  return openapi_pb.TNode.deserializeBinary(new Uint8Array(buffer_arg));
}


// IOpenApi 开放api
var OpenApiService = exports.OpenApiService = {
  isDir: {
    path: '/service.OpenApi/IsDir',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: dto_comm_pb.ResultOfBool,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_dto_ResultOfBool,
    responseDeserialize: deserialize_dto_ResultOfBool,
  },
  isFile: {
    path: '/service.OpenApi/IsFile',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: dto_comm_pb.ResultOfBool,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_dto_ResultOfBool,
    responseDeserialize: deserialize_dto_ResultOfBool,
  },
  isExist: {
    path: '/service.OpenApi/IsExist',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: dto_comm_pb.ResultOfBool,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_dto_ResultOfBool,
    responseDeserialize: deserialize_dto_ResultOfBool,
  },
  getFileSize: {
    path: '/service.OpenApi/GetFileSize',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: dto_comm_pb.ResultOfInt64,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_dto_ResultOfInt64,
    responseDeserialize: deserialize_dto_ResultOfInt64,
  },
  getNode: {
    path: '/service.OpenApi/GetNode',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: openapi_pb.TNode,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_service_TNode,
    responseDeserialize: deserialize_service_TNode,
  },
  getNodes: {
    path: '/service.OpenApi/GetNodes',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfStrings,
    responseType: openapi_pb.DirNodeListDto,
    requestSerialize: serialize_dto_QryOfStrings,
    requestDeserialize: deserialize_dto_QryOfStrings,
    responseSerialize: serialize_service_DirNodeListDto,
    responseDeserialize: deserialize_service_DirNodeListDto,
  },
  getDirNameList: {
    path: '/service.OpenApi/GetDirNameList',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QueryLimitOfString,
    responseType: openapi_pb.DirNameListDto,
    requestSerialize: serialize_dto_QueryLimitOfString,
    requestDeserialize: deserialize_dto_QueryLimitOfString,
    responseSerialize: serialize_service_DirNameListDto,
    responseDeserialize: deserialize_service_DirNameListDto,
  },
  getDirNodeList: {
    path: '/service.OpenApi/GetDirNodeList',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QueryLimitOfString,
    responseType: openapi_pb.DirNodeListDto,
    requestSerialize: serialize_dto_QueryLimitOfString,
    requestDeserialize: deserialize_dto_QueryLimitOfString,
    responseSerialize: serialize_service_DirNodeListDto,
    responseDeserialize: deserialize_service_DirNodeListDto,
  },
  doMkDir: {
    path: '/service.OpenApi/DoMkDir',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: dto_comm_pb.ResultOfString,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_dto_ResultOfString,
    responseDeserialize: deserialize_dto_ResultOfString,
  },
  doDelete: {
    path: '/service.OpenApi/DoDelete',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: dto_comm_pb.ResultOfBool,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_dto_ResultOfBool,
    responseDeserialize: deserialize_dto_ResultOfBool,
  },
  doRename: {
    path: '/service.OpenApi/DoRename',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.RenameCmd,
    responseType: dto_comm_pb.ResultOfBool,
    requestSerialize: serialize_service_RenameCmd,
    requestDeserialize: deserialize_service_RenameCmd,
    responseSerialize: serialize_dto_ResultOfBool,
    responseDeserialize: deserialize_dto_ResultOfBool,
  },
  doCopy: {
    path: '/service.OpenApi/DoCopy',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.MoveCmd,
    responseType: dto_comm_pb.ResultOfString,
    requestSerialize: serialize_service_MoveCmd,
    requestDeserialize: deserialize_service_MoveCmd,
    responseSerialize: serialize_dto_ResultOfString,
    responseDeserialize: deserialize_dto_ResultOfString,
  },
  doMove: {
    path: '/service.OpenApi/DoMove',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.CopyCmd,
    responseType: dto_comm_pb.ResultOfBool,
    requestSerialize: serialize_service_CopyCmd,
    requestDeserialize: deserialize_service_CopyCmd,
    responseSerialize: serialize_dto_ResultOfBool,
    responseDeserialize: deserialize_dto_ResultOfBool,
  },
  doQueryToken: {
    path: '/service.OpenApi/DoQueryToken',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_service_StreamToken,
    responseDeserialize: deserialize_service_StreamToken,
  },
  doAskReadToken: {
    path: '/service.OpenApi/DoAskReadToken',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_service_StreamToken,
    responseDeserialize: deserialize_service_StreamToken,
  },
  doAskWriteToken: {
    path: '/service.OpenApi/DoAskWriteToken',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_service_StreamToken,
    responseDeserialize: deserialize_service_StreamToken,
  },
  doRefreshToken: {
    path: '/service.OpenApi/DoRefreshToken',
    requestStream: false,
    responseStream: false,
    requestType: dto_comm_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_dto_QryOfString,
    requestDeserialize: deserialize_dto_QryOfString,
    responseSerialize: serialize_service_StreamToken,
    responseDeserialize: deserialize_service_StreamToken,
  },
  doSubmitWriteToken: {
    path: '/service.OpenApi/DoSubmitWriteToken',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.SubmitTokenCmd,
    responseType: openapi_pb.TNode,
    requestSerialize: serialize_service_SubmitTokenCmd,
    requestDeserialize: deserialize_service_SubmitTokenCmd,
    responseSerialize: serialize_service_TNode,
    responseDeserialize: deserialize_service_TNode,
  },
  getReadStreamURL: {
    path: '/service.OpenApi/GetReadStreamURL',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryStreamURLCmd,
    responseType: dto_comm_pb.ResultOfString,
    requestSerialize: serialize_service_QryStreamURLCmd,
    requestDeserialize: deserialize_service_QryStreamURLCmd,
    responseSerialize: serialize_dto_ResultOfString,
    responseDeserialize: deserialize_dto_ResultOfString,
  },
  getWriteStreamURL: {
    path: '/service.OpenApi/GetWriteStreamURL',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryStreamURLCmd,
    responseType: dto_comm_pb.ResultOfString,
    requestSerialize: serialize_service_QryStreamURLCmd,
    requestDeserialize: deserialize_service_QryStreamURLCmd,
    responseSerialize: serialize_dto_ResultOfString,
    responseDeserialize: deserialize_dto_ResultOfString,
  },
};

exports.OpenApiClient = grpc.makeGenericClientConstructor(OpenApiService);
