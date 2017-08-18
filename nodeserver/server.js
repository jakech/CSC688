var path = require('path')
var PROTO_PATH = path.resolve(__dirname, '../proto/rpc.proto')
var config = require('./config')

var grpc = require('grpc')
var rpc_proto = grpc.load(PROTO_PATH).rpc

var category = require('./category')

function main() {
    var server = new grpc.Server()
    server.addService(rpc_proto.Rpc.service, category)
    server.bind(
        config.host + ':' + config.port,
        grpc.ServerCredentials.createInsecure()
    )
    server.start()
}

main()
