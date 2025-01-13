{
  buildGoModule,
  lib,
}:
buildGoModule rec {
  pname = "terraform-provider-nix-signature";
  version = "0.1.0";

  src = lib.cleanSource ./.;

  vendorHash = "sha256-6UuRD0GvCEVY3N50y4mOJB3IaNgwfbT/I2eseokXHt0=";

  ldflags = [
    "-s"
    "-w"
    "-X main.version=${version}"
  ];

  subPackages = ["."];

  postInstall = ''
    dir=$out/libexec/terraform-providers/registry.terraform.io/aldoborrero/nix-signature/${version}/''${GOOS}_''${GOARCH}
    mkdir -p "$dir"
    mv $out/bin/* "$dir/terraform-provider-nix-signature_${version}"
    rmdir $out/bin
  '';

  meta = with lib; {
    description = "Terraform provider for nix-signature";
    homepage = "https://github.com/aldoborrero/terraform-provider-nix-signature";
    license = licenses.mit;
    maintainers = with maintainers; [aldoborrero];
    mainProgram = "terraform-provider-nix-signature";
  };
}
