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

* **Orchestrator**: Defines the unified workflow interface. The Orchestrator operates as a pure execution engine using a single method `ExecuteWorkflowNode(ExecuteWorkflowNodeRequest)`. Clients pass a specific `NodeConfiguration` along with the current `GlobalState`, which the Orchestrator executes according to the chosen `OrchestratorPrimitive` (e.g. `PRIMITIVE_AI_STEP`, `PRIMITIVE_HUMAN_STEP`), returning state mutations back to the client. **Rule: This interface is not to be altered unless explicitly told to do so.**
* **Processor**: Defines the AI logic methods for both async document processing and fast synchronous chunk extraction.
