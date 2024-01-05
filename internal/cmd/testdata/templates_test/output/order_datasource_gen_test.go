package provider_test

import (
	"context"
	"testing"

	"terraform-provider-petstore/internal/provider/datasource_order"
)

// Just a simple test for ensuring the order data source template data is passed here
func TestOrderDataSourceSchema(t *testing.T) {
	t.Parallel()

	schema := datasource_order.OrderDataSourceSchema(context.Background())
	if schema == nil {
		t.Fatal("no schema generated!")
	}
}
