package shared_protos

import (
	"fmt"
	"strings"
)

// ValidateIdentity enforces strict fail-fast validation on InfrastructureIdentity.
// Returns an error immediately if the identity pointer is nil or any field is empty.
func ValidateIdentity(id *InfrastructureIdentity) error {
	if id == nil {
		return fmt.Errorf("infrastructure identity is nil: fail-fast validation error")
	}
	if strings.TrimSpace(id.GetTenantId()) == "" {
		return fmt.Errorf("infrastructure identity missing mandatory tenant_id")
	}
	if strings.TrimSpace(id.GetAppId()) == "" {
		return fmt.Errorf("infrastructure identity missing mandatory app_id")
	}
	if strings.TrimSpace(id.GetJobId()) == "" {
		return fmt.Errorf("infrastructure identity missing mandatory job_id")
	}
	return nil
}

// StorageNamespace generates the canonical multi-tenant storage URI:
// internal-store://{tenant_id}/{app_id}/{job_id}/{filename}
func (x *InfrastructureIdentity) StorageNamespace(filename string) (string, error) {
	if err := ValidateIdentity(x); err != nil {
		return "", err
	}
	cleanFile := strings.TrimPrefix(strings.TrimSpace(filename), "/")
	if cleanFile == "" {
		return fmt.Sprintf("internal-store://%s/%s/%s/", x.GetTenantId(), x.GetAppId(), x.GetJobId()), nil
	}
	return fmt.Sprintf("internal-store://%s/%s/%s/%s", x.GetTenantId(), x.GetAppId(), x.GetJobId(), cleanFile), nil
}
