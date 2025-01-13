package provider

import (
	"context"
	"crypto/rand"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nix-community/go-nix/pkg/narinfo/signature"
)

type NixSigningKeyResource struct{}

var _ resource.Resource = (*NixSigningKeyResource)(nil)

func NewNixSigningKeyResource() resource.Resource {
	return &NixSigningKeyResource{}
}

type NixSigningKeyResourceModel struct {
	Name       types.String `tfsdk:"name"`
	PrivateKey types.String `tfsdk:"private_key"`
	PublicKey  types.String `tfsdk:"public_key"`
}

func (r *NixSigningKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nix_signing_key"
}

func (r *NixSigningKeyResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates a Nix signing key pair for use with binary caches",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name used to identify the signing key",
			},
			"private_key": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The private signing key in Nix-compatible format",
			},
			"public_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The public verification key in Nix-compatible format",
			},
		},
	}
}

func (r *NixSigningKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NixSigningKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	secretKey, publicKey, err := signature.GenerateKeypair(plan.Name.ValueString(), rand.Reader)
	if err != nil {
		resp.Diagnostics.AddError("Error generating keypair", err.Error())
		return
	}

	plan.PrivateKey = types.StringValue(secretKey.String())
	plan.PublicKey = types.StringValue(publicKey.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NixSigningKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NixSigningKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
}

func (r *NixSigningKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Updates not supported", "This resource does not support updates")
}

func (r *NixSigningKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
