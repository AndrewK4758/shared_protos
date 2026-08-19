# Core Components

This repository contains the shared Protocol Buffer definitions for the MPloyee microservices.

## Services

* **Orchestrator**: Defines the unified workflow interface. The Orchestrator operates as a pure execution engine using a single method `ExecuteWorkflowNode(ExecuteWorkflowNodeRequest)`. Clients pass a specific `NodeConfiguration` along with the current `GlobalState`, which the Orchestrator executes according to the chosen `OrchestratorPrimitive` (e.g. `PRIMITIVE_AI_STEP`, `PRIMITIVE_HUMAN_STEP`), returning state mutations back to the client. **Rule: This interface is not to be altered unless explicitly told to do so.**
* **Processor**: Defines the AI logic methods for both async document processing and fast synchronous chunk extraction.
