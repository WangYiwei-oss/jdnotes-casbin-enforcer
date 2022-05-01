protoc --proto_path=src/pbfiles --plugin=protoc-gen-go=C:\Users\WangYiWei\go\bin\protoc-gen-go.exe --go_out=../ enforce_models.proto

protoc --proto_path=src/pbfiles --go_grpc_out=../ --plugin=protoc-gen-go=C:\Users\WangYiWei\go\bin\protoc-gen-go.exe --plugin=protoc-gen-go_grpc=C:\Users\WangYiWei\go\bin\protoc-gen-go-grpc.exe enforce_services.proto