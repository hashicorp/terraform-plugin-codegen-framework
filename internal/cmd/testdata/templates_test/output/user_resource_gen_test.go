package provider_test

import (
	"context"
	"testing"

	"terraform-provider-petstore/internal/provider/resource_user"
)

// Just a simple test for ensuring the user resource template data is passed here
func TestUserResourceSchema(t *testing.T) {
	t.Parallel()

	schema := resource_user.UserResourceSchema(context.Background())
	if schema == nil {
		t.Fatal("no schema generated!")
	}
}
