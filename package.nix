{
  terraform-providers,
  lib,
  homepage ? "https://registry.terraform.io/providers/aldoborrero/nix-signature",
  provider-source-address ? null,
  ...
}:
terraform-providers.mkProvider {
  owner = "aldoborrero";
  repo = "terraform-provider-nix-signature";
  rev = "v0.1.0";
  version = "0.1.0";
  vendorHash = "sha256-6UuRD0GvCEVY3N50y4mOJB3IaNgwfbT/I2eseokXHt0=";
  inherit homepage;
  mkProviderFetcher = _: lib.cleanSource ./.;
  hash = lib.fakeHash;
  spdx = "MIT";
}
// lib.optionalAttrs (provider-source-address != null) {
  inherit provider-source-address;
}
