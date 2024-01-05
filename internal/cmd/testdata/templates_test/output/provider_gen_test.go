package provider_test

import (
	"terraform-provider-petstore/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// Just a simple test for ensuring the petstore provider template data is passed here
func protoV5ProviderFactories() map[string]func() (tfprotov5.ProviderServer, error) {
	return map[string]func() (tfprotov5.ProviderServer, error){
		"petstore": providerserver.NewProtocol5WithError(provider.New()),
	}
}
