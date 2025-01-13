package provider

import (
	"regexp"
	"testing"

	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNixSigningKey_basic(t *testing.T) {
	r.Test(t, r.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []r.TestStep{
			{
				Config: `
					resource "nix_signing_key" "test" {
						name = "test-key"
					}
				`,
				Check: r.ComposeAggregateTestCheckFunc(
					r.TestCheckResourceAttr("nix_signing_key.test", "name", "test-key"),
					r.TestMatchResourceAttr("nix_signing_key.test", "private_key", regexp.MustCompile(`^secret-key:`)),
					r.TestMatchResourceAttr("nix_signing_key.test", "public_key", regexp.MustCompile(`^public-key:`)),
				),
			},
		},
	})
}

func TestAccNixSigningKey_invalidName(t *testing.T) {
	r.Test(t, r.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []r.TestStep{
			{
				Config: `
					resource "nix_signing_key" "test" {
						name = ""
					}
				`,
				ExpectError: regexp.MustCompile(`name cannot be empty`),
			},
		},
	})
}

func TestAccNixSigningKey_update(t *testing.T) {
	r.Test(t, r.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []r.TestStep{
			{
				Config: `
					resource "nix_signing_key" "test" {
						name = "test-key"
					}
				`,
			},
			{
				Config: `
					resource "nix_signing_key" "test" {
						name = "updated-key"
					}
				`,
				ExpectError: regexp.MustCompile(`Updates not supported`),
			},
		},
	})
}

func TestAccNixSigningKey_import(t *testing.T) {
	r.Test(t, r.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []r.TestStep{
			{
				Config: `
					resource "nix_signing_key" "test" {
						name = "test-key"
					}
				`,
			},
			{
				ResourceName:      "nix_signing_key.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
