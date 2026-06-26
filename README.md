# Shared Protos

This repository contains the shared Protocol Buffer definitions for the Document Processor microservices.

## Generating Go Code

To generate the Go code from these `.proto` definitions, make sure you have `protoc` and the Go plugins installed:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Then run the generation command inside this directory:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    *.proto
```

## Services

* **Orchestrator**: Defines the workflow interface and streaming callbacks for status updates.
* **Processor**: Defines the AI logic methods for both async document processing and fast synchronous chunk extraction.
