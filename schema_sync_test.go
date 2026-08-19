package shared_protos

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestProtobufReflectionSync_Identity(t *testing.T) {
	idMsg := (&InfrastructureIdentity{}).ProtoReflect()
	descriptor := idMsg.Descriptor()

	expectedFields := map[protoreflect.FieldNumber]string{
		1: "tenant_id",
		2: "app_id",
		3: "job_id",
	}

	if descriptor.Fields().Len() != len(expectedFields) {
		t.Fatalf("expected %d fields in InfrastructureIdentity, found %d", len(expectedFields), descriptor.Fields().Len())
	}

	for num, name := range expectedFields {
		fd := descriptor.Fields().ByNumber(num)
		if fd == nil {
			t.Fatalf("missing field number %d in InfrastructureIdentity", num)
		}
		if string(fd.Name()) != name {
			t.Fatalf("expected field %d to be %q, got %q", num, name, fd.Name())
		}
		if fd.Kind() != protoreflect.StringKind {
			t.Fatalf("expected field %d to be string, got %v", num, fd.Kind())
		}
	}
}

func TestProtobufReflectionSync_ValidationContracts(t *testing.T) {
	reqMsg := (&FieldValidationRequest{}).ProtoReflect()
	descReq := reqMsg.Descriptor()

	if descReq.Fields().ByNumber(1) == nil || string(descReq.Fields().ByNumber(1).Name()) != "identity" {
		t.Fatalf("expected field 1 of FieldValidationRequest to be identity")
	}

	reportMsg := (&DocumentValidationReport{}).ProtoReflect()
	descReport := reportMsg.Descriptor()
	if descReport.Fields().ByNumber(3) == nil || string(descReport.Fields().ByNumber(3).Name()) != "field_results" {
		t.Fatalf("expected field 3 of DocumentValidationReport to be field_results")
	}
}

func TestProtobufReflectionSync_MLWorker(t *testing.T) {
	redactReq := (&RedactDocumentRequest{}).ProtoReflect()
	descRedact := redactReq.Descriptor()

	if descRedact.Fields().ByNumber(7) == nil || string(descRedact.Fields().ByNumber(7).Name()) != "dynamic_labels" {
		t.Fatalf("expected field 7 of RedactDocumentRequest to be dynamic_labels")
	}
	if descRedact.Fields().ByNumber(8) == nil || string(descRedact.Fields().ByNumber(8).Name()) != "custom_boxes" {
		t.Fatalf("expected field 8 of RedactDocumentRequest to be custom_boxes")
	}
}

func TestProtobufReflectionSync_DocumentProcessing(t *testing.T) {
	pageEval := (&PageEvaluation{}).ProtoReflect()
	desc := pageEval.Descriptor()

	if desc.Fields().ByNumber(1) == nil || string(desc.Fields().ByNumber(1).Name()) != "page_number" {
		t.Fatalf("expected field 1 of PageEvaluation to be page_number")
	}
	if desc.Fields().ByNumber(5) == nil || string(desc.Fields().ByNumber(5).Name()) != "boundary_confidence" {
		t.Fatalf("expected field 5 of PageEvaluation to be boundary_confidence")
	}
}
