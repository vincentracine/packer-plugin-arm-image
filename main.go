package main

import (
	"fmt"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/solo-io/packer-plugin-arm-image/pkg/builder"
	"github.com/solo-io/packer-plugin-arm-image/pkg/postprocessor"
	"github.com/solo-io/packer-plugin-arm-image/version"
	"os"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, builder.NewBuilder())
	pps.RegisterPostProcessor(plugin.DEFAULT_NAME, postprocessor.NewFlasher())
	pps.SetVersion(version.PluginVersion)
	fmt.Println(os.Args[1:])
	err := pps.Run()
	if err != nil {
		panic(err)
	}
}
