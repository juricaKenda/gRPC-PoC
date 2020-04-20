const grpc = require("grpc");
const loader = require("@grpc/proto-loader");
const pkg = loader.loadSync("../../proto/topics.proto", {});
const grpcObject = grpc.loadPackageDefinition(pkg);
const topicPkg = grpcObject.topics;

const client = new topicPkg.TopicService("localhost:50000", grpc.credentials.createInsecure());

let topic = subscription_message("time");
const call = client.Pull(topic);
listen();

function subscription_message(topic) {
    return {
        body: topic,
    };
}

function listen() {
    call.on(
        "data", message => {
            console.log("pero received: " + message.body)
        }
    )
}
