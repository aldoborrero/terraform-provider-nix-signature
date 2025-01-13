package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aldoborrero/terraform-provider-nix-signature/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// automatically filled by goreleaser
var version = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address:         "registry.terraform.io/aldoborrero/nix-signature",
		Debug:           debug,
		ProtocolVersion: 6,
	})
	if err != nil {
		fmt.Printf("failed to initialize provider: %v\n", err)
		os.Exit(1)
	}
}
