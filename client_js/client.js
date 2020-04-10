const grpc = require("grpc");
const loader = require("@grpc/proto-loader");
const pkg = loader.loadSync("proto/grpc_ping_pong.proto", {});
const grpcObject = grpc.loadPackageDefinition(pkg);
const pingPongPackage = grpcObject.pingPongPackage;


const client = new pingPongPackage.PingPongService("localhost:50000", grpc.credentials.createInsecure());

const call = client.Chat();

function askForTime(n) {
    let ping = buildPingMessage();
    call.write(ping);

    if (n > 0) {
        setTimeout(() => askForTime(n - 1), 1000);
    }
}

function buildPingMessage() {
    console.log("JS client_golang sending ping message..")
    return {
        message: "What is the time?"
    };
}

function listenForResponse() {
    call.on("data", pong => {
        console.log("JS client_golang received a response: " + pong.message)
    });
}

askForTime(10);

listenForResponse();

