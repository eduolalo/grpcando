var PROTO_FILE = __dirname + '/../chat.proto';
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_FILE, {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
);
var chat = grpc.loadPackageDefinition(packageDefinition).chat;
var client = new chat.ChatService('127.0.0.1:9000', grpc.credentials.createInsecure());

function holaPerro() {


    client.holaPrro({
        body: 'Ahí te va la de hacer hijos xD'
    }, function(error, mssg) {

        if (error) {

            console.log('Error: ', error)
            return
        }
        console.log('mensaje: ', mssg.body)
    })
}

holaPerro()