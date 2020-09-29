gen_proto:
	protoc --proto_path=../api/protos --go_out=../api \
		   --go-grpc_out=../api --go-grpc_opt=paths=source_relative \
		   --go_opt=paths=source_relative ../api/protos/plans/plans.proto
