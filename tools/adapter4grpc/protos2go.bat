call go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
call go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
@REM call protoc -I ./protos/ ./protos/*.proto --go_out=plugins=grpc:.
call protoc --go_out=. --go_opt=paths=import --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=import ./protos/*.proto