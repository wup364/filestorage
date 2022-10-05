// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var openapi_pb = require('./openapi_pb.js');

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

function serialize_service_QryOfString(arg) {
  if (!(arg instanceof openapi_pb.QryOfString)) {
    throw new Error('Expected argument of type service.QryOfString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_QryOfString(buffer_arg) {
  return openapi_pb.QryOfString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_QryOfStrings(arg) {
  if (!(arg instanceof openapi_pb.QryOfStrings)) {
    throw new Error('Expected argument of type service.QryOfStrings');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_QryOfStrings(buffer_arg) {
  return openapi_pb.QryOfStrings.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_service_QueryLimitOfString(arg) {
  if (!(arg instanceof openapi_pb.QueryLimitOfString)) {
    throw new Error('Expected argument of type service.QueryLimitOfString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_QueryLimitOfString(buffer_arg) {
  return openapi_pb.QueryLimitOfString.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_service_ResultOfBool(arg) {
  if (!(arg instanceof openapi_pb.ResultOfBool)) {
    throw new Error('Expected argument of type service.ResultOfBool');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_ResultOfBool(buffer_arg) {
  return openapi_pb.ResultOfBool.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_ResultOfInt64(arg) {
  if (!(arg instanceof openapi_pb.ResultOfInt64)) {
    throw new Error('Expected argument of type service.ResultOfInt64');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_ResultOfInt64(buffer_arg) {
  return openapi_pb.ResultOfInt64.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_service_ResultOfString(arg) {
  if (!(arg instanceof openapi_pb.ResultOfString)) {
    throw new Error('Expected argument of type service.ResultOfString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_service_ResultOfString(buffer_arg) {
  return openapi_pb.ResultOfString.deserializeBinary(new Uint8Array(buffer_arg));
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
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.ResultOfBool,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_ResultOfBool,
    responseDeserialize: deserialize_service_ResultOfBool,
  },
  isFile: {
    path: '/service.OpenApi/IsFile',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.ResultOfBool,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_ResultOfBool,
    responseDeserialize: deserialize_service_ResultOfBool,
  },
  isExist: {
    path: '/service.OpenApi/IsExist',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.ResultOfBool,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_ResultOfBool,
    responseDeserialize: deserialize_service_ResultOfBool,
  },
  getFileSize: {
    path: '/service.OpenApi/GetFileSize',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.ResultOfInt64,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_ResultOfInt64,
    responseDeserialize: deserialize_service_ResultOfInt64,
  },
  getNode: {
    path: '/service.OpenApi/GetNode',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.TNode,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_TNode,
    responseDeserialize: deserialize_service_TNode,
  },
  getNodes: {
    path: '/service.OpenApi/GetNodes',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfStrings,
    responseType: openapi_pb.DirNodeListDto,
    requestSerialize: serialize_service_QryOfStrings,
    requestDeserialize: deserialize_service_QryOfStrings,
    responseSerialize: serialize_service_DirNodeListDto,
    responseDeserialize: deserialize_service_DirNodeListDto,
  },
  getDirNameList: {
    path: '/service.OpenApi/GetDirNameList',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QueryLimitOfString,
    responseType: openapi_pb.DirNameListDto,
    requestSerialize: serialize_service_QueryLimitOfString,
    requestDeserialize: deserialize_service_QueryLimitOfString,
    responseSerialize: serialize_service_DirNameListDto,
    responseDeserialize: deserialize_service_DirNameListDto,
  },
  getDirNodeList: {
    path: '/service.OpenApi/GetDirNodeList',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QueryLimitOfString,
    responseType: openapi_pb.DirNodeListDto,
    requestSerialize: serialize_service_QueryLimitOfString,
    requestDeserialize: deserialize_service_QueryLimitOfString,
    responseSerialize: serialize_service_DirNodeListDto,
    responseDeserialize: deserialize_service_DirNodeListDto,
  },
  doMkDir: {
    path: '/service.OpenApi/DoMkDir',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.ResultOfString,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_ResultOfString,
    responseDeserialize: deserialize_service_ResultOfString,
  },
  doDelete: {
    path: '/service.OpenApi/DoDelete',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.ResultOfBool,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_ResultOfBool,
    responseDeserialize: deserialize_service_ResultOfBool,
  },
  doRename: {
    path: '/service.OpenApi/DoRename',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.RenameCmd,
    responseType: openapi_pb.ResultOfBool,
    requestSerialize: serialize_service_RenameCmd,
    requestDeserialize: deserialize_service_RenameCmd,
    responseSerialize: serialize_service_ResultOfBool,
    responseDeserialize: deserialize_service_ResultOfBool,
  },
  doCopy: {
    path: '/service.OpenApi/DoCopy',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.MoveCmd,
    responseType: openapi_pb.ResultOfString,
    requestSerialize: serialize_service_MoveCmd,
    requestDeserialize: deserialize_service_MoveCmd,
    responseSerialize: serialize_service_ResultOfString,
    responseDeserialize: deserialize_service_ResultOfString,
  },
  doMove: {
    path: '/service.OpenApi/DoMove',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.CopyCmd,
    responseType: openapi_pb.ResultOfBool,
    requestSerialize: serialize_service_CopyCmd,
    requestDeserialize: deserialize_service_CopyCmd,
    responseSerialize: serialize_service_ResultOfBool,
    responseDeserialize: deserialize_service_ResultOfBool,
  },
  doQueryToken: {
    path: '/service.OpenApi/DoQueryToken',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_StreamToken,
    responseDeserialize: deserialize_service_StreamToken,
  },
  doAskReadToken: {
    path: '/service.OpenApi/DoAskReadToken',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_StreamToken,
    responseDeserialize: deserialize_service_StreamToken,
  },
  doAskWriteToken: {
    path: '/service.OpenApi/DoAskWriteToken',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
    responseSerialize: serialize_service_StreamToken,
    responseDeserialize: deserialize_service_StreamToken,
  },
  doRefreshToken: {
    path: '/service.OpenApi/DoRefreshToken',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryOfString,
    responseType: openapi_pb.StreamToken,
    requestSerialize: serialize_service_QryOfString,
    requestDeserialize: deserialize_service_QryOfString,
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
    responseType: openapi_pb.ResultOfString,
    requestSerialize: serialize_service_QryStreamURLCmd,
    requestDeserialize: deserialize_service_QryStreamURLCmd,
    responseSerialize: serialize_service_ResultOfString,
    responseDeserialize: deserialize_service_ResultOfString,
  },
  getWriteStreamURL: {
    path: '/service.OpenApi/GetWriteStreamURL',
    requestStream: false,
    responseStream: false,
    requestType: openapi_pb.QryStreamURLCmd,
    responseType: openapi_pb.ResultOfString,
    requestSerialize: serialize_service_QryStreamURLCmd,
    requestDeserialize: deserialize_service_QryStreamURLCmd,
    responseSerialize: serialize_service_ResultOfString,
    responseDeserialize: deserialize_service_ResultOfString,
  },
};

exports.OpenApiClient = grpc.makeGenericClientConstructor(OpenApiService);
