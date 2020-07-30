package main

import (
	"github.com/MruV-RP/mruv-service-template/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("port", "3001")
	viper.SetDefault("texture_studio.path", "Texture-Studio")

	cmd.Execute()
}
