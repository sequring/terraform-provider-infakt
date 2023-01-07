package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/sequring/terraform-provider-infakt/infakt"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: infakt.Provider})
}
