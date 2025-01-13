# terraform-provider-nix-signature

A Terraform provider for creating Nix signing keys used with binary caches.

## Features

- Generate Nix signing keypairs
- Automated key management through Terraform
- Compatible with Nix binary cache signing requirements

## Installation

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    nix-signature = {
      source = "aldoborrero/nix-signature"
    }
  }
}
```

## Usage

Generate a Nix signing key:

```hcl
resource "nix-signature_nix_signing_key" "cache_key" {
  name = "my-binary-cache"
}
```

### Output Attributes

- `private_key`: The generated private key (sensitive value)
- `public_key`: The public verification key to share with users

## Security Considerations

- Private keys are marked as sensitive and will be hidden in logs
- Keys are generated using `nixcommunity/go-nix` code
- The provider prevents key updates to avoid accidental key rotation

## Development

Requirements:

- Nix with Flakes enabled
- direnv (recommended)

Setup:

```bash
# Enter dev environment
direnv allow   # If using direnv
# or
nix develop
```

Commands:

- `fmt`: Format code
- `check`: Run checks

## License

See [LICENSE.md] for more information.
