// Copyright (C) Andrew Harris 2024 - CyberArk, Inc.

package main

import (
	"context"
	"log"

	// CyberArk API
	provider "github.com/conjurdemos/terraform-provider-cyberark/cyberark"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name cyberark

var (
	version string = "0.9.9"
)

func main() {
	opts := providerserver.ServeOpts{
		Debug:   false,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}