gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

.PHONY: broker
broker:
	go build broker/main.go
	go run broker/main.go

.PHONY: gocli
gocli:
	go build clients/golang/main.go
	go run clients/golang/main.go

.PHONY: jscli
jscli:
	node clients/js/client.js