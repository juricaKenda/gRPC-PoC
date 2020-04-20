gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

broker:
	go build broker/main.go
	go run broker/main.go

client_go:
	go build clients/golang/main.go
	go run clients/golang/main.go