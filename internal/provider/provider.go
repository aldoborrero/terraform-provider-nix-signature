package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type NixSignatureProvider struct{}

// Enforce interfaces we want provider to implement
var _ provider.Provider = (*NixSignatureProvider)(nil)

func New() provider.Provider {
	return &NixSignatureProvider{}
}

func (p *NixSignatureProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nix-signature"
}

func (p *NixSignatureProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider configuration",
	}
}

func (p *NixSignatureProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *NixSignatureProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNixSigningKeyResource,
	}
}

func (p *NixSignatureProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
