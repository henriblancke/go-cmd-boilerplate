package main

import (
	"github.com/henriblancke/go-cmd-boilerplate/cmd"
	"github.com/henriblancke/go-cmd-boilerplate/config"
	"github.com/henriblancke/go-cmd-boilerplate/pkg/log"
)

func main() {

	config.Init()
	log.Init()
	cmd.Execute()

}
