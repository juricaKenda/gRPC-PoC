const grpc = require("grpc");
const loader = require("@grpc/proto-loader");
const pkg = loader.loadSync("../proto/grpc_ping_pong.proto", {});
const grpcObject = grpc.loadPackageDefinition(pkg);
const pingPongPackage = grpcObject.pingPongPackage;


const client = new pingPongPackage.PingPongService("localhost:50000", grpc.credentials.createInsecure());

const call = client.Chat();


function notifyAboutTheListening() {
    console.log("Node client sending listening message..")
    return {
        message: "node_client"
    };
}

function listenForResponse() {
    call.on("data", pong => {
        console.log("JS client_golang received a response: " + pong.message)
    });
}


let ping = notifyAboutTheListening();
call.write(ping);

listenForResponse();

