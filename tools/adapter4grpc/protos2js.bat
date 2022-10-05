SET OUT_DIR=./sdk/js/libs

call npm install request -g
call npm config set unsafe-perm true
call npm install protoc-gen-grpc -g
@REM  build
call protoc-gen-grpc --js_out=import_style=commonjs,binary:%OUT_DIR% --grpc_out=grpc_js:%OUT_DIR% --proto_path ./protos ./protos/*.proto
call protoc-gen-grpc-ts --ts_out=grpc_js:%OUT_DIR% --proto_path ./protos ./protos/*.proto