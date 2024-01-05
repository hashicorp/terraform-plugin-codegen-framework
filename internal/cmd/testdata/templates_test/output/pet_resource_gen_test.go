package provider_test

import (
	"context"
	"testing"

	"terraform-provider-petstore/internal/provider/resource_pet"
)

// Just a simple test for ensuring the pet resource template data is passed here
func TestPetResourceSchema(t *testing.T) {
	t.Parallel()

	schema := resource_pet.PetResourceSchema(context.Background())
	if schema == nil {
		t.Fatal("no schema generated!")
	}
}
