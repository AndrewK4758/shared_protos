package shared_protos

import (
	"testing"

	"google.golang.org/protobuf/types/known/structpb"
)

func TestStateStruct_UnconditionalHydrationAndNativeUnwrap(t *testing.T) {
	initialPayload := map[string]any{
		"EmailBody":            "Here is the invoice for property closing at 100 Main St.",
		"rawHtmlBody":          "<html><body>Here is the invoice...</body></html>",
		"storage.document_uri": "internal-store://tenant-1/app-alpha/job-1001/closing.pdf",
		"nested_metadata": map[string]any{
			"sender": "closing@titlecompany.com",
			"pages":  73,
		},
	}

	stateStruct, err := structpb.NewStruct(initialPayload)
	if err != nil {
		t.Fatalf("failed to construct protobuf struct: %v", err)
	}

	fields := stateStruct.GetFields()
	if fields["EmailBody"].GetStringValue() != "Here is the invoice for property closing at 100 Main St." {
		t.Fatalf("unexpected EmailBody value: %v", fields["EmailBody"])
	}
	if fields["storage.document_uri"].GetStringValue() != "internal-store://tenant-1/app-alpha/job-1001/closing.pdf" {
		t.Fatalf("unexpected document URI: %v", fields["storage.document_uri"])
	}

	nested := fields["nested_metadata"].GetStructValue().GetFields()
	if nested["pages"].GetNumberValue() != 73 {
		t.Fatalf("unexpected nested pages value: %v", nested["pages"])
	}

	// Native in-place mutation without stringification
	fields["extraction_result"] = structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"total_amount": structpb.NewNumberValue(150000.50),
			"verified":     structpb.NewBoolValue(true),
		},
	})

	extracted := stateStruct.GetFields()["extraction_result"].GetStructValue().GetFields()
	if !extracted["verified"].GetBoolValue() {
		t.Fatalf("expected verified = true")
	}
	if extracted["total_amount"].GetNumberValue() != 150000.50 {
		t.Fatalf("expected total_amount = 150000.50")
	}
}
