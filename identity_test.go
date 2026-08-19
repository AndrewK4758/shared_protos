package shared_protos

import (
	"testing"
)

func TestValidateIdentity_TableDriven(t *testing.T) {
	tests := []struct {
		name        string
		identity    *InfrastructureIdentity
		expectError bool
		errorSubstr string
	}{
		{
			name:        "nil identity",
			identity:    nil,
			expectError: true,
			errorSubstr: "nil",
		},
		{
			name: "empty tenant_id",
			identity: &InfrastructureIdentity{
				TenantId: "",
				AppId:    "app-1",
				JobId:    "job-1",
			},
			expectError: true,
			errorSubstr: "tenant_id",
		},
		{
			name: "whitespace tenant_id",
			identity: &InfrastructureIdentity{
				TenantId: "   ",
				AppId:    "app-1",
				JobId:    "job-1",
			},
			expectError: true,
			errorSubstr: "tenant_id",
		},
		{
			name: "empty app_id",
			identity: &InfrastructureIdentity{
				TenantId: "tenant-1",
				AppId:    "",
				JobId:    "job-1",
			},
			expectError: true,
			errorSubstr: "app_id",
		},
		{
			name: "empty job_id",
			identity: &InfrastructureIdentity{
				TenantId: "tenant-1",
				AppId:    "app-1",
				JobId:    "",
			},
			expectError: true,
			errorSubstr: "job_id",
		},
		{
			name: "valid identity",
			identity: &InfrastructureIdentity{
				TenantId: "tenant-1",
				AppId:    "app-1",
				JobId:    "job-1",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateIdentity(tt.identity)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for %s, got nil", tt.name)
				}
				if tt.errorSubstr != "" && !contains(err.Error(), tt.errorSubstr) {
					t.Fatalf("expected error containing %q, got %v", tt.errorSubstr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestStorageNamespace_TableDriven(t *testing.T) {
	validId := &InfrastructureIdentity{
		TenantId: "acme-corp",
		AppId:    "invoicing",
		JobId:    "batch-992",
	}

	tests := []struct {
		name        string
		identity    *InfrastructureIdentity
		filename    string
		expected    string
		expectError bool
	}{
		{
			name:        "standard filename",
			identity:    validId,
			filename:    "document.pdf",
			expected:    "internal-store://acme-corp/invoicing/batch-992/document.pdf",
			expectError: false,
		},
		{
			name:        "filename with leading slash",
			identity:    validId,
			filename:    "/subfolder/packet.pdf",
			expected:    "internal-store://acme-corp/invoicing/batch-992/subfolder/packet.pdf",
			expectError: false,
		},
		{
			name:        "empty filename returns directory URI",
			identity:    validId,
			filename:    "",
			expected:    "internal-store://acme-corp/invoicing/batch-992/",
			expectError: false,
		},
		{
			name: "invalid identity returns error",
			identity: &InfrastructureIdentity{
				TenantId: "",
				AppId:    "app",
				JobId:    "job",
			},
			filename:    "doc.pdf",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri, err := tt.identity.StorageNamespace(tt.filename)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for %s, got nil", tt.name)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if uri != tt.expected {
				t.Fatalf("expected URI %q, got %q", tt.expected, uri)
			}
		})
	}
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && (str == substr || len(substr) == 0 || (len(str) > 0 && len(substr) > 0 && searchSubstr(str, substr)))
}

func searchSubstr(s, substr string) bool {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
