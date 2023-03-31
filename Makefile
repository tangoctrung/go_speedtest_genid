gen_proto:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
run_server:
	go run main.go
run_client:
	go run client/client.go