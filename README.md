# gRPC Topic Broker

## Try it yourself
    cd topicbroker/ 
    make broker
    make gocli 
    node client.js

## Expected result
- delivered data is the same on each client subscribed to a topic (one client does not exclusively consume the topic update)
- single client can be subscribed to multiple topics
