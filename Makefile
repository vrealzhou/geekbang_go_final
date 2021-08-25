build:
	go build -o $(GOPATH)/bin/dev_env cmd/dev_env/*.go

build_container:
	go build -o $(GOPATH)/bin/effected cmd/effected/*.go

gen_grpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/*.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/servestale/api/*.proto