{
  description = "terraform-provider-nix-sinagure / A Terraform Provider for Nix Signatures";

  nixConfig = {
    extra-substituters = [
      "https://nix-community.cachix.org"
      "https://numtide.cachix.org"
      "https://cache.garnix.io"
    ];
    extra-trusted-public-keys = [
      "nix-community.cachix.org-1:mB9FSh9qf2dCimDSUo8Zy7bkq5CX+/rkCWyvRCYg3Fs="
      "numtide.cachix.org-1:2ps1kLBUWjxIneOy1Ik6cQjb41X0iXVXeHigGmycPPE="
      "cache.garnix.io:CTFPyKSLcx5RMJKfLo5EEPUObbA78b0YQ2DTCJXqr9g="
    ];
  };

  inputs = {
    # packages
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

    # flake-parts
    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };

    # go
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs = {
        nixpkgs.follows = "nixpkgs";
      };
    };

    # utilities
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    devshell = {
      url = "github:numtide/devshell";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    systems.url = "github:nix-systems/default";
    flake-compat = {
      url = "github:nix-community/flake-compat";
      flake = false;
    };
  };

  outputs = inputs @ {
    flake-parts,
    nixpkgs,
    ...
  }:
    flake-parts.lib.mkFlake
    {
      inherit inputs;
    }
    {
      imports = [
        inputs.devshell.flakeModule
        inputs.flake-parts.flakeModules.easyOverlay
        inputs.treefmt-nix.flakeModule
      ];

      debug = false;

      systems = import inputs.systems;

      perSystem = {
        pkgs,
        system,
        ...
      }: {
        # nixpkgs
        _module.args = {
          pkgs = import nixpkgs {
            inherit system;
            config.allowUnfree = true;
            overlays = [
              inputs.gomod2nix.overlays.default
            ];
          };
        };

        # packages
        packages = {
          terraform-provider-nix-signature = pkgs.callPackage ./package.nix {};
        };

        # devshells
        devshells.default = {
          name = "terraform-provider-nix-sinagure";
          packages = with pkgs; [
            cosign
            delve
            go
            golangci-lint
            gomod2nix
            goreleaser
            gotools
            terraform
            terraform-plugin-docs
          ];
          commands = [
            {
              name = "fmt";
              category = "nix";
              help = "format the source tree";
              command = ''nix fmt'';
            }
            {
              name = "check";
              category = "nix";
              help = "check the source tree";
              command = ''nix flake check'';
            }
            {
              name = "docs";
              category = "terraform";
              help = "generate terraform provider docs";
              command = ''tfplugindocs generate'';
            }
          ];
        };

        # treefmt
        treefmt.config = {
          flakeCheck = true;
          flakeFormatter = true;
          projectRootFile = "flake.nix";
          programs = {
            alejandra.enable = true;
            deadnix.enable = true;
            gofmt.enable = true;
            hclfmt.enable = true;
            shfmt.enable = true;
            statix.enable = true;
          };
          settings.formatter = {
            alejandra.priority = 3;
            deadnix.priority = 1;
            statix.priority = 2;
          };
        };
      };
    };
}
