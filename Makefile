proto:
	protoc -I./proto --go_out=plugins=grpc:./distance_service/trippb ./proto/*.proto && \
	grpc_tools_ruby_protoc -I ./proto --ruby_out=./price_calculation_service/proto --grpc_out=./price_calculation_service/proto ./proto/*.proto