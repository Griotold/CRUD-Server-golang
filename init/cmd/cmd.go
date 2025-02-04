package cmd

import (
	"awesomeProject/config"
	"awesomeProject/network"
	"fmt"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	fmt.Println(c.config.Server.Port)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
