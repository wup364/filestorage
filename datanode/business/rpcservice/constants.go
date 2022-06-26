package rpcservice

import "errors"

// ErrorAuthentication 认证失败
var ErrorAuthentication = errors.New("authentication failed")
var ErrorPermissionLess = errors.New("insufficient permissions")
