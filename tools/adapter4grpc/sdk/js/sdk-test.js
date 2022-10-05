var grpc = require('@grpc/grpc-js');

var dtoOfComm = require('./libs/dto_comm_pb');
var dtoOfOpenAPI = require('./libs/openapi_pb');
var services = require('./libs/openapi_grpc_pb');


function main() {
    var target = '127.0.0.1:5059';
    var client = new services.OpenApiClient(target, grpc.credentials.createInsecure());
    var request = new dtoOfComm.QryOfString();
    request.setQuery("/")
    client.isDir(request, function (err, response) {
        if (!err) {
            console.log('Greeting:', response.getMessage());
        } else {
            console.error(err);
        }
    });
}

main();