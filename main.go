package main

import (
	"github.com/ewbankkit/EvilSuperstars/decode"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: decode.Provider,
	})
}
